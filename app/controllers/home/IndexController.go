package home

import (
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/html"
	"flashCoder/app/models"
	"net/http"
)

type IndexController struct {
	ctr.BaseController
}

func (c *IndexController) Index(r *http.Request, w http.ResponseWriter) {
	page := c.ParsePage(r)
	pageSize := 10
	list := models.Cron.GetCronList(page, pageSize)
	total := models.Cron.GetCronListCount()
	pages := html.NewPage(page, pageSize, total, "/cron/index")
	data := map[string]interface{}{
		"list": list,
		"page": pages.Show(),
	}
	c.View(w, data)
}
