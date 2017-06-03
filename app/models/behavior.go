package models

import (
	"encoding/json"
	"flashCoder/utils"
	"time"
)

type BehaviorModel struct {
}

func (m *BehaviorModel) GetBehaviorList(page, pageSize int) []FlashBehavior {
	sql := "select * from flash_behavior order by addtime desc"
	sql += getPageSql(page, pageSize)
	condition := make([]interface{}, 0)
	result, err := DB.Select(sql, condition)
	utils.CheckError(err)
	var res []FlashBehavior
	json.Unmarshal([]byte(result), &res)
	return res
}

func (m *BehaviorModel) GetBehaviorListCount() int {
	sql := "select count(opid) as count from flash_behavior"
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

func (m *BehaviorModel) AddBehavior(bname string, opid int, paramsdef []byte, remark string) int64 {
	sql := "insert into flash_behavior(opid, bname, paramsdef, remark, addtime, updtime) values(?, ?, ?, ?, ?, ?)"
	data := make([]interface{}, 6)
	data[0] = opid
	data[1] = bname
	data[2] = paramsdef
	data[3] = remark
	data[4] = time.Now().Unix()
	data[5] = data[4]
	lastid, err := DB.Insert(sql, data)
	utils.CheckError(err)
	return lastid
}

func (m *BehaviorModel) UpdateBehavior(bid int, bname string, opid int, paramsdef []byte, remark string) bool {
	sql := "update flash_behavior set opid=?, bname=?, paramsdef=?, remark=?, updtime=? where bid=?"
	updtime := time.Now().Unix()
	params := []interface{}{opid, bname, paramsdef, remark, updtime, bid}
	err := DB.Update(sql, params)
	if err != nil {
		return false
	} else {
		return true
	}

}

func (m *BehaviorModel) GetBehavior(bid int) FlashBehavior {
	sql := "select * from flash_behavior where bid = ?"
	condition := []interface{}{bid}
	result, err := DB.Select(sql, condition)
	utils.CheckError(err)
	var res []FlashBehavior
	json.Unmarshal([]byte(result), &res)
	return res[0]
}

func (m *BehaviorModel) BehaviorTest() []FlashBehavior {
	sql := "select * from flash_behavior where bid in(2, 3, 4) order by bid asc"
	condition := make([]interface{}, 0)
	result, err := DB.Select(sql, condition)
	utils.CheckError(err)
	var res []FlashBehavior
	json.Unmarshal([]byte(result), &res)
	return res
}
