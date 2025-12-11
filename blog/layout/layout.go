package layout

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"

	"github.com/titpetric/vuego"
)

// Renderer handles page and layout rendering with vuego templates
type Renderer struct {
	root fs.FS
	data map[string]any
}

// NewRenderer creates a new Renderer with the given filesystem and shared data
func NewRenderer(root fs.FS, data map[string]any) *Renderer {
	return &Renderer{
		root: root,
		data: data,
	}
}

// template creates a vuego template with shared data and custom functions
func (r *Renderer) template(data map[string]any, funcs map[string]interface{}) vuego.Template {
	tpl := vuego.Load(r.root, vuego.WithLessProcessor())
	return tpl.Funcs(funcs).Fill(data)
}

// RenderPage loads a page template, renders it, and wraps it with a layout
// It passes templateData to the page, then adds the rendered page as "content" to the layout
func (r *Renderer) RenderPage(ctx context.Context, pagePath string, templateData map[string]any, funcs map[string]interface{}) (string, error) {
	tpl := r.template(templateData, funcs)

	// Render the page template
	var pageBuf bytes.Buffer
	if err := tpl.Render(ctx, &pageBuf, pagePath); err != nil {
		return "", err
	}

	templateData["content"] = pageBuf.String()

	// Get layout name from template metadata
	layoutName := "layouts/base.vuego"
	if layout, ok := tpl.GetVar("layout").(string); ok && layout != "" {
		layoutName = fmt.Sprintf("layouts/%s.vuego", layout)
	}

	// Render with layout
	return r.RenderLayout(ctx, layoutName, templateData, funcs)
}

// RenderLayout renders a layout template with the provided data
func (r *Renderer) RenderLayout(ctx context.Context, layoutPath string, data map[string]any, funcs map[string]interface{}) (string, error) {
	// Merge shared data with provided data
	mergedData := make(map[string]any)
	for k, v := range r.data {
		mergedData[k] = v
	}
	for k, v := range data {
		mergedData[k] = v
	}

	// Render the layout with merged data
	var layoutBuf bytes.Buffer
	if err := r.template(mergedData, funcs).Render(ctx, &layoutBuf, layoutPath); err != nil {
		return "", err
	}
	return layoutBuf.String(), nil
}
