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
	c.DefaultView = c.SetView(c.Mn, c.Cn, c.Fn)
}

func (c *BaseController) SetView(mn, cn, fn string) string {
	var mp, cp, fp string
	if mn == "/" {
		mp = VIEWPATH + "home/"
	} else {
		if len(mn) > 0 {
			mp = VIEWPATH + mn + "/"
		} else {
			mp = ""
		}

	}

	if cn == "/" {
		cp = "index/"
	} else {
		if len(cn) > 0 {
			cp = cn + "/"
		} else {
			cp = ""
		}

	}

	fp = fn
	return mp + cp + fp + TPLEXT
}

func (c *BaseController) View(w Reponse, data interface{}) {
	headerView := c.SetView(c.Mn, "public", "header")
	footerView := c.SetView(c.Mn, "public", "footer")
	if c.DefaultView != "" {
		t, err := template.ParseFiles(c.DefaultView, headerView, footerView)
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		t.Execute(w, data)
	} else {
		fmt.Fprint(w, "template is empty")
	}

}
