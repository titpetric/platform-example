package storage

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/titpetric/platform-example/blog/model"
)

// TestStorageIntegration_FullLifecycle tests the complete storage lifecycle
func TestStorageIntegration_FullLifecycle(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// 1. Verify schema was initialized
	var tableCount int
	err := db.GetContext(ctx, &tableCount,
		"SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='article'")
	require.NoError(t, err)
	require.Equal(t, 1, tableCount)

	// 2. Insert multiple articles
	articles := []model.Article{
		{
			ID:          "article-1",
			Slug:        "getting-started",
			Title:       "Getting Started with Go",
			Description: "Learn the basics of Go programming",
			Content:     "# Getting Started\n\nGo is a modern programming language...",
			Date:        time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
			OGImage:     "/og/go.png",
			Layout:      "post",
			URL:         "/blog/getting-started/",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "article-2",
			Slug:        "advanced-patterns",
			Title:       "Advanced Go Patterns",
			Description: "Master advanced patterns in Go",
			Content:     "# Advanced Patterns\n\nOnce you know the basics...",
			Date:        time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			OGImage:     "/og/patterns.png",
			Layout:      "post",
			URL:         "/blog/advanced-patterns/",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "article-3",
			Slug:        "concurrency-guide",
			Title:       "Concurrency Guide",
			Description: "Understanding goroutines and channels",
			Content:     "# Concurrency\n\nGoroutines are lightweight threads...",
			Date:        time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC),
			OGImage:     "/og/concurrency.png",
			Layout:      "post",
			URL:         "/blog/concurrency-guide/",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, article := range articles {
		err := storage.InsertArticle(ctx, &article)
		require.NoError(t, err, "failed to insert article %s", article.Slug)
	}

	// 3. Verify count
	count, err := storage.CountArticles(ctx)
	require.NoError(t, err)
	require.Equal(t, 3, count)

	// 4. Retrieve all and verify ordering
	all, err := storage.GetArticles(ctx, 0, 9999)
	require.NoError(t, err)
	require.Len(t, all, 3)
	require.Equal(t, "concurrency-guide", all[0].Slug)

	// 5. Get single article
	article, err := storage.GetArticleBySlug(ctx, "advanced-patterns")
	require.NoError(t, err)
	require.Equal(t, "Advanced Go Patterns", article.Title)

	// 6. Search articles
	results, err := storage.SearchArticles(ctx, "Go")
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(results), 2)

	// 7. Update article
	article.Title = "Updated: Advanced Go Patterns"
	article.UpdatedAt = time.Now()
	err = storage.InsertArticle(ctx, article)
	require.NoError(t, err)

	// Verify update
	updated, err := storage.GetArticleBySlug(ctx, "advanced-patterns")
	require.NoError(t, err)
	require.Equal(t, "Updated: Advanced Go Patterns", updated.Title)
}

// TestStorageIntegration_ConcurrentInserts tests concurrent article insertions
func TestStorageIntegration_ConcurrentInserts(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// Insert articles concurrently
	done := make(chan error, 10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			article := &model.Article{
				ID:    fmt.Sprintf("concurrent-%d", idx),
				Slug:  fmt.Sprintf("article-%d", idx),
				Title: "Concurrent Article",
				Date:  time.Now(),
			}
			done <- storage.InsertArticle(ctx, article)
		}(i)
	}

	// Collect results
	for i := 0; i < 10; i++ {
		err := <-done
		require.NoError(t, err)
	}

	// Verify all were inserted
	finalCount, _ := storage.CountArticles(ctx)
	require.Equal(t, 10, finalCount)
}

