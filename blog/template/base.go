package template

import (
	"bytes"
	"context"

	"github.com/titpetric/platform-example/blog/model"
)

// BaseData holds the data required for rendering the base layout
type BaseData struct {
	Title       string
	Description string
	OGImage     string
	Content     string
	Classnames  string
}

// Base renders the base layout template
func (v *Views) Base(ctx context.Context, data *BaseData) (string, error) {
	var vue = v.vue

	// Build the context data
	templateData := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"ogImage":     data.OGImage,
		"content":     data.Content,
		"classnames":  data.Classnames,
	}
	for k, v := range v.data {
		if _, ok := templateData[k]; !ok {
			templateData[k] = v
		}
	}

	var buf bytes.Buffer
	err := vue.RenderFragment(&buf, "theme/layouts/base.vuego", templateData)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// BaseFromArticle creates BaseData from an Article
func (v *Views) BaseFromArticle(article *model.Article, content string) *BaseData {
	return &BaseData{
		Title:       article.Title,
		Description: article.Description,
		OGImage:     article.OGImage,
		Content:     content,
		Classnames:  "prose",
	}
}
