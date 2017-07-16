package api

import (
	"flashCoder/app/models"
	"flashCoder/supplier/ctr"
	"flashCoder/supplier/html"
	"flashCoder/utils"
	"fmt"
	"html/template"
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

func (c *IndexController) Client(r *http.Request, w http.ResponseWriter) {
	tplView := "client/dist/index.html"
	tplName := "index.html"
	t := template.New(tplName)
	t, err := t.ParseFiles(tplView)
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	err = t.Execute(w, nil)
	utils.CheckError("error", err)
}
