package template

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
	"time"

	"github.com/titpetric/vuego"
)

type Views struct {
	vue  *vuego.Vue
	data map[string]any
}

func NewViews(root fs.FS) (*Views, error) {
	vue := vuego.NewVue(root)
	vue.RegisterNodeProcessor(vuego.NewLessProcessor(root))

	funcMap := vue.DefaultFuncMap()
	funcMap["postDate"] = func(v any) any {
		layoutStr := "2006/01/02 15:04"

		switch t := v.(type) {
		case time.Time:
			return t.Format(layoutStr)
		case string:
			// Try to parse as RFC3339 or Unix timestamp
			if parsed, err := time.Parse(time.RFC3339, t); err == nil {
				return parsed.Format(layoutStr)
			}
		}
		return v
	}
	funcMap["readingTime"] = func(content string) string {
		// average reading speed ~200 words/minute
		words := len(strings.Fields(content))
		minutes := words / 200
		if words%200 != 0 {
			minutes++
		}

		if minutes <= 2 {
			return "a few minutes"
		}
		return fmt.Sprintf("%d minutes", minutes)
	}

	funcMap["metaTitle"] = func(title string) string {
		return title
	}
	funcMap["metaDescription"] = func(description string) string {
		return description
	}
	funcMap["metaOGImage"] = func(in string) string {
		return in
	}
	funcMap["getCss"] = func(string) string {
		return ""
	}
	funcMap["getJs"] = func(string) string {
		return ""
	}
	funcMap["json"] = func(in any) (string, error) {
		b, err := json.MarshalIndent(in, "", "  ")
		return string(b), err
	}
	vue.Funcs(funcMap)

	data := map[string]any{}
	if err := fillTemplateData(&data); err != nil {
		return nil, err
	}

	return &Views{
		vue:  vue,
		data: data,
	}, nil
}
