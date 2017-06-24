package tester

import (
	"encoding/json"
	"flashCoder/flashdb"
	"flashCoder/utils"
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	db := flashdb.SetDbHandler(flashdb.DRMySQL, "root:@/flashCoder")
	//事务开始
	tx, err := db.TransBegin() //使用事务确保mysql数据表类型为Innodb
	utils.CheckError("error", err)

	defer tx.Rollback()

	//选择一条数据
	sql := "SELECT behavior FROM flash_behavior WHERE name = ?"
	condition := []interface{}{13}
	var squareNum int
	res := []interface{}{&squareNum}
	err = db.SelectOne(sql, condition, res)
	utils.CheckError("error", err)
	fmt.Println(squareNum)
	//选择多条数据
	sql = "SELECT Name FROM flash_behavior"
	condition = make([]interface{}, 0)
	abc, err := db.Select(sql, condition)
	utils.CheckError("error", err)
	var p []flashdb.FlashBehavior
	json.Unmarshal([]byte(abc), &p)
	fmt.Println(p[13].Behavior, p[13].Name)

	//更新数据
	sql = "Update flash_coding set code=? where tid=?"
	params := []interface{}{"welcome to shanghai", 30}
	err = db.TransUpdate(tx, sql, params)
	utils.CheckError("error", err)
	//插入数据
	sql = "INSERT INTO flash_coding VALUES(?,?,?)"
	contents := []interface{}{25, 77, "hello SS"}
	lastid, err := db.TransInsert(tx, sql, contents)
	utils.CheckError("error", err)
	fmt.Println(lastid)
	tx.Commit()

	db.Close()
}
