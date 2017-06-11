package home

import (
	"encoding/json"
	"flashCoder/app/kernel/crontab"
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/html"
	"flashCoder/app/models"
	"net/http"
	// "time"
)

type CronController struct {
	ctr.BaseController
}

func (c *CronController) Index(r *http.Request, w http.ResponseWriter) {
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

func (c *CronController) Add(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		second := r.Form["second"][0]
		minute := r.Form["minute"][0]
		hour := r.Form["hour"][0]
		day := r.Form["day"][0]
		month := r.Form["month"][0]
		week := r.Form["week"][0]
		taskType := r.Form["taskType"][0]
		var itemList []models.TaskItem
		if taskType == "basicTask" {
			taskList := r.Form["basicTask"][0]
			json.Unmarshal([]byte(taskList), &itemList)
		} else {
			taskList := r.Form["compositeTask"][0]
			json.Unmarshal([]byte(taskList), &itemList)
		}
		if len(second) > 0 && len(minute) > 0 && len(hour) > 0 && len(day) > 0 && len(month) > 0 && len(week) > 0 && len(itemList) > 0 {
			taskId := itemList[0].ItemId
			models.Cron.AddCron(taskId, second, minute, hour, day, month, week)
			crontab.Reload()
			c.Success(w, "定时添加成功", "/cron/index")
		} else {
			c.Error(w, "请选择数据", "")
		}
	} else {
		page := c.ParsePage(r)
		pageSize := 10
		//基础任务
		basicTaskList := models.Task.GetTaskList(1, page, pageSize)
		//复合任务
		compositeTaskList := models.Task.GetTaskList(2, page, pageSize)

		data := map[string]interface{}{
			"basicTaskList":     basicTaskList,
			"compositeTaskList": compositeTaskList,
		}

		c.View(w, data)
	}
}
