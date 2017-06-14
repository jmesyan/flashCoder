package tester

import (
	"fmt"
	html "html/template"
	"os"
	"testing"
	text "text/template"
)

var textTmpl = text.Must(text.New("").Parse(`<script type="text/babel">
		var list = {{.list}}
		console.log(list)
		</script>`))

var htmlTmpl = html.Must(html.New("").Parse(`<script type="text/babel">
		var list = {{.list}}
		alert("hello");
		</script>`))

type Grade struct {
	Gid   int64
	Gname string
}

func TestTemplate(t *testing.T) {

	list := []Grade{
		{1, "Name1"},
		{2, "Name2"},
	}

	data := map[string]interface{}{
		"list": list,
	}

	fmt.Printf("text/template:\n\n")
	textTmpl.Execute(os.Stdout, data)

	fmt.Printf("\n\nhtml/template:\n\n")
	htmlTmpl.Execute(os.Stdout, data)
}