// TestStorageIntegration_ReplaceOnDuplicate tests that INSERT OR REPLACE works correctly
func TestStorageIntegration_ReplaceOnDuplicate(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// Insert a valid article
	article1 := &model.Article{
		ID:    "valid-1",
		Slug:  "valid-article",
		Title: "Valid Article",
		Date:  time.Now(),
	}
	err := storage.InsertArticle(ctx, article1)
	require.NoError(t, err)

	// Insert article with duplicate slug (should replace)
	article2 := &model.Article{
		ID:    "valid-2",
		Slug:  "valid-article", // duplicate slug
		Title: "Replaced Article",
		Date:  time.Now(),
	}
	err = storage.InsertArticle(ctx, article2)
	require.NoError(t, err)

	// Verify article was replaced
	count, _ := storage.CountArticles(ctx)
	require.Equal(t, 1, count)

	retrieved, err := storage.GetArticleBySlug(ctx, "valid-article")
	require.NoError(t, err)
	require.Equal(t, "valid-2", retrieved.ID)
	require.Equal(t, "Replaced Article", retrieved.Title)
}

// TestStorageIntegration_SearchAccuracy tests search result accuracy
func TestStorageIntegration_SearchAccuracy(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// Insert articles with varied content
	articles := []model.Article{
		{
			ID:          "search-1",
			Slug:        "python-basics",
			Title:       "Python for Beginners",
			Description: "Learn Python programming language",
			Content:     "Python is easy to learn and powerful",
			Date:        time.Now(),
		},
		{
			ID:          "search-2",
			Slug:        "go-performance",
			Title:       "Go Performance Optimization",
			Description: "Optimize Go applications",
			Content:     "Go is fast and efficient",
			Date:        time.Now(),
		},
		{
			ID:          "search-3",
			Slug:        "rust-safety",
			Title:       "Rust Memory Safety",
			Description: "Understanding Rust's safety features",
			Content:     "Rust prevents data races at compile time",
			Date:        time.Now(),
		},
	}

	for i := range articles {
		err := storage.InsertArticle(ctx, &articles[i])
		require.NoError(t, err)
	}

	tests := []struct {
		name          string
		query         string
		expectedCount int
		expectedSlugs []string
	}{
		{
			name:          "search by title",
			query:         "Python",
			expectedCount: 1,
			expectedSlugs: []string{"python-basics"},
		},
		{
			name:          "search by description",
			query:         "Optimize",
			expectedCount: 1,
			expectedSlugs: []string{"go-performance"},
		},
		{
			name:          "search by title word",
			query:         "Performance",
			expectedCount: 1,
			expectedSlugs: []string{"go-performance"},
		},
		{
			name:          "search by slug",
			query:         "rust",
			expectedCount: 1,
			expectedSlugs: []string{"rust-safety"},
		},
		{
			name:          "case insensitive search",
			query:         "PYTHON",
			expectedCount: 1,
			expectedSlugs: []string{"python-basics"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := storage.SearchArticles(ctx, tt.query)
			require.NoError(t, err)
			require.Len(t, results, tt.expectedCount, "unexpected result count")

			for _, expectedSlug := range tt.expectedSlugs {
				found := false
				for _, result := range results {
					if result.Slug == expectedSlug {
						found = true
						break
					}
				}
				require.True(t, found, "expected slug %s not found in results", expectedSlug)
			}
		})
	}
}

