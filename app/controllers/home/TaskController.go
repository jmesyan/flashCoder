package home

import (
	"encoding/json"
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
		taskName := r.Form["taskName"][0]
		taskType := r.Form["taskType"][0]
		var itemList []models.TaskItem
		if taskType == "basicTask" {
			taskList := r.Form["basicTaskBehaviors"][0]
			json.Unmarshal([]byte(taskList), &itemList)
		} else {
			taskList := r.Form["compositeTaskBasics"][0]
			json.Unmarshal([]byte(taskList), &itemList)
		}
		if len(itemList) <= 0 || len(taskName) <= 0 {
			c.Error(w, "请选择数据", "")
		} else {
			if taskType == "basicTask" {
				res := models.Task.AddBasicTask(taskName, itemList)
				if res {
					c.Success(w, "添加任务成功", "/task/index")
				} else {
					c.Error(w, "添加任务失败,请重试", "")
				}
			}
		}
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
