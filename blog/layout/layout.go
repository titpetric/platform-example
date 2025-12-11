package layout

import (
	"bytes"
	"context"
	"io"
	"io/fs"
	"log"

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
func (r *Renderer) template(data map[string]any) vuego.Template {
	tpl := vuego.NewFS(r.root, vuego.WithLessProcessor())
	return tpl.Funcs(Funcs).Fill(data)
}

// Render loads a template, and if the template contains "layout" in the metadata, it will
// load another template from layouts/%s.vuego; Layouts can be chained so one layout can
// again trigger another layout, like `blog.vuego -> layouts/post.vuego -> layouts/base.vuego`.
func (r *Renderer) Render(ctx context.Context, w io.Writer, filename string, data map[string]any) error {
	tpl := r.template(data)
	for {
		var buf bytes.Buffer
		if err := tpl.Load(filename).Render(ctx, &buf); err != nil {
			return err
		}

		log.Printf("Render: %s %q", filename, tpl.Get("layout"))

		content := buf.String()
		data["content"] = content

		layout := tpl.Get("layout")
		if layout == "" {
			filename = "layouts/base.vuego"
			break
		}

		delete(data, "layout")
		tpl = r.template(data)
		filename = "layouts/" + layout + ".vuego"
		continue
	}
	log.Printf("Render: %s %q", filename, tpl.Get("layout"))

	var buf bytes.Buffer
	if err := tpl.Load(filename).Render(ctx, &buf); err != nil {
		return err
	}
	_, err := io.Copy(w, &buf)
	return err
}
