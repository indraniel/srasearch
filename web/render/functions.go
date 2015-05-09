package render

import (
	"html/template"
	"strings"
	"time"
)

var funcMap template.FuncMap

func init() {
	funcMap = template.FuncMap{
		"add":        Add,
		"mapGet":     MapGet,
		"timeFmt":    TimeFmt,
		"safeHTML":   SafeHTML,
		"shortType":  ShortType,
		"minUint64":  MinUint64,
		"pluralType": PluralType,
	}
}

func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
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
