package view

import (
	"context"
	"time"

	"github.com/titpetric/platform-example/blog/model"
)

// PostData holds the data required for rendering the post layout
type PostData struct {
	Slug        string
	Title       string
	Description string
	OGImage     string
	Content     string
	Date        time.Time
	Classnames  string
}

// Map converts PostData to a map[string]any
func (d *PostData) Map() map[string]any {
	return map[string]any{
		"slug":        d.Slug,
		"title":       d.Title,
		"description": d.Description,
		"ogImage":     d.OGImage,
		"content":     d.Content,
		"date":        d.Date,
		"classnames":  d.Classnames,
	}
}

// Post renders the post layout template
func (v *Views) Post(ctx context.Context, data *PostData) (string, error) {
	// Build the context data
	templateData := data.Map()
	for k, v := range v.data {
		if _, ok := templateData[k]; !ok {
			templateData[k] = v
		}
	}

	// Render the post layout
	return v.RenderPage(ctx, "layouts/post.vuego", templateData)
}

// PostFromArticle creates PostData from an Article
func (v *Views) PostFromArticle(article *model.Article, content string) *PostData {
	return &PostData{
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		OGImage:     article.OGImage,
		Content:     content,
		Date:        article.Date,
		Classnames:  "prose",
	}
}
