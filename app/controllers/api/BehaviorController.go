package api

import (
	"flashCoder/app/models"
	"flashCoder/supplier/ctr"
	"flashCoder/supplier/html"
	"flashCoder/utils"
	"net/http"
	"strconv"
)

type BehaviorController struct {
	ctr.BaseController
}

func (c *BehaviorController) List(r *http.Request, w http.ResponseWriter) {
	page := c.ParsePage(r)
	pageSize := 10
	list := models.Behavior.GetBehaviorList(page, pageSize)
	total := models.Behavior.GetBehaviorListCount()
	pages := html.NewPage(page, pageSize, total, "/behavior/list")
	data := map[string]interface{}{
		"list": list,
		"page": pages.Show(),
	}
	c.Jsonp(w, data)
}

func (c *BehaviorController) Add(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		if len(r.Form["bname"]) > 0 && len(r.Form["operate"]) > 0 && len(r.Form["paramsList"]) > 0 {
			opid, err := strconv.Atoi(r.Form["operate"][0])
			utils.CheckError("error", err)
			remark := "-"
			if len(r.Form["remark"]) > 0 {
				remark = r.Form["remark"][0]
			}
			bname := r.Form["bname"][0]
			//检查行为名称是否存在
			if models.Behavior.IsExistBehavior(0, bname) {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "行为名已存在，请调整",
				}
				c.Jsonp(w, data)
			}
			//添加数据
			models.Behavior.AddBehavior(bname, opid, []byte(r.Form["paramsList"][0]), remark)
			data := map[string]interface{}{
				"ret": 0,
				"msg": "行为添加成功",
			}
			c.Jsonp(w, data)
		} else {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "数据不能为空",
			}
			c.Jsonp(w, data)
		}
	}

}

func (c *BehaviorController) Update(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		if len(r.Form["bid"]) > 0 && len(r.Form["bname"]) > 0 && len(r.Form["paramsList"]) > 0 {
			bid, err := strconv.Atoi(r.Form["bid"][0])
			if err != nil || bid < 1 {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "行为不存在",
				}
				c.Jsonp(w, data)
			}
			remark := "-"
			if len(r.Form["remark"]) > 0 {
				remark = r.Form["remark"][0]
			}
			bname := r.Form["bname"][0]
			//检查行为名称是否存在
			if models.Behavior.IsExistBehavior(bid, bname) {
				data := map[string]interface{}{
					"ret": 1,
					"msg": "行为不存在",
				}
				c.Jsonp(w, data)
			}
			models.Behavior.UpdateBehavior(bid, bname, []byte(r.Form["paramsList"][0]), remark)
			data := map[string]interface{}{
				"ret": 0,
				"msg": "行为更新成功",
			}
			c.Jsonp(w, data)
		} else {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "数据不能为空",
			}
			c.Jsonp(w, data)
		}
	} else {
		r.ParseForm()
		var bid int64
		id, err := strconv.Atoi(r.Form["bid"][0])
		if err != nil || id < 1 {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "行为不存在",
			}
			c.Jsonp(w, data)
		}
		bid = int64(id)
		behavior := models.Behavior.GetBehavior(bid)
		c.Jsonp(w, behavior)
	}
}

func (c *BehaviorController) Delete(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	var bid int64
	bint, _ := strconv.Atoi(r.Form["bid"][0])
	bid = int64(bint)
	if bid < 1 {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "参数不正确",
		}
		c.Jsonp(w, data)
	}
	//检查行为是否已被任务使用
	count := models.Task.GetBehaviorCountInTask(bid)
	if count > 0 {
		data := map[string]interface{}{
			"ret": 1,
			"msg": "该行为已被任务使用",
		}
		c.Jsonp(w, data)
	} else {
		if models.Behavior.DeleteBehavior(bid) {
			data := map[string]interface{}{
				"ret": 0,
				"msg": "删除成功",
			}
			c.Jsonp(w, data)
		} else {
			data := map[string]interface{}{
				"ret": 1,
				"msg": "删除失败",
			}
			c.Jsonp(w, data)
		}
	}

}
