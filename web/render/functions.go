package render

import (
	"html/template"
	"strings"
	"time"
)

var funcMap template.FuncMap

func init() {
	funcMap = template.FuncMap{
		"safeHTML":  SafeHTML,
		"mapGet":    MapGet,
		"timeFmt":   TimeFmt,
		"shortType": ShortType,
	}
}

func ShortType(member string) string {
	abbrev := member[0:2]
	return strings.Title(abbrev)
}

func TimeFmt(ts, fmt string) string {
	t, err := time.Parse("2006-01-02T15:04:05Z", ts)
	if err != nil {
		return ""
	}
	return t.Format(fmt)
}

func SafeHTML(text string) template.HTML {
	return template.HTML(text)
}

func MapGet(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		return val.(string)
	}
	return ""
}
