package crontab

import (
	"errors"
	"flashCoder/app/jobs"
	"flashCoder/app/models"
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
		jobs.TaskExecute(taskDetail)
		fmt.Println(tid, "执行完成")
	} else {
		fmt.Println(tid, "任务不存在")
	}
}
