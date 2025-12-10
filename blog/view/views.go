package view

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"

	"github.com/titpetric/vuego"
)

type Views struct {
	root fs.FS
	data map[string]any
}

func NewViews(root fs.FS) (*Views, error) {
	data := map[string]any{}
	if err := fillTemplateData(&data); err != nil {
		return nil, err
	}

	return &Views{
		root: root,
		data: data,
	}, nil
}

func (v *Views) template(data map[string]any) vuego.Template {
	tpl := vuego.Load(v.root, vuego.WithLessProcessor())
	return tpl.Funcs(Funcs).Fill(data)
}

// RenderPage loads a page template, renders it, builds data with "content" key, and calls RenderLayout
func (v *Views) RenderPage(ctx context.Context, pagePath string, templateData map[string]any) (string, error) {
	tpl := v.template(templateData)

	// Render the page template
	var pageBuf bytes.Buffer
	if err := tpl.Render(ctx, &pageBuf, pagePath); err != nil {
		return "", err
	}

	templateData["content"] = pageBuf.String()

	// Get layout name from template metadata
	layoutName := "layouts/base.vuego"
	if layout := tpl.GetString("layout"); layout != "" {
		layoutName = fmt.Sprintf("layouts/%s.vuego", layout)
	}

	// Render with layout
	return v.RenderLayout(ctx, layoutName, templateData)
}

// RenderLayout renders the page content within a specified layout using the provided data map
func (v *Views) RenderLayout(ctx context.Context, layoutPath string, data map[string]any) (string, error) {
	// Render the layout with provided data
	var layoutBuf bytes.Buffer
	if err := v.template(data).Render(ctx, &layoutBuf, layoutPath); err != nil {
		return "", err
	}
	return layoutBuf.String(), nil
}
