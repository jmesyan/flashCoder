package models

import (
	"encoding/json"
	"flashCoder/utils"
	"time"
)

type BehaviorModel struct {
}

func (m *BehaviorModel) GetBehaviorList(page, pageSize int) []FlashBehavior {
	sql := "select a.*,b.opname from flash_behavior as a inner join flash_operate as b on a.opid=b.opid order by addtime desc"
	sql += getPageSql(page, pageSize)
	condition := make([]interface{}, 0)
	result, err := DB.Select(sql, condition)
	utils.CheckError("error", err)
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
	utils.CheckError("error", err)
	return lastid
}

func (m *BehaviorModel) UpdateBehavior(bid int, bname string, paramsdef []byte, remark string) bool {
	sql := "update flash_behavior set bname=?, paramsdef=?, remark=?, updtime=? where bid=?"
	updtime := time.Now().Unix()
	params := []interface{}{bname, paramsdef, remark, updtime, bid}
	err := DB.Update(sql, params)
	if err != nil {
		return false
	} else {
		return true
	}

}

func (m *BehaviorModel) GetBehavior(bid int64) FlashBehavior {
	sql := "select a.*,b.opname from flash_behavior as a inner join flash_operate as b on a.opid=b.opid where a.bid = ?"
	condition := []interface{}{bid}
	result, err := DB.Select(sql, condition)
	utils.CheckError("error", err)
	var res []FlashBehavior
	json.Unmarshal([]byte(result), &res)
	if len(res) > 0 {
		return res[0]
	} else {
		return FlashBehavior{}
	}
}

func (m *BehaviorModel) IsExistBehavior(bid int, bname string) bool {
	sql := "select count(opid) as count from flash_behavior where bname = ? and bid <> ?"
	condition := make([]interface{}, 2)
	condition[0] = bname
	condition[1] = bid
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

func (m *BehaviorModel) DeleteBehavior(bid int64) bool {
	sql := "delete from flash_behavior where bid = ?"
	params := []interface{}{bid}
	err := DB.Update(sql, params)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (m *BehaviorModel) GetOperateCountInBehavior(opid int64) int {
	sql := "select count(bid) as count from flash_behavior where opid=?"
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
