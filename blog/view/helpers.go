package view

import (
	"strings"
)

// StripFrontMatter removes YAML front matter from content and returns just the body
func StripFrontMatter(content string) string {
	// Check if content starts with ---
	if !strings.HasPrefix(content, "---") {
		// No front matter, return as-is
		return content
	}

	// Find the closing ---
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		// No closing ---, return as-is
		return content
	}

	// Return content without front matter (parts[2])
	return parts[2]
}
