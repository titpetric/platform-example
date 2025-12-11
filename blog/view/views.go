package view

import (
	"context"
	"io/fs"

	"github.com/titpetric/platform-example/blog/layout"
)

type Views struct {
	renderer *layout.Renderer
	data     map[string]any
}

func NewViews(root fs.FS) (*Views, error) {
	data := map[string]any{}
	if err := fillTemplateData(&data); err != nil {
		return nil, err
	}

	return &Views{
		renderer: layout.NewRenderer(root, data),
		data:     data,
	}, nil
}

// RenderPage loads a page template, renders it, builds data with "content" key, and calls RenderLayout
func (v *Views) RenderPage(ctx context.Context, pagePath string, templateData map[string]any) (string, error) {
	return v.renderer.RenderPage(ctx, pagePath, templateData, Funcs)
}

// RenderLayout renders the page content within a specified layout using the provided data map
func (v *Views) RenderLayout(ctx context.Context, layoutPath string, data map[string]any) (string, error) {
	return v.renderer.RenderLayout(ctx, layoutPath, data, Funcs)
}
