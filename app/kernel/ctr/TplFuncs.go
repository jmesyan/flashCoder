package ctr

import (
	"html/template"
	"time"
)

var funcMaps = template.FuncMap{
	"timeformat": func(timestamp int64) string {
		return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	},
}
