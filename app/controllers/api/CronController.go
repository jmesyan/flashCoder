package api

import (
	"encoding/json"
	"flashCoder/app/models"
	"flashCoder/supplier/crontab"
	"flashCoder/supplier/ctr"
	"flashCoder/supplier/html"
	"net/http"
	"strconv"
	// "fmt"
)

type CronController struct {
	ctr.BaseController
}

func (c *CronController) List(r *http.Request, w http.ResponseWriter) {
	page := 1
	pageSize := 10
	list := models.Cron.GetCronList(page, pageSize)
	total := models.Cron.GetCronListCount()
	pages := html.NewPage(page, pageSize, total, "/cron/list")
	data := map[string]interface{}{
		"list": list,
		"page": pages.Show(),
	}
	c.Jsonp(w, data)
}

func (c *CronController) Add(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		second := r.Form["Second"][0]
		minute := r.Form["Minute"][0]
		hour := r.Form["Hour"][0]
		day := r.Form["Day"][0]
		month := r.Form["Month"][0]
		week := r.Form["Week"][0]
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
			data := map[string]interface{}{
				"ret": 0,
				"msg": "定时任務添加成功",
			}
			c.Jsonp(w, data)
		} else {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "请选择数据",
			}
			c.Jsonp(w, data)
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

		c.Jsonp(w, data)
	}
}

func (c *CronController) Update(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	var crid int64
	crint, _ := strconv.Atoi(r.Form["crid"][0])
	crid = int64(crint)
	if crid < 1 {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "参数不正确",
		}
		c.Jsonp(w, data)
	}

	if r.Method == "POST" {
		second := r.Form["Second"][0]
		minute := r.Form["Minute"][0]
		hour := r.Form["Hour"][0]
		day := r.Form["Day"][0]
		month := r.Form["Month"][0]
		week := r.Form["Week"][0]
		if len(second) > 0 && len(minute) > 0 && len(hour) > 0 && len(day) > 0 && len(month) > 0 && len(week) > 0 {
			if models.Cron.UpdateCron(crid, second, minute, hour, day, month, week) {
				crontab.Reload()
				data := map[string]interface{}{
					"ret": 0,
					"msg": "更新数据成功",
				}
				c.Jsonp(w, data)
			} else {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "更新数据失败",
				}
				c.Jsonp(w, data)
			}
		} else {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "请填写数据",
			}
			c.Jsonp(w, data)
		}
	} else {
		cron := models.Cron.GetCron(crid)
		task := models.Task.GetTask(cron.Tid)
		data := map[string]interface{}{
			"ret":  0,
			"cron": cron,
			"task": task,
		}

		c.Jsonp(w, data)

	}
}

func (c *CronController) UpdateState(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	var crid int64
	crint, _ := strconv.Atoi(r.Form["crid"][0])
	crid = int64(crint)
	if crid < 1 {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "參數不正確",
		}
		c.Jsonp(w, data)
		return
	}
	cron := models.Cron.GetCron(crid)
	var state uint8
	var stateDesc string
	if cron.State == 1 {
		state = 0
		stateDesc = "开启"
	} else {
		state = 1
		stateDesc = "关闭"
	}

	if models.Cron.UpdateCronState(crid, state) {
		crontab.Reload()
		message := stateDesc + "成功！"
		data := map[string]interface{}{
			"ret": 0,
			"msg": message,
		}
		c.Jsonp(w, data)
	} else {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "操作失败",
		}
		c.Jsonp(w, data)
	}
}

func (c *CronController) Delete(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	var crid int64
	crint, _ := strconv.Atoi(r.Form["crid"][0])
	crid = int64(crint)
	if crid < 1 {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "參數不正確",
		}
		c.Jsonp(w, data)
		return
	}
	if models.Cron.DeleteCron(crid) {
		crontab.Reload()
		data := map[string]interface{}{
			"ret": 0,
			"msg": "刪除成功",
		}
		c.Jsonp(w, data)
	} else {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "删除失败",
		}
		c.Jsonp(w, data)
	}

}
