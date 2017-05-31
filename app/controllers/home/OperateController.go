package home

import (
	"encoding/json"
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/db"
	"flashCoder/utils"
	"fmt"
	"net/http"
	// "reflect"
	// "time"
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

	} else {
		c.View(w, nil)
	}
}

func (c *OperateController) JsonOperateList(r *http.Request, w http.ResponseWriter) {
	sql := "select optag as value, opname as name from flash_operate order by addtime desc"
	condition := make([]interface{}, 0)
	operates, err := c.DB.Select(sql, condition)
	utils.CheckError(err)
	fmt.Fprint(w, string(operates))

}
