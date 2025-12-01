package template

import (
	"context"
	"fmt"
	"html"
	"time"

	"github.com/titpetric/platform-example/blog/model"
)

// AtomFeed generates an Atom XML feed for articles
func (v *Views) AtomFeed(ctx context.Context, articles []model.Article) (string, error) {
	var meta map[string]any
	var author map[string]any

	if m, ok := v.data["meta"].(map[string]any); ok {
		meta = m
	}
	if m, ok := meta["author"].(map[string]any); ok {
		author = m
	}

	// Find the most recent article date
	var newestDate time.Time
	if len(articles) > 0 {
		newestDate = articles[0].Date
		for _, a := range articles {
			if a.Date.After(newestDate) {
				newestDate = a.Date
			}
		}
	} else {
		newestDate = time.Now()
	}

	var language string
	if lang, ok := meta["language"].(string); ok {
		language = lang
	}

	xml := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xml:base="%s">
  <title>%s</title>
  <subtitle>Blogging general thoughts and rambles, code snippets, and front-end web dev discoveries</subtitle>
  <link href="%s/feed.xml" rel="self"/>
  <link href="%s"/>
  <updated>%s</updated>
  <id>%s</id>
  <author>
    <name>%s</name>
    <email>%s</email>
  </author>
`, meta["url"], escapeXML(meta["title"]), meta["url"], meta["url"], newestDate.Format(time.RFC3339), meta["url"], escapeXML(author["name"]), author["email"])

	// Add entries for each article
	for _, article := range articles {
		// Strip front matter from content
		contentWithoutFrontMatter := StripFrontMatter(article.Content)

		entryXML := fmt.Sprintf(`  <entry>
    <title>%s</title>
    <link href="%s/blog/%s"/>
    <updated>%s</updated>
    <id>%s/blog/%s</id>
    <content xml:lang="%s" type="html">%s</content>
  </entry>
`, escapeXML(article.Title), meta["url"], article.Slug, article.Date.Format(time.RFC3339), meta["url"], article.Slug, language, escapeXML(contentWithoutFrontMatter))

		xml += entryXML
	}

	xml += `</feed>`

	return xml, nil
}

// escapeXML escapes special XML characters
func escapeXML(s any) string {
	return html.EscapeString(s.(string))
}

// stringReplace is a simple string replacement helper
func stringReplace(s, old, new string) string {
	result := ""
	i := 0
	for i < len(s) {
		idx := -1
		for j := i; j <= len(s)-len(old); j++ {
			if s[j:j+len(old)] == old {
				idx = j
				break
			}
		}
		if idx == -1 {
			result += s[i:]
			break
		}
		result += s[i:idx] + new
		i = idx + len(old)
	}
	return result
}
