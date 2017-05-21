package kernel

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const RESOURCEPATH = "resources/"
const VIEWPATH = RESOURCEPATH + "views/"
const TPLEXT = ".html"

type Reponse http.ResponseWriter

type BaseController struct {
	Mn          string
	Cn          string
	Fn          string
	DefaultView string
}

func (c *BaseController) SetBase(mn, cn, fn string) {
	c.Mn, c.Cn, c.Fn = strings.ToLower(mn), strings.ToLower(cn), strings.ToLower(fn)
	c.SetView(c.Mn, c.Cn, c.Fn)
}

func (c *BaseController) SetView(mn, cn, fn string) {
	var mp, cp, fp string
	if mn == "/" {
		mp = VIEWPATH + "home/"
	} else {
		mp = VIEWPATH + mn + "/"
	}

	if cn == "/" {
		cp = "index/"
	} else {
		cp = cn + "/"
	}

	fp = fn
	c.DefaultView = mp + cp + fp + TPLEXT

}

func (c *BaseController) View(w Reponse, data interface{}) {
	if c.DefaultView != "" {
		t, err := template.ParseFiles(c.DefaultView)
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		t.Execute(w, data)
	} else {
		fmt.Fprint(w, "template is empty")
	}

}
