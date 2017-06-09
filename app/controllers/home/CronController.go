package home

import (
	"flashCoder/app/kernel/cron"
	"flashCoder/app/kernel/ctr"
	// "math/rand"
	"net/http"
)

type CronController struct {
	ctr.BaseController
}

func (c *CronController) TestCron(r *http.Request, w http.ResponseWriter) {
	var num int64
	num = 1
	task := cron.CronTask{TaskId: num, Crontab: "*/5 * * * * ?"}
	cron.Task <- task
}