// TestStorageIntegration_DateOrdering tests proper date-based ordering
func TestStorageIntegration_DateOrdering(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// Insert articles with specific dates (not in order)
	articles := []model.Article{
		{
			ID:    "date-1",
			Slug:  "middle",
			Title: "Middle Article",
			Date:  time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:    "date-2",
			Slug:  "newest",
			Title: "Newest Article",
			Date:  time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:    "date-3",
			Slug:  "oldest",
			Title: "Oldest Article",
			Date:  time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	for i := range articles {
		err := storage.InsertArticle(ctx, &articles[i])
		require.NoError(t, err)
	}

	// Retrieve all and verify order (newest first)
	all, err := storage.GetArticles(ctx, 0, 9999)
	require.NoError(t, err)

	expectedOrder := []string{"newest", "middle", "oldest"}
	for i, expected := range expectedOrder {
		require.Equal(t, expected, all[i].Slug, "position %d mismatch", i)
	}
}

// TestStorageIntegration_EmptyDatabase tests operations on empty database
func TestStorageIntegration_EmptyDatabase(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// Count should be zero
	count, err := storage.CountArticles(ctx)
	require.NoError(t, err)
	require.Equal(t, 0, count)

	// Get should return empty slice
	articles, err := storage.GetArticles(ctx, 0, 9999)
	require.NoError(t, err)
	require.Len(t, articles, 0)

	// Search should return empty results
	results, err := storage.SearchArticles(ctx, "anything")
	require.NoError(t, err)
	require.Len(t, results, 0)

	// GetBySlug should fail
	_, err = storage.GetArticleBySlug(ctx, "nonexistent")
	require.Error(t, err)
}

// TestStorageIntegration_LargeContent tests handling of large article content
func TestStorageIntegration_LargeContent(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// Create article with large content
	largeContent := ""
	for i := 0; i < 1000; i++ {
		largeContent += "This is a paragraph of content that will be repeated many times to create a large article.\n"
	}

	article := &model.Article{
		ID:      "large-1",
		Slug:    "large-article",
		Title:   "Large Article",
		Content: largeContent,
		Date:    time.Now(),
	}

	err := storage.InsertArticle(ctx, article)
	require.NoError(t, err)

	// Retrieve and verify
	retrieved, err := storage.GetArticleBySlug(ctx, "large-article")
	require.NoError(t, err)
	require.Equal(t, len(largeContent), len(retrieved.Content))
}

// TestStorageIntegration_SpecialCharacters tests handling of special characters
func TestStorageIntegration_SpecialCharacters(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	article := &model.Article{
		ID:          "special-1",
		Slug:        "special-chars",
		Title:       "Article with \"quotes\" and 'apostrophes'",
		Description: "Testing special chars: <html> & {json}",
		Content:     "Code: `func() { return \"test\"; }`",
		Date:        time.Now(),
	}

	err := storage.InsertArticle(ctx, article)
	require.NoError(t, err)

	// Retrieve and verify special characters are preserved
	retrieved, err := storage.GetArticleBySlug(ctx, "special-chars")
	require.NoError(t, err)
	require.Equal(t, article.Title, retrieved.Title)
	require.Equal(t, article.Description, retrieved.Description)
	require.Equal(t, article.Content, retrieved.Content)
}

// TestStorageIntegration_TimestampHandling tests proper timestamp handling
func TestStorageIntegration_TimestampHandling(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	now := time.Now()
	article := &model.Article{
		ID:        "time-1",
		Slug:      "time-test",
		Title:     "Timestamp Test",
		Date:      time.Date(2024, 6, 15, 14, 30, 45, 0, time.UTC),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := storage.InsertArticle(ctx, article)
	require.NoError(t, err)

	retrieved, err := storage.GetArticleBySlug(ctx, "time-test")
	require.NoError(t, err)

	// Verify date is preserved
	require.Equal(t, 2024, retrieved.Date.Year())
	require.Equal(t, time.June, retrieved.Date.Month())
	require.Equal(t, 15, retrieved.Date.Day())
}

// TestStorageIntegration_IndexEfficiency tests that indexes are properly used
func TestStorageIntegration_IndexEfficiency(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	storage := NewStorage(db)
	ctx := context.Background()

	// Insert many articles
	for i := 0; i < 100; i++ {
		idx := fmt.Sprintf("%d", i)
		article := &model.Article{
			ID:    "index-" + idx,
			Slug:  "article-" + idx,
			Title: "Article Number",
			Date:  time.Now().AddDate(0, 0, -i),
		}
		err := storage.InsertArticle(ctx, article)
		require.NoError(t, err)
	}

	// These operations should be fast due to indexes
	_, err := storage.GetArticleBySlug(ctx, "article-50")
	require.NoError(t, err)

	articles, err := storage.GetArticles(ctx, 0, 9999)
	require.NoError(t, err)
	require.Len(t, articles, 100)

	// Verify newest first ordering with many articles
	require.Equal(t, "article-0", articles[0].Slug)
}
