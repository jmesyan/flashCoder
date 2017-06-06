package home

import (
	"encoding/json"
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/html"
	"flashCoder/app/models"
	"flashCoder/app/operates"
	// "flashCoder/utils"
	// "fmt"
	"net/http"
	"reflect"
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
			task := models.Task.GetTask(tid)
			if task.Tid > 0 {
				taskBehavior := models.Task.GetTaskBehavior(task.Tid, task.Tcate)
				var lastRes interface{}
				for k, v := range taskBehavior {
					bv := models.Behavior.GetBehavior(v.Bid)
					optag := models.Operate.GetOperateTagById(bv.Opid)
					var params []models.OperateParams
					json.Unmarshal([]byte(v.Paramsin), &params)
					var paramsList map[string]string
					paramsList = make(map[string]string)
					for _, param := range params {
						paramsList[param.Name] = param.Value
					}

					if operate, ok := operates.Operates[optag]; ok {
						ber := reflect.ValueOf(operate)
						in := make([]reflect.Value, 2)
						in[0] = reflect.ValueOf(paramsList)
						if k == 0 {
							in[1] = reflect.ValueOf(true)
						} else {
							in[1] = reflect.ValueOf(lastRes)
						}
						last := ber.MethodByName("Execute").Call(in)
						lastRes = last[0].Interface()
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
	if r.Method == "GET" {
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
				data := map[string]interface{}{
					"list": behavior,
				}
				c.View(w, data)
			} else {
				c.Error(w, "行为不存在", "")
			}
		}
	} else {
		c.Error(w, "操作失败", "")
	}
}
