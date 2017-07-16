package api

import (
	"flashCoder/app/models"
	"flashCoder/supplier/ctr"
	"flashCoder/supplier/html"
	"fmt"
	"net/http"
	"strconv"
)

type OperateController struct {
	ctr.BaseController
}

func (c *OperateController) Index(r *http.Request, w http.ResponseWriter) {
	page := c.ParsePage(r)
	pageSize := 10
	list := models.Operate.GetOperateList(page, pageSize)
	total := models.Operate.GetOperateListCount()
	pages := html.NewPage(page, pageSize, total, "/operate/index")
	data := map[string]interface{}{
		"list": list,
		"page": pages.Show(),
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
			remark := "-"
			if len(r.Form["remark"]) > 0 {
				remark = r.Form["remark"][0]
			}
			opname := r.Form["opname"][0]
			optag := r.Form["optag"][0]
			//检查操作名称是否存在
			if models.Operate.IsExistOperate(0, optag) {
				c.Error(w, "操作标识已存在，请调整", "")
				return
			}
			//加入数据
			models.Operate.AddOperate(opname, optag, remark)
			c.Success(w, "保存数据成功", "/operate/index")
			return
		}
	} else {
		c.View(w, nil)
	}
}

func (c *OperateController) Delete(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	var opid int64
	oint, _ := strconv.Atoi(r.Form["opid"][0])
	opid = int64(oint)
	count := models.Operate.GetOperateCount(opid)
	if opid < 1 || count < 1 {
		c.Error(w, "不存在该操作", "")
		return
	} else {
		//检查操作是否已被行为使用
		count := models.Behavior.GetOperateCountInBehavior(opid)
		if count > 0 {
			c.Error(w, "该操作已被行为使用", "")
		} else {
			if models.Operate.DeleteOperate(opid) {
				c.Success(w, "删除成功", "")
			} else {
				c.Error(w, "删除失败", "")
			}
		}
	}
}

func (c *OperateController) JsonOperateList(r *http.Request, w http.ResponseWriter) {
	operates := models.Operate.GetOperateSelectItems()
	fmt.Fprint(w, string(operates))

}
