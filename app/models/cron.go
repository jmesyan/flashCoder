package models

import (
	"encoding/json"
	"flashCoder/utils"
	"time"
)

type CronModel struct {
}

func (m *CronModel) GetCron(crid int64) FlashCron {
	sql := "select * from flash_cron where crid = ?"
	condition := []interface{}{crid}
	result, err := DB.Select(sql, condition)
	utils.LogError("error", err)
	var res []FlashCron
	json.Unmarshal([]byte(result), &res)
	if len(res) > 0 {
		return res[0]
	} else {
		return FlashCron{}
	}
}

func (m *CronModel) UpdateCronState(crid int64, state uint8) bool {
	sql := "update flash_cron set state=? where crid=?"
	params := []interface{}{state, crid}
	err := DB.Update(sql, params)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (m *CronModel) GetValidList() []FlashCron {
	sql := "select * from flash_cron where state = 0 order by crid desc"
	condition := make([]interface{}, 0)
	result, err := DB.Select(sql, condition)
	utils.LogError("error", err)
	var res []FlashCron
	json.Unmarshal([]byte(result), &res)
	return res
}

func (m *CronModel) GetCronList(page, pageSize int) []FlashCron {
	sql := "select * from flash_cron order by crid desc	"
	sql += getPageSql(page, pageSize)
	condition := make([]interface{}, 0)
	result, err := DB.Select(sql, condition)
	utils.LogError("error", err)
	var res []FlashCron
	json.Unmarshal([]byte(result), &res)
	return res
}

func (m *CronModel) GetCronListCount() int {
	sql := "select count(crid) as count from flash_cron"
	condition := make([]interface{}, 0)
	var count int
	res := []interface{}{&count}
	err := DB.SelectOne(sql, condition, res)
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (m *CronModel) AddCron(tid int64, second, minute, hour, day, month, week string) int64 {
	sql := "insert into flash_cron(second, minute, hour, day, month, week,tid,addtime) values(?,?,?,?,?,?,?,?)"
	addtime := time.Now().Unix()
	contents := []interface{}{second, minute, hour, day, month, week, tid, addtime}
	lastid, err := DB.Insert(sql, contents)
	utils.LogError("error", err)
	return lastid
}

func (m *CronModel) UpdateCron(crid int64, second, minute, hour, day, month, week string) bool {
	sql := "update flash_cron set second = ?,minute=?, hour=?,day=?,month=?,week=? where crid = ?"
	params := []interface{}{second, minute, hour, day, month, week, crid}
	err := DB.Update(sql, params)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (m *CronModel) DeleteCron(crid int64) bool {
	sql := "delete from flash_cron where crid = ?"
	params := []interface{}{crid}
	err := DB.Update(sql, params)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (m *CronModel) GetTaskCountInCron(tid int64) int {
	sql := "select count(crid) as count from flash_cron where tid=?"
	condition := []interface{}{tid}
	var count int
	res := []interface{}{&count}
	err := DB.SelectOne(sql, condition, res)
	if err != nil {
		return 0
	} else {
		return count
	}
}
