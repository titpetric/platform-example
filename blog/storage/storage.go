package storage

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/titpetric/platform-example/blog/model"
	"github.com/titpetric/platform-example/blog/schema"
)

// Storage provides database operations for the blog module
type Storage struct {
	db *sqlx.DB
}

// NewStorage creates a new Storage instance
func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// GetArticleBySlug retrieves an article by its slug
func (s *Storage) GetArticleBySlug(ctx context.Context, slug string) (*model.Article, error) {
	return GetArticleBySlug(ctx, s.db, slug)
}

// GetArticles retrieves all articles
func (s *Storage) GetArticles(ctx context.Context, start, length int) ([]model.Article, error) {
	return GetArticles(ctx, s.db, start, length)
}

// SearchArticles performs a full-text search on articles
func (s *Storage) SearchArticles(ctx context.Context, query string) ([]model.Article, error) {
	return SearchArticles(ctx, s.db, query)
}

// InsertArticle inserts a new article
func (s *Storage) InsertArticle(ctx context.Context, article *model.Article) error {
	return InsertArticle(ctx, s.db, article)
}

// CountArticles returns the total count of articles
func (s *Storage) CountArticles(ctx context.Context) (int, error) {
	return CountArticles(ctx, s.db)
}

// InitSchema initializes the database schema from embedded schema
func (s *Storage) InitSchema(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, schema.InitialSchema)
	return err
}
