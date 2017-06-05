package models

import (
	"encoding/json"
	"flashCoder/utils"
	"time"
)

type OperateParams struct {
	Type  int
	Name  string
	Value string
}

type OperateModel struct {
}

func (m *OperateModel) GetOperateList(page, pageSize int) []FlashOperate {
	sql := "select * from flash_operate order by addtime desc"
	sql += getPageSql(page, pageSize)
	condition := make([]interface{}, 0)
	result, err := DB.Select(sql, condition)
	utils.CheckError(err)
	var res []FlashOperate
	json.Unmarshal([]byte(result), &res)
	return res
}

func (m *OperateModel) GetOperateListCount() int {
	sql := "select count(opid) as count from flash_operate"
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

func (m *OperateModel) GetOperateCount(opid int) int {
	sql := "select count(opid) as count from flash_operate where opid = ?"
	condition := []interface{}{opid}
	var count int
	res := []interface{}{&count}
	err := DB.SelectOne(sql, condition, res)
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (m *OperateModel) DeleteOperate(opid int) bool {
	sql := "delete from flash_operate where opid = ?"
	params := []interface{}{opid}
	err := DB.Update(sql, params)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (m *OperateModel) GetOperateSelectItems() string {
	sql := "select opid as value, opname as name from flash_operate order by addtime desc"
	condition := make([]interface{}, 0)
	operates, err := DB.Select(sql, condition)
	if err != nil {
		return ""
	} else {
		return string(operates)
	}
}

func (m *OperateModel) AddOperate(opname, optag, remark string) int64 {
	sql := "insert into flash_operate(opname, optag, remark, addtime) values (?, ?, ?, ?) "
	data := make([]interface{}, 4)
	data[0] = opname
	data[1] = optag
	data[2] = remark
	data[3] = time.Now().Unix()
	lastid, err := DB.Insert(sql, data)
	utils.CheckError(err)
	return lastid
}

func (m *OperateModel) GetOperateTagById(opid int64) string {
	sql := "select optag from flash_operate where opid = ?"
	condition := []interface{}{opid}
	var optag string
	res := []interface{}{&optag}
	err := DB.SelectOne(sql, condition, res)
	if err != nil {
		return ""
	} else {
		return optag
	}
}

func (m *OperateModel) IsExistOperate(opid int, optag string) bool {
	sql := "select count(opid) as count from flash_operate where optag = ? and opid <> ?"
	condition := make([]interface{}, 2)
	condition[0] = optag
	condition[1] = opid
	var count int
	res := []interface{}{&count}
	err := DB.SelectOne(sql, condition, res)
	if err != nil {
		return false
	} else {
		if count > 0 {
			return true
		} else {
			return false
		}
	}
}

func (m *OperateModel) IsOperateUsed(opid int) bool {
	sql := "select count(tbid) as count from flash_task_behavior as a left join flash_behavior as b on a.bid = b.bid where b.opid=? "
	condition := make([]interface{}, 1)
	condition[0] = opid
	var count int
	res := []interface{}{&count}
	err := DB.SelectOne(sql, condition, res)
	if err != nil {
		return false
	} else {
		if count > 0 {
			return true
		} else {
			return false
		}
	}
}
