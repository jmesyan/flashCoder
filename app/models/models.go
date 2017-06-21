package models

import (
	"flashCoder/app/kernel/db"
	// "os"
	// "os/exec"
	// "path/filepath"
	"strconv"
)

var DB flashdb.FlashDB
var Operate *OperateModel
var Behavior *BehaviorModel
var Task *TaskModel
var Cron *CronModel

func init() {
	dbType, connstr := flashdb.DBConnection("default")
	DB = flashdb.SetDbHandler(dbType, connstr)
	Operate = new(OperateModel)
	Behavior = new(BehaviorModel)
	Task = new(TaskModel)
	Cron = new(CronModel)
}

func getPageSql(page, pageSize int) string {
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		sql := (" limit " + strconv.Itoa(offset) + ", " + strconv.Itoa(pageSize))
		return sql
	} else {
		return ""
	}
}
