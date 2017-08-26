package api

import (
	"encoding/json"
	"flashCoder/app/jobs"
	"flashCoder/app/models"
	"flashCoder/supplier/ctr"
	"flashCoder/supplier/html"
	"net/http"
	"strconv"
)

type TaskController struct {
	ctr.BaseController
}

func (c *TaskController) List(r *http.Request, w http.ResponseWriter) {
	page := c.ParsePage(r)
	pageSize := 10
	list := models.Task.GetTaskList(0, page, pageSize)
	total := models.Task.GetTaskListCount(0)
	pages := html.NewPage(page, pageSize, total, "/task/list")
	data := map[string]interface{}{
		"list": list,
		"page": pages.Show(),
	}
	c.Jsonp(w, data)
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
			data := map[string]interface{}{
				"ret": 1,
				"msg": "请选择数据",
			}
			c.Jsonp(w, data)
		} else {
			if taskType == "basicTask" {
				res := models.Task.AddBasicTask(taskName, itemList)
				if res {
					data := map[string]interface{}{
						"ret": 0,
						"msg": "添加任务成功",
					}
					c.Jsonp(w, data)
				} else {
					data := map[string]interface{}{
						"ret": 1,
						"msg": "添加任务失败,请重试",
					}
					c.Jsonp(w, data)
				}
			} else {
				res := models.Task.AddCompositeTask(taskName, itemList)
				if res {
					data := map[string]interface{}{
						"ret": 0,
						"msg": "添加任务成功",
					}
					c.Jsonp(w, data)
				} else {
					data := map[string]interface{}{
						"ret": 1,
						"msg": "添加任务失败,请重试",
					}
					c.Jsonp(w, data)
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

		c.Jsonp(w, data)
	}

}

func (c *TaskController) TaskExecute(r *http.Request, w http.ResponseWriter) {
	if r.Method == "GET" {
		r.ParseForm()
		var tid int64
		tidstr := r.Form["tid"][0]
		if len(tidstr) <= 0 {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "任务信息不全",
			}
			c.Jsonp(w, data)
		} else {
			td, err := strconv.Atoi(tidstr)
			if err != nil {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "任务id应该是整数",
				}
				c.Jsonp(w, data)
				return
			}
			tid = int64(td)
			taskDetail := models.Task.GetTask(tid)
			if taskDetail.Tid > 0 {
				jobs.TaskExecute(taskDetail)
				data := map[string]interface{}{
					"ret": 0,
					"msg": "任务执行完成",
				}
				c.Jsonp(w, data)
			} else {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "任务不存在",
				}
				c.Jsonp(w, data)
			}
		}
	}
}

func (c *TaskController) EditTaskBehaviors(r *http.Request, w http.ResponseWriter) {
	if r.Method == "GET" {
		r.ParseForm()
		var tid int64
		tidstr := r.Form["tid"][0]
		if len(tidstr) <= 0 {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "任务信息不全",
			}
			c.Jsonp(w, data)
		} else {
			td, err := strconv.Atoi(tidstr)
			if err != nil {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "任务id应该是整数",
				}
				c.Jsonp(w, data)
			}
			tid = int64(td)
			task := models.Task.GetTask(tid)
			if task.Tid > 0 {
				taskBehavior := models.Task.GetTaskBehavior(task.Tid, task.Tcate)
				data := map[string]interface{}{
					"list": taskBehavior,
				}
				c.Jsonp(w, data)
			} else {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "任务不存在",
				}
				c.Jsonp(w, data)
			}
		}
	} else {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "操作失败",
		}
		c.Jsonp(w, data)
	}

}

func (c *TaskController) TaskBehaviorParams(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		var tbid int64
		tbidstr := r.Form["tbid"][0]
		if len(tbidstr) <= 0 {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "任务信息不全",
			}
			c.Jsonp(w, data)
		} else {
			td, err := strconv.Atoi(tbidstr)
			if err != nil {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "任务id应该是整数",
				}
				c.Jsonp(w, data)
				return
			}
			tbid = int64(td)
			paramsin := r.Form["paramsList"][0]
			models.Task.UpdateTaskBehaviorById(tbid, []byte(paramsin))
			behavior := models.Task.GetTaskBehaviorById(tbid)
			var tid int
			if behavior.Ctid > 0 {
				tid = int(behavior.Ctid)
			} else {
				tid = int(behavior.Tid)
			}
			data := map[string]interface{}{
				"ret": 0,
				"msg": "行为更新成功",
				"tid":tid,
			}
			c.Jsonp(w, data)
		}
	} else {
		r.ParseForm()
		var tbid int64
		tbidstr := r.Form["tbid"][0]
		if len(tbidstr) <= 0 {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "任务信息不全",
			}
			c.Jsonp(w, data)
		} else {
			td, err := strconv.Atoi(tbidstr)
			if err != nil {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "任务id应该是整数",
				}
				c.Jsonp(w, data)
				return
			}
			tbid = int64(td)
			behavior := models.Task.GetTaskBehaviorById(tbid)
			if behavior.Bid > 0 {
				baseBehavior := models.Behavior.GetBehavior(behavior.Bid)
				data := map[string]interface{}{
					"ret": 0,
					"base":   baseBehavior,
					"params": behavior.Paramsin,
					"tbid":   tbid,
				}
				c.Jsonp(w, data)
			} else {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "行为不存在",
				}
				c.Jsonp(w, data)
			}
		}
	}
}

func (c *TaskController) Delete(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	var tid int64
	tint, _ := strconv.Atoi(r.Form["tid"][0])
	tid = int64(tint)
	if tid < 1 {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "参数不正确",
		}
		c.Jsonp(w, data)
	}
	//检查行为是否已被任务使用
	count := models.Cron.GetTaskCountInCron(tid)
	if count > 0 {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "该任务存在于计划当中",
		}
		c.Jsonp(w, data)
	} else {
		if models.Task.DeleteTask(tid) {
			data := map[string]interface{}{
				"ret": 0,
				"msg": "删除成功",
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

}
