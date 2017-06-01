package home

import (
	"encoding/json"
	"flashCoder/app/behaviors"
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/db"
	"flashCoder/utils"
	"fmt"
	"net/http"
	// "os"
	"reflect"
	"time"
)

type BehaviorController struct {
	ctr.BaseController
}

func (c *BehaviorController) Index(r *http.Request, w http.ResponseWriter) {
	sql := "select * from flash_behavior order by updtime desc"
	condition := make([]interface{}, 0)
	result, err := c.DB.Select(sql, condition)
	utils.CheckError(err)
	var res []flashdb.FlashBehavior
	json.Unmarshal([]byte(result), &res)
	data := map[string]interface{}{
		"list": res,
	}
	c.View(w, data)
}

func (c *BehaviorController) Add(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		data := make([]interface{}, 4)
		if len(r.Form["name"]) > 0 {
			data[0] = r.Form["name"][0]
			if len(r.Form["paramName[]"]) > 0 && len(r.Form["paramValue[]"]) == len(r.Form["paramName[]"]) {
				paramsin := make(map[string]string)
				for k, v := range r.Form["paramName[]"] {
					paramsin[v] = r.Form["paramValue[]"][k]
				}
				jsonData, err := json.Marshal(paramsin)
				utils.CheckError(err)
				data[1] = jsonData
			} else {
				data[1] = nil
			}
			data[2] = time.Now().Unix()
			data[3] = data[2]
			sql := "insert into flash_behavior(name, paramsin, addtime, updtime) values(?, ?, ?, ?)"
			lastid, err := c.DB.Insert(sql, data)
			utils.CheckError(err)
			fmt.Fprintln(w, "保存数据成功", lastid)
		} else {
			fmt.Fprintln(w, "行为名称不能为空")
		}
	} else {
		c.View(w, nil)
	}

}

func (c *BehaviorController) Btest(r *http.Request, w http.ResponseWriter) {
	sql := "select * from flash_behavior where bid in(2, 3, 4) order by bid asc"
	condition := make([]interface{}, 0)
	result, err := c.DB.Select(sql, condition)
	utils.CheckError(err)
	var res []flashdb.FlashBehavior
	json.Unmarshal([]byte(result), &res)
	var lastRes interface{}
	fmt.Println(res)
	for k, v := range res {
		var params map[string]string
		json.Unmarshal([]byte(v.Paramsdef), &params)
		if behavior, ok := behaviors.Behaviors[v.Bname]; ok {
			ber := reflect.ValueOf(behavior)
			in := make([]reflect.Value, 2)
			in[0] = reflect.ValueOf(params)
			if k == 0 {
				in[1] = reflect.ValueOf(0)
			} else {
				in[1] = reflect.ValueOf(lastRes)
			}
			last := ber.MethodByName("Execute").Call(in)
			lastRes = last[0].Interface()
		}

	}
}
