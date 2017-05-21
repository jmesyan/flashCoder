package home

import (
	ke "flashCoder/app/kernel"
	"fmt"
	"net/http"
)

type Reponse http.ResponseWriter
type Request *http.Request

type IndexController struct {
	ke.BaseController
}

func (c *IndexController) Index(r Request, w Reponse) {
	fmt.Println(r.Method)
	fmt.Println("this is home indexController welcome")
	w.Header().Set("Server", "golang")
	fmt.Fprintf(w, "jmesyan,%q", "location in index")
}

func (c *IndexController) Welcome(r Request, w Reponse) {
	fmt.Fprintf(w, "jmesyan,%q", "location in welcome")
}
