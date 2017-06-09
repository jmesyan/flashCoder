package cron

import (
	"encoding/json"
	"flashCoder/app/models"
	"flashCoder/app/operates"
	"fmt"
	"github.com/robfig/cron"
	"reflect"
)

var Task chan CronTask
var CronTasks []CronTask
var Handler *cron.Cron

type CronWatcher struct {
}

func (c *CronWatcher) Watching() {
	Task = make(chan CronTask)
	Handler = cron.New()
	Handler.Start()
	go func() {
		for {
			select {
			case task := <-Task:
				Handler.AddJob(task.Crontab, task)
			}
		}
	}()
}

type CronTask struct {
	TaskId  int64
	Crontab string
}

func (c CronTask) Run() {
	tid := c.TaskId
	c.excute(tid)
}

func (c CronTask) excute(tid int64) {
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
		fmt.Println(tid, "执行完成")
	} else {
		fmt.Println(tid, "任务不存在")
	}
}
