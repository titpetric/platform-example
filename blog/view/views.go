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

// LoadTemplate loads a template from the filesystem with LESS processor and default funcmap
func (v *Views) LoadTemplate(filename string) (vuego.Template, error) {
	tpl, err := vuego.Load(v.root, filename, vuego.WithLessProcessor())
	if err != nil {
		return nil, err
	}
	return tpl.Funcs(Funcs), nil
}

// RenderPage loads a page template, renders it, builds data with "content" key, and calls RenderLayout
func (v *Views) RenderPage(ctx context.Context, pagePath string, templateData map[string]interface{}) (string, error) {
	// Load the page template
	tpl, err := v.LoadTemplate(pagePath)
	if err != nil {
		return "", err
	}

	// Render the page template
	var pageBuf bytes.Buffer
	if err := tpl.Fill(templateData).Render(ctx, &pageBuf); err != nil {
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
	// Load the layout template
	layout, err := v.LoadTemplate(layoutPath)
	if err != nil {
		return "", err
	}

	// Render the layout with provided data
	var layoutBuf bytes.Buffer
	if err := layout.Fill(data).Render(ctx, &layoutBuf); err != nil {
		return "", err
	}
	return layoutBuf.String(), nil
}
