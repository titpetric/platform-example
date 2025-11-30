package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/titpetric/platform-example/blog/model"
)

// GetArticleBySlug retrieves a single article by slug
func GetArticleBySlug(ctx context.Context, db *sqlx.DB, slug string) (*model.Article, error) {
	query := `
	SELECT * FROM article
	WHERE slug = ?
	LIMIT 1
	`

	var article model.Article
	err := db.GetContext(ctx, &article, query, slug)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// GetArticles retrieves all articles ordered by date descending
func GetArticles(ctx context.Context, db *sqlx.DB, start, length int) ([]model.Article, error) {
	query := fmt.Sprintf(`SELECT * FROM article ORDER BY date DESC LIMIT %d, %d`, start, length)

	var articles []model.Article
	err := db.SelectContext(ctx, &articles, query)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// SearchArticles performs a search on articles by title, description, or content
func SearchArticles(ctx context.Context, db *sqlx.DB, query string) ([]model.Article, error) {
	searchTerm := "%" + query + "%"

	sqlQuery := `
	SELECT * FROM article
	WHERE title LIKE ? OR description LIKE ? OR slug LIKE ?
	ORDER BY date DESC
	`

	var articles []model.Article
	err := db.SelectContext(ctx, &articles, sqlQuery, searchTerm, searchTerm, searchTerm)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// InsertArticle inserts a new article into the database
func InsertArticle(ctx context.Context, db *sqlx.DB, article *model.Article) error {
	query := `
	INSERT OR REPLACE INTO article (
		id, slug, title, description, content, date,
		og_image, layout, source, url, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.ExecContext(ctx, query,
		article.ID,
		article.Slug,
		article.Title,
		article.Description,
		article.Content,
		article.Date,
		article.OGImage,
		article.Layout,
		article.Source,
		article.URL,
		article.CreatedAt,
		article.UpdatedAt,
	)

	return err
}

// CountArticles returns the total number of articles
func CountArticles(ctx context.Context, db *sqlx.DB) (int, error) {
	var count int
	err := db.GetContext(ctx, &count, "SELECT COUNT(*) FROM article")
	return count, err
}
