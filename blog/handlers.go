package blog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"

	"github.com/titpetric/platform-example/blog/markdown"
	"github.com/titpetric/platform-example/blog/model"
	"github.com/titpetric/platform-example/blog/storage"
	"github.com/titpetric/platform-example/blog/view"
)

// Handlers handles HTTP requests for the blog module
type Handlers struct {
	repository *storage.Storage
	views      *view.Views
}

// NewHandlers creates a new Handlers instance with the given storage
func NewHandlers(repo *storage.Storage) (*Handlers, error) {
	views, err := view.NewViews(os.DirFS("theme"))
	if err != nil {
		return nil, err
	}

	return &Handlers{
		repository: repo,
		views:      views,
	}, nil
}

// ListArticlesJSON returns a JSON list of all articles
func (h *Handlers) ListArticlesJSON(w http.ResponseWriter, r *http.Request) {
	articles, err := h.repository.GetArticles(r.Context(), 0, 9999)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch articles: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=300")

	list := &model.ArticleList{
		Articles: articles,
		Total:    len(articles),
		Page:     1,
		PageSize: len(articles),
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetArticleJSON returns a single article as JSON
func (h *Handlers) GetArticleJSON(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	article, err := h.repository.GetArticleBySlug(r.Context(), slug)
	if err != nil {
		http.Error(w, fmt.Sprintf("article not found: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=3600")

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SearchArticlesJSON performs full-text search on articles
func (h *Handlers) SearchArticlesJSON(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing 'q' query parameter", http.StatusBadRequest)
		return
	}

	articles, err := h.repository.SearchArticles(r.Context(), query)
	if err != nil {
		http.Error(w, fmt.Sprintf("search failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=300")

	result := map[string]interface{}{
		"articles": articles,
		"total":    len(articles),
		"query":    query,
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// IndexHTML returns an HTML index page listing blogs
func (h *Handlers) IndexHTML(w http.ResponseWriter, r *http.Request) {
	articles, err := h.repository.GetArticles(r.Context(), 0, 5)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch articles: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=300")

	// Create index component to render list
	indexData := h.views.IndexFromArticles(articles)
	html, err := h.views.Index(r.Context(), indexData)
	if err != nil {
		http.Error(w, fmt.Sprintf("render failed: %v\n (%s)", err, html), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(html))
}

// ListArticlesHTML returns an HTML list of articles
func (h *Handlers) ListArticlesHTML(w http.ResponseWriter, r *http.Request) {
	articles, err := h.repository.GetArticles(r.Context(), 0, 9999)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch articles: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=300")

	// Create blog list and render
	blogData := h.views.IndexFromArticles(articles)
	html, err := h.views.Blog(r.Context(), blogData)
	if err != nil {
		http.Error(w, fmt.Sprintf("render failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(html))
}

// GetArticleHTML returns a single article as HTML
func (h *Handlers) GetArticleHTML(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	article, err := h.repository.GetArticleBySlug(r.Context(), slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")

	// Strip front matter and convert markdown to HTML with syntax highlighting
	contentWithoutFrontMatter := view.StripFrontMatter(article.Content)
	mdRenderer := markdown.NewRenderer()
	htmlContent := mdRenderer.Render([]byte(contentWithoutFrontMatter))

	// Create PostData and render
	postData := h.views.PostFromArticle(article, string(htmlContent))
	html, err := h.views.Post(r.Context(), postData)
	if err != nil {
		http.Error(w, fmt.Sprintf("render failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(html))
}

// GetAtomFeed returns an Atom XML feed of all articles
func (h *Handlers) GetAtomFeed(w http.ResponseWriter, r *http.Request) {
	articles, err := h.repository.GetArticles(r.Context(), 0, 20)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch articles: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")

	// Generate atom feed
	xml, err := h.views.AtomFeed(r.Context(), articles)
	if err != nil {
		http.Error(w, fmt.Sprintf("feed generation failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(xml))
}
