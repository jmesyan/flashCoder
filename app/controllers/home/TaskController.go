package home

import (
	"encoding/json"
	"flashCoder/app/models"
	"flashCoder/app/operates"
	"flashCoder/supplier/ctr"
	"flashCoder/supplier/html"
	// "flashCoder/utils"
	// "fmt"
	"context"
	"net/http"
	// "reflect"
	"strconv"
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
			} else {
				res := models.Task.AddCompositeTask(taskName, itemList)
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

func (c *TaskController) TaskExecute(r *http.Request, w http.ResponseWriter) {
	if r.Method == "GET" {
		r.ParseForm()
		var tid int64
		tidstr := r.Form["tid"][0]
		if len(tidstr) <= 0 {
			c.Error(w, "任务信息不全", "")
		} else {
			td, err := strconv.Atoi(tidstr)
			if err != nil {
				c.Error(w, "任务id应该是整数", "")
				return
			}
			tid = int64(td)
			taskDetail := models.Task.GetTask(tid)
			if taskDetail.Tid > 0 {
				taskBehavior := models.Task.GetTaskBehavior(taskDetail.Tid, taskDetail.Tcate)
				task, _ := context.WithCancel(context.Background())
				global := context.WithValue(task, operates.ParamsGlobal, map[string]string{})
				resolve := map[string]string{}
				for _, v := range taskBehavior {
					bv := models.Behavior.GetBehavior(v.Bid)
					optag := models.Operate.GetOperateTagById(bv.Opid)
					var params []models.OperateParams
					json.Unmarshal([]byte(v.Paramsin), &params)
					if optag == "ParamsGlobal" {
						if pa := global.Value(operates.ParamsGlobal).(map[string]string); pa != nil {
							for _, param := range params {
								pa[param.Name] = param.Value
							}
							global = context.WithValue(task, operates.ParamsGlobal, pa)

						}
					} else {
						current := make(map[string]string)
						for _, param := range params {
							current[param.Name] = param.Value
						}
						val := map[string]map[string]string{
							"current": current,
							"resolve": resolve,
						}
						curres := context.WithValue(global, operates.ParamsCurRes, val)
						if operate, ok := operates.Operates[optag]; ok {
							resolve = operate.Execute(curres)
						}
					}

				}
				c.Success(w, "任务执行完成", "")
			} else {
				c.Error(w, "任务不存在", "")
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
			c.Error(w, "任务信息不全", "")
		} else {
			td, err := strconv.Atoi(tidstr)
			if err != nil {
				c.Error(w, "任务id应该是整数", "")
				return
			}
			tid = int64(td)
			task := models.Task.GetTask(tid)
			if task.Tid > 0 {
				taskBehavior := models.Task.GetTaskBehavior(task.Tid, task.Tcate)
				data := map[string]interface{}{
					"list": taskBehavior,
				}
				c.View(w, data)
			} else {
				c.Error(w, "任务不存在", "")
			}
		}
	} else {
		c.Error(w, "操作失败", "")
	}

}

func (c *TaskController) TaskBehaviorParams(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		var tbid int64
		tbidstr := r.Form["tbid"][0]
		if len(tbidstr) <= 0 {
			c.Error(w, "任务信息不全", "")
		} else {
			td, err := strconv.Atoi(tbidstr)
			if err != nil {
				c.Error(w, "任务id应该是整数", "")
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
			c.Success(w, "行为更新成功", "/task/editTaskBehaviors?tid="+strconv.Itoa(tid))
		}
	} else {
		r.ParseForm()
		var tbid int64
		tbidstr := r.Form["tbid"][0]
		if len(tbidstr) <= 0 {
			c.Error(w, "任务信息不全", "")
		} else {
			td, err := strconv.Atoi(tbidstr)
			if err != nil {
				c.Error(w, "任务id应该是整数", "")
				return
			}
			tbid = int64(td)
			behavior := models.Task.GetTaskBehaviorById(tbid)
			if behavior.Bid > 0 {
				baseBehavior := models.Behavior.GetBehavior(behavior.Bid)
				data := map[string]interface{}{
					"base":   baseBehavior,
					"params": behavior.Paramsin,
					"tbid":   tbid,
				}
				c.View(w, data)
			} else {
				c.Error(w, "行为不存在", "")
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
		c.Error(w, "参数不正确", "")
		return
	}
	//检查行为是否已被任务使用
	count := models.Cron.GetTaskCountInCron(tid)
	if count > 0 {
		c.Error(w, "该任务存在于计划当中", "")
	} else {
		if models.Task.DeleteTask(tid) {
			c.Success(w, "删除成功", "")
		} else {
			c.Error(w, "删除失败", "")
		}
	}

}
