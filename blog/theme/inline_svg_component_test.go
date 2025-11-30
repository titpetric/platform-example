package theme_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/titpetric/vuego"
)

func TestInlineSVGComponent(t *testing.T) {
	validSVG := "assets/icons/avatar.svg"
	nonexistentSVG := "assets/icons/nonexistent.svg"

	tests := []struct {
		name        string
		src         string
		expectError bool
	}{
		{
			name:        "load valid svg",
			src:         validSVG,
			expectError: false,
		},
		{
			name:        "load nonexistent svg",
			src:         nonexistentSVG,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use DirFS from the workspace root
			vue := vuego.NewVue(os.DirFS("."))

			// Create template data
			templateData := map[string]interface{}{
				"src": tt.src,
			}

			var buf bytes.Buffer
			err := vue.RenderFragment(&buf, "components/inline-svg.vuego", templateData)
			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			result := buf.String()
			t.Log(result)

			require.Contains(t, result, "<svg")
			require.Contains(t, result, "</svg>")
			require.True(t, !strings.Contains(result, "<svg></svg>"))
		})
	}
}
