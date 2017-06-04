package models

import (
	"encoding/json"
	"flashCoder/utils"
	"time"
)

type TaskItem struct {
	ItemId   int64
	ItemName string
}

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
	sql := "select count(tid) as count from flash_task where 1"
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

func (m *TaskModel) AddBasicTask(name string, berhaviors []TaskItem) bool {
	tx, err := DB.TransBegin() //使用事务确保mysql数据表类型为Innodb
	if err != nil {
		return false
	}
	//添加一个任务
	sql := "insert into flash_task(tname, tcate, bids, addtime, updtime) values(?, ?, ?, ?, ?)"
	contents := make([]interface{}, 5)
	contents[0] = name
	contents[1] = 1
	json, err := json.Marshal(berhaviors)
	if err != nil {
		return false
	}
	contents[2] = string(json)
	contents[3] = time.Now().Unix()
	contents[4] = contents[3]
	tid, err := DB.TransInsert(tx, sql, contents)
	if err != nil {
		tx.Rollback()
		return false
	}

	//添加任务行为序列
	for k, v := range berhaviors {
		bid := v.ItemId
		border := k
		behavior := Behavior.GetBehavior(bid)
		paramsin := behavior.Paramsdef
		sql = "insert into flash_task_behavior(bid, tid, border, paramsin) values(?, ?, ?, ?)"
		contents = []interface{}{bid, tid, border, paramsin}
		_, err := DB.TransInsert(tx, sql, contents)
		if err != nil {
			tx.Rollback()
			return false
		}
	}
	tx.Commit()
	return true
}

func (m *TaskModel) GetTask(tid int64) FlashTask {
	sql := "select * from flash_task where tid = ?"
	condition := []interface{}{tid}
	result, err := DB.Select(sql, condition)
	utils.CheckError(err)
	var res []FlashTask
	json.Unmarshal([]byte(result), &res)
	return res[0]
}

func (m *TaskModel) GetTaskBehavior(tid int64, tcate uint8) []FlashTaskBehavior {
	var sql string
	if tcate == 1 {
		sql = "select * from flash_task_behavior where tid = ? order by border asc"
	} else {
		sql = "select * from flash_task_behavior where ctid = ? order by torder asc, border asc"
	}
	var condition []interface{}
	condition = make([]interface{}, 1)
	condition[0] = tid
	result, err := DB.Select(sql, condition)
	utils.CheckError(err)
	var res []FlashTaskBehavior
	json.Unmarshal([]byte(result), &res)
	return res
}
