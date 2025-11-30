package model

import (
	"time"
)

// Article represents a blog article parsed from markdown files
type Article struct {
	ID          string    `db:"id" json:"id"`
	Slug        string    `db:"slug" json:"slug"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Content     string    `db:"content" json:"content"`
	Date        time.Time `db:"date" json:"date"`
	OGImage     string    `db:"og_image" json:"ogImage,omitempty"`
	Layout      string    `db:"layout" json:"layout"`
	Source      string    `db:"source" json:"source,omitempty"`
	ReadingTime string    `db:"reading_time" json:"readingTime,omitempty"`
	URL         string    `db:"url" json:"url"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

// Metadata represents the YAML front matter of a markdown file
type Metadata struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	OGImage     string `yaml:"ogImage"`
	Date        string `yaml:"date"`
	Layout      string `yaml:"layout"`
	Source      string `yaml:"source"`
}

// ArticleList represents a paginated list of articles
type ArticleList struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
}
