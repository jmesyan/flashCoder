package models

import (
	"encoding/json"
	"flashCoder/utils"
	// "time"
)

type TaskModel struct {
}

func (m *TaskModel) GetTaskList(tcate, page, pageSize int) []FlashTask {
	sql := "select * from flash_task where 1"
	if tcate > 0 {
		sql += " and tcate= ? "
	}
	sql += " order by addtime desc "
	sql += getPageSql(page, pageSize)
	var condition []interface{}
	if tcate > 0 {
		condition = make([]interface{}, 1)
		condition[0] = tcate
	} else {
		condition = make([]interface{}, 0)
	}
	result, err := DB.Select(sql, condition)
	utils.CheckError(err)
	var res []FlashTask
	json.Unmarshal([]byte(result), &res)
	return res
}

func (m *TaskModel) GetTaskListCount(tcate int) int {
	sql := "select count(tid) as countfrom flash_task where 1"
	if tcate > 0 {
		sql += " and tcate= ? "
	}
	var condition []interface{}
	if tcate > 0 {
		condition = make([]interface{}, 1)
		condition[0] = tcate
	} else {
		condition = make([]interface{}, 0)
	}
	var count int
	res := []interface{}{&count}
	err := DB.SelectOne(sql, condition, res)
	if err != nil {
		return 0
	} else {
		return count
	}
}
