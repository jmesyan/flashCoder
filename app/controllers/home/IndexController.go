package home

import (
	"flashCoder/app/kernel/ctr"
	"net/http"
)

type IndexController struct {
	ctr.BaseController
}

func (c *IndexController) Index(r *http.Request, w http.ResponseWriter) {
	data := map[string]interface{}{
		"hello": "world",
		"good":  123,
		"yes":   "yes is good",
		"Title": "这是header测试",
	}
	c.View(w, data)
}
