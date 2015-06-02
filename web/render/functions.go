package render

import (
	"github.com/arschles/go-bindata-html-template"

	"time"
)

var funcMap template.FuncMap

func init() {
	funcMap = template.FuncMap{
		"add":          Add,
		"mapGet":       MapGet,
		"timeFmt":      TimeFmt,
		"incrementIdx": IncrementIndex,
		"shortType":    ShortType,
		"minUint64":    MinUint64,
		"pluralType":   PluralType,
		//		"safeHTML":     SafeHTML,
	}
}

func IncrementIndex(counter int, index uint64) int {
	val := counter + int(index)
	return val
}

func PluralType(member string) string {
	plural := member
	switch {
	case member == "study":
		plural = "Studies"
	case member == "sample":
		plural = "Samples"
	case member == "experiment":
		plural = "Experiments"
	case member == "run":
		plural = "Runs"
	case member == "analysis":
		plural = "Analyses"
	case member == "submission":
		plural = "Submissions"
	}
	return plural
}

func MinUint64(args ...uint64) uint64 {
	var min uint64 = 100000000
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func ShortType(member string) string {
	abbrev := "-"
	switch {
	case member == "study":
		abbrev = "St"
	case member == "sample":
		abbrev = "Sa"
	case member == "experiment":
		abbrev = "E"
	case member == "run":
		abbrev = "R"
	case member == "analysis":
		abbrev = "A"
	case member == "submission":
		abbrev = "Su"
	}
	return abbrev
}

func TimeFmt(ts, fmt string) string {
	t, err := time.Parse("2006-01-02T15:04:05Z", ts)
	if err != nil {
		return ""
	}
	return t.Format(fmt)
}

// func SafeHTML(text string) template.HTML {
// 	return template.HTML(text)
// }

func MapGet(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		return val.(string)
	}
	return ""
}
