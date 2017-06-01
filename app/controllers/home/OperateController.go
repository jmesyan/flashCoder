package home

import (
	"encoding/json"
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/db"
	"flashCoder/utils"
	"fmt"
	"net/http"
	// "reflect"
	"strconv"
	"time"
)

type OperateController struct {
	ctr.BaseController
}

func (c *OperateController) Index(r *http.Request, w http.ResponseWriter) {
	sql := "select * from flash_operate order by addtime desc"
	condition := make([]interface{}, 0)
	result, err := c.DB.Select(sql, condition)
	utils.CheckError(err)
	var res []flashdb.FlashOperate
	json.Unmarshal([]byte(result), &res)
	data := map[string]interface{}{
		"list": res,
	}
	c.View(w, data)
}

func (c *OperateController) Add(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["opname"][0] == "" || r.Form["optag"][0] == "" {
			c.Error(w, "操作名称或者标识不能为空", "")
			return
		} else {
			//检查是否存在标识

			//加入数据
			sql := "insert into flash_operate(opname, optag, remark, addtime) values (?, ?, ?, ?) "
			data := make([]interface{}, 4)
			data[0] = r.Form["opname"][0]
			data[1] = r.Form["optag"][0]
			data[2] = r.Form["remark"][0]
			data[3] = time.Now().Unix()
			_, err := c.DB.Insert(sql, data)
			utils.CheckError(err)
			c.Success(w, "保存数据成功", "")
			return
		}
	} else {
		c.View(w, nil)
	}
}

func (c *OperateController) Delete(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	opid, err := strconv.Atoi(r.Form["opid"][0])
	utils.CheckError(err)
	sql := "select count(opid) as count from flash_operate where opid = ?"
	condition := []interface{}{opid}
	var count int
	res := []interface{}{&count}
	err = c.DB.SelectOne(sql, condition, res)
	utils.CheckError(err)
	if opid < 1 || count < 1 {
		c.Error(w, "不存在该操作", "")
		return
	} else {
		//删除前检查数据是否被使用

		//删除数据
		sql := "delete from flash_operate where opid = ?"
		params := []interface{}{opid}
		err = c.DB.Update(sql, params)
		utils.CheckError(err)
		c.Success(w, "删除数据成功", "")
		return
	}
}

func (c *OperateController) JsonOperateList(r *http.Request, w http.ResponseWriter) {
	sql := "select optag as value, opname as name from flash_operate order by addtime desc"
	condition := make([]interface{}, 0)
	operates, err := c.DB.Select(sql, condition)
	utils.CheckError(err)
	fmt.Fprint(w, string(operates))

}
