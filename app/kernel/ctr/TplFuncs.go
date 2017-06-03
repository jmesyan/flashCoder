package ctr

import (
	"html/template"
	"time"
)

var funcMaps = template.FuncMap{
	"timeformat": func(timestamp int64) string {
		return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	},
	"str2html": func(raw string) template.HTML {
		return template.HTML(raw)
	},
}
