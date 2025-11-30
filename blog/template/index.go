package template

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/titpetric/platform-example/blog/model"
)

// IndexData holds the data required for rendering the index page
type IndexData struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	OGImage     string          `json:"ogImage"`
	Articles    []model.Article `json:"articles"`
	Total       int             `json:"total"`
}

// Index renders the blog index/list page
func (v *Views) Index(ctx context.Context, data *IndexData) (string, error) {
	var vue = v.vue

	// Build the context data
	templateData := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"ogImage":     data.OGImage,
		"articles":    data.Articles,
		"total":       data.Total,
	}
	for k, v := range v.data {
		if _, ok := templateData[k]; !ok {
			templateData[k] = v
		}
	}

	// First, render the index content as a fragment
	var fragmentBuf strings.Builder
	if err := vue.RenderFragment(&fragmentBuf, "pages/index.vuego", templateData); err != nil {
		return "", err
	}

	// Pass the rendered content to the base layout
	templateData["content"] = fragmentBuf.String()

	var layoutBuf strings.Builder
	// Re-apply funcMap before rendering base layout
	if err := vue.Render(&layoutBuf, "layouts/base.vuego", templateData); err != nil {
		return "", err
	}
	return layoutBuf.String(), nil
}

func fillTemplateData(w *map[string]any) error {
	if err := loadFile(w, "navigation", "config/navigation.json"); err != nil {
		return err
	}
	if err := loadFile(w, "themes", "config/themes.json"); err != nil {
		return err
	}
	if err := loadFile(w, "meta", "config/meta.json"); err != nil {
		return err
	}
	return nil
}

func loadFile(w *map[string]any, key string, filename string) error {
	var result any
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&result); err != nil {
		return fmt.Errorf("error loading %s: %w", filename, err)
	}
	(*w)[key] = result
	log.Println("loaded ok:", key, filename)
	return nil
}

// IndexFromArticles creates IndexData from a list of articles
func (v *Views) IndexFromArticles(articles []model.Article) *IndexData {
	return &IndexData{
		Title:       "Blog",
		Description: "Read my latest articles and posts",
		Articles:    articles,
		Total:       len(articles),
	}
}
