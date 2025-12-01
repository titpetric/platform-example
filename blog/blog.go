package blog

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"

	"github.com/titpetric/platform"
	yaml "gopkg.in/yaml.v3"

	"github.com/titpetric/platform-example/blog/model"
	"github.com/titpetric/platform-example/blog/storage"
)

// Module implements the blog module for the platform
type Module struct {
	platform.UnimplementedModule

	// Data directory for markdown files
	dataDir string

	// Storage for database operations
	repository *storage.Storage

	// Articles index for in-memory access
	articles map[string]*model.Article
}

// NewModule creates a new blog module instance
func NewModule(dataDir string) *Module {
	return &Module{
		dataDir:  dataDir,
		articles: make(map[string]*model.Article),
	}
}

// Name returns the module name
func (m *Module) Name() string {
	return "blog"
}

// Mount registers the blog routes with the router
func (m *Module) Mount(_ context.Context, router platform.Router) error {
	// Create handlers using the module's storage
	h, err := NewHandlers(m.repository)
	if err != nil {
		return err
	}

	// Static files
	assetFS := http.StripPrefix("/assets", http.FileServer(http.Dir("theme/assets")))
	router.Get("/assets/css/*", func(w http.ResponseWriter, r *http.Request) { assetFS.ServeHTTP(w, r) })
	router.Get("/assets/fonts/*", func(w http.ResponseWriter, r *http.Request) { assetFS.ServeHTTP(w, r) })
	router.Get("/assets/icons/*", func(w http.ResponseWriter, r *http.Request) { assetFS.ServeHTTP(w, r) })
	router.Get("/assets/favicon/*", func(w http.ResponseWriter, r *http.Request) { assetFS.ServeHTTP(w, r) })
	router.Get("/assets/robots.txt", func(w http.ResponseWriter, r *http.Request) { assetFS.ServeHTTP(w, r) })
	router.Get("/assets/site.webmanifest", func(w http.ResponseWriter, r *http.Request) { assetFS.ServeHTTP(w, r) })

	// API Routes (JSON)
	router.Get("/api/blog/articles", h.ListArticlesJSON)
	router.Get("/api/blog/articles/{slug}", h.GetArticleJSON)
	router.Get("/api/blog/search", h.SearchArticlesJSON)

	// HTML Routes
	router.Get("/", h.IndexHTML)
	router.Get("/blog/", h.ListArticlesHTML)
	router.Get("/blog/{slug}", h.GetArticleHTML)
	router.Get("/blog/{slug}/", h.GetArticleHTML)

	// Feed Routes
	router.Get("/feed.xml", h.GetAtomFeed)

	return nil
}

// Start initializes the blog module by scanning markdown files and building the index
func (m *Module) Start(ctx context.Context) error {
	// Get database connection from platform
	db, err := storage.DB(ctx)
	if err != nil {
		return fmt.Errorf("failed to get blog database: %w", err)
	}

	// Create storage instance
	m.repository = storage.NewStorage(db)

	// Create schema
	if err := m.repository.InitSchema(ctx); err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	// Scan and index markdown files
	count, err := m.scanMarkdownFiles(ctx)
	if err != nil {
		return fmt.Errorf("failed to scan markdown files: %w", err)
	}

	fmt.Printf("[blog] scanned %d markdown files from %s\n", count, m.dataDir)

	// Verify articles were inserted
	total, err := m.repository.CountArticles(ctx)
	if err != nil {
		return fmt.Errorf("failed to count articles: %w", err)
	}
	fmt.Printf("[blog] verified %d articles in database\n", total)

	return nil
}

// Stop is called when the module is shutting down
func (m *Module) Stop(context.Context) error {
	// Nothing to clean up - database is managed by platform
	return nil
}

// SetRepository sets the repository on the module
func (m *Module) SetRepository(repo *storage.Storage) {
	m.repository = repo
}

// ScanMarkdownFiles scans the data directory for markdown files and indexes them
// Returns the count of scanned files
func (m *Module) ScanMarkdownFiles(ctx context.Context) (int, error) {
	return m.scanMarkdownFiles(ctx)
}

// scanMarkdownFiles scans the data directory for markdown files and indexes them
// Returns the count of scanned files
func (m *Module) scanMarkdownFiles(ctx context.Context) (int, error) {
	count := 0
	err := filepath.WalkDir(m.dataDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only process markdown files
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		count++

		article, err := m.parseMarkdownFile(path)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		// Store in memory map
		m.articles[article.Slug] = article

		// Insert into database
		err = m.repository.InsertArticle(ctx, article)
		if err != nil {
			return fmt.Errorf("failed to insert article %s: %w", article.Slug, err)
		}

		return nil
	})
	return count, err
}

// parseMarkdownFile parses a markdown file and extracts metadata
func (m *Module) parseMarkdownFile(filePath string) (*model.Article, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	content := string(data)

	// Extract YAML front matter
	var meta model.Metadata
	var markdownContent string

	// Check if file starts with ---
	if strings.HasPrefix(content, "---") {
		parts := strings.SplitN(content, "---", 3)
		if len(parts) >= 3 {
			if err := yaml.Unmarshal([]byte(parts[1]), &meta); err != nil {
				return nil, fmt.Errorf("failed to parse YAML front matter: %w", err)
			}
			markdownContent = parts[2]
		}
	} else {
		markdownContent = content
	}

	// Generate article ID and slug
	fileName := filepath.Base(filePath)
	slug := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	id := generateID(slug)

	// Parse date
	var date time.Time
	if meta.Date != "" {
		var err error
		date, err = time.Parse("2006-01-02", meta.Date)
		if err != nil {
			date = time.Now()
		}
	} else {
		date = time.Now()
	}

	// Set default layout if not provided
	layout := meta.Layout
	if layout == "" {
		layout = "post"
	}

	article := &model.Article{
		ID:          id,
		Slug:        slug,
		Title:       meta.Title,
		Description: meta.Description,
		Content:     markdownContent,
		Date:        date,
		OGImage:     meta.OGImage,
		Layout:      layout,
		Source:      meta.Source,
		URL:         "/blog/" + slug + "/",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return article, nil
}

// generateID creates a unique ID from slug
func generateID(slug string) string {
	return slug + "-" + time.Now().Format("20060102150405")
}
