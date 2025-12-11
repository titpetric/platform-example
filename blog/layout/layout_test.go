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
			Data: []byte(`BASE: [{{ content }}]`),
		},
		"layouts/post.vuego": &fstest.MapFile{
			Data: []byte(`POST: [{{ content }}]`),
		},
		"layouts/index.vuego": &fstest.MapFile{
			Data: []byte(`INDEX: [{{ content }}]`),
		},
		"pages/post.vuego": &fstest.MapFile{
			Data: []byte(`---
layout: post
---
PostContent`),
		},
		"pages/index.vuego": &fstest.MapFile{
			Data: []byte(`---
layout: index
---
IndexContent`),
		},
	}

	sharedData := map[string]any{
		"siteName": "Test Blog",
	}

	renderer := NewRenderer(fsys, sharedData)
	ctx := context.Background()

	t.Run("page content is wrapped in layout", func(t *testing.T) {
		output, err := renderer.RenderPage(ctx, "pages/post.vuego", map[string]any{}, nil)
		assert.NoError(t, err)
		// Should use default base layout since front matter layout detection happens elsewhere
		assert.Contains(t, output, "BASE: [")
		assert.Contains(t, output, "PostContent")
		assert.Contains(t, output, "]")
	})

	t.Run("shared data is available in layout", func(t *testing.T) {
		fsys2 := fstest.MapFS{
			"layouts/base.vuego": &fstest.MapFile{
				Data: []byte(`SITE: {{ siteName }}`),
			},
			"pages/test.vuego": &fstest.MapFile{
				Data: []byte(`Test`),
			},
		}
		renderer := NewRenderer(fsys2, sharedData)
		output, err := renderer.RenderPage(ctx, "pages/test.vuego", map[string]any{}, nil)
		assert.NoError(t, err)
		assert.Contains(t, output, "SITE: Test Blog")
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
