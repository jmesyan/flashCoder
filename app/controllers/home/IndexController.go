package home

import (
	ke "flashCoder/app/kernel"
	// "fmt"
	"net/http"
)

type Reponse http.ResponseWriter
type Request *http.Request

type IndexController struct {
	ke.BaseController
}

func (c *IndexController) Index(r Request, w Reponse) {
	data := map[string]interface{}{
		"hello": "world",
		"good":  123,
		"yes":   "yes is good",
	}
	c.View(w, data)
}
