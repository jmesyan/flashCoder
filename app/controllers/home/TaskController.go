package home

import (
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/html"
	"flashCoder/app/models"
	// "flashCoder/utils"
	// "fmt"
	"net/http"
	// "strconv"
)

type TaskController struct {
	ctr.BaseController
}

func (c *TaskController) Index(r *http.Request, w http.ResponseWriter) {
	page := c.ParsePage(r)
	pageSize := 10
	list := models.Task.GetTaskList(0, page, pageSize)
	total := models.Task.GetTaskListCount(0)
	pages := html.NewPage(page, pageSize, total, "/task/index")
	data := map[string]interface{}{
		"list": list,
		"page": pages.Show(),
	}
	c.View(w, data)
}

func (c *TaskController) Add(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()

	} else {
		page := c.ParsePage(r)
		pageSize := 10
		//基础任务获取行为列表
		behaviorList := models.Behavior.GetBehaviorList(page, pageSize)
		//复合任务获取基础任务列表
		basicTaskList := models.Task.GetTaskList(1, page, pageSize)

		data := map[string]interface{}{
			"behaviorList":  behaviorList,
			"basicTaskList": basicTaskList,
		}

		c.View(w, data)
	}

}
