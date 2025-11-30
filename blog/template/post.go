package template

import (
	"bytes"
	"context"
	"time"

	"github.com/titpetric/platform-example/blog/model"
)

// PostData holds the data required for rendering the post layout
type PostData struct {
	Title       string
	Description string
	OGImage     string
	Content     string
	Date        time.Time
	Classnames  string
}

// Post renders the post layout template
func (v *Views) Post(ctx context.Context, data *PostData) (string, error) {
	var vue = v.vue

	// Build the context data
	templateData := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"ogImage":     data.OGImage,
		"content":     data.Content,
		"date":        data.Date,
		"classnames":  data.Classnames,
	}
	for k, v := range v.data {
		if _, ok := templateData[k]; !ok {
			templateData[k] = v
		}
	}

	// First, render the post content as a fragment
	var fragmentBuf bytes.Buffer
	err := vue.RenderFragment(&fragmentBuf, "layouts/post.vuego", templateData)
	if err != nil {
		return "", err
	}

	// Pass the rendered content to the base layout
	templateData["content"] = fragmentBuf.String()

	var layoutBuf bytes.Buffer
	// Re-apply funcMap before rendering base layout
	err = vue.Render(&layoutBuf, "layouts/base.vuego", templateData)
	if err != nil {
		return "", err
	}

	return layoutBuf.String(), nil
}

// PostFromArticle creates PostData from an Article
func (v *Views) PostFromArticle(article *model.Article, content string) *PostData {
	return &PostData{
		Title:       article.Title,
		Description: article.Description,
		OGImage:     article.OGImage,
		Content:     content,
		Date:        article.Date,
		Classnames:  "prose",
	}
}
