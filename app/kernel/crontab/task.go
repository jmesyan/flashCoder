package crontab

import (
	"context"
	"encoding/json"
	"errors"
	"flashCoder/app/models"
	"flashCoder/app/operates"
	"fmt"
	"github.com/robfig/cron"
	"strconv"
	"strings"
	"time"
)

var Task chan CronTask
var CronTasks []CronTask
var Handler *cron.Cron

func Check(cron models.FlashCron) bool {
	monthNow := int(time.Now().Month())
	monthCron, err := getMaxDate(cron.Month)
	if err == nil && monthNow > monthCron {
		return false
	}
	return true
}

func getMaxDate(date string) (int, error) {
	var err error
	if strings.Contains(date, "*") || strings.Contains(date, "?") {
		return 0, errors.New("the date is not num")
	}

	if strings.Contains(date, "-") {
		result := strings.Split(date, "-")
		res, err := strconv.Atoi(result[1])
		return res, err
	}

	if strings.Contains(date, ",") {
		result := strings.Split(date, ",")
		res, err := strconv.Atoi(result[1])
		return res, err
	}

	res, err := strconv.Atoi(date)
	return res, err
}

func Watching() {
	Task = make(chan CronTask)
	Reload()
	// go func() {
	// 	for {
	// 		select {
	// 		case task := <-Task:
	// 			Handler.AddJob(task.Crontab, task)
	// 		}
	// 	}
	// }()
}

func Reload() {
	if Handler != nil {
		Handler.Stop()
	}
	Handler = cron.New()
	Handler.Start()
	tasks := models.Cron.GetValidList()
	for _, v := range tasks {
		if Check(v) {
			sep := " "
			crontab := v.Second + sep + v.Minute + sep + v.Hour + sep + v.Day + sep + v.Month + sep + v.Week
			fmt.Println(crontab)
			task := CronTask{TaskId: v.Tid, Crontab: crontab}
			Handler.AddJob(task.Crontab, task)
		}
	}
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
		fmt.Println(tid, "执行完成")
	} else {
		fmt.Println(tid, "任务不存在")
	}
}
