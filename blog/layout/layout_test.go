package layout

import (
	"context"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestRenderPageWithLayout(t *testing.T) {
	// Create a test filesystem with layouts and pages  
	fsys := fstest.MapFS{
		"layouts/base.vuego": &fstest.MapFile{
			Data: []byte(`<p>Layout: base</p>
{{ content }}`),
		},
		"layouts/post.vuego": &fstest.MapFile{
			Data: []byte(`<p>Layout: post</p>
{{ content }}`),
		},
		"layouts/index.vuego": &fstest.MapFile{
			Data: []byte(`<p>Layout: index</p>
{{ content }}`),
		},
		"pages/post.vuego": &fstest.MapFile{
			Data: []byte(`Post Content`),
		},
		"pages/index.vuego": &fstest.MapFile{
			Data: []byte(`Index Content`),
		},
	}

	sharedData := map[string]any{
		"siteName": "Test Blog",
	}

	renderer := NewRenderer(fsys, sharedData)
	ctx := context.Background()

	t.Run("page content is wrapped in default base layout", func(t *testing.T) {
		output, err := renderer.RenderPage(ctx, "pages/post.vuego", map[string]any{}, nil)
		assert.NoError(t, err)
		assert.Contains(t, output, "<p>Layout: base</p>")
		assert.Contains(t, output, "Post Content")
	})

	t.Run("custom post layout can be rendered directly", func(t *testing.T) {
		// Verify that the post layout template has the expected content
		output, err := renderer.RenderLayout(ctx, "layouts/post.vuego", map[string]any{
			"content": "Post Content",
		}, nil)
		assert.NoError(t, err)
		assert.Contains(t, output, "<p>Layout: post</p>")
		assert.Contains(t, output, "Post Content")
	})

	t.Run("custom index layout can be rendered directly", func(t *testing.T) {
		// Verify that the index layout template has the expected content
		output, err := renderer.RenderLayout(ctx, "layouts/index.vuego", map[string]any{
			"content": "Index Content",
		}, nil)
		assert.NoError(t, err)
		assert.Contains(t, output, "<p>Layout: index</p>")
		assert.Contains(t, output, "Index Content")
	})

	t.Run("shared data is available in layout", func(t *testing.T) {
		fsys2 := fstest.MapFS{
			"layouts/base.vuego": &fstest.MapFile{
				Data: []byte(`<header>{{ siteName }}</header>
{{ content }}`),
			},
			"pages/test.vuego": &fstest.MapFile{
				Data: []byte(`Test Content`),
			},
		}
		renderer := NewRenderer(fsys2, sharedData)
		output, err := renderer.RenderPage(ctx, "pages/test.vuego", map[string]any{}, nil)
		assert.NoError(t, err)
		assert.Contains(t, output, "<header>Test Blog</header>")
		assert.Contains(t, output, "Test Content")
	})
}

func TestRenderLayout(t *testing.T) {
	fsys := fstest.MapFS{
		"layouts/base.vuego": &fstest.MapFile{
			Data: []byte(`WRAP: [{{ content }}]`),
		},
	}

	renderer := NewRenderer(fsys, map[string]any{})
	ctx := context.Background()

	output, err := renderer.RenderLayout(ctx, "layouts/base.vuego", map[string]any{
		"content": "INNER",
	}, nil)

	assert.NoError(t, err)
	assert.Contains(t, output, "WRAP: [INNER]")
}
