package home

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

func (c *BehaviorController) Index(r *http.Request, w http.ResponseWriter) {
	page := c.ParsePage(r)
	pageSize := 10
	list := models.Behavior.GetBehaviorList(page, pageSize)
	total := models.Behavior.GetBehaviorListCount()
	pages := html.NewPage(page, pageSize, total, "/behavior/index")
	data := map[string]interface{}{
		"list": list,
		"page": pages.Show(),
	}
	c.View(w, data)
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
				c.Error(w, "行为名已存在，请调整", "")
				return
			}
			//添加数据
			models.Behavior.AddBehavior(bname, opid, []byte(r.Form["paramsList"][0]), remark)
			c.Success(w, "行为添加成功", "/behavior/index")
		} else {
			c.Error(w, "数据不能为空", "")
		}
	} else {
		c.View(w, nil)
	}

}

func (c *BehaviorController) Update(r *http.Request, w http.ResponseWriter) {
	if r.Method == "POST" {
		r.ParseForm()
		if len(r.Form["bid"]) > 0 && len(r.Form["bname"]) > 0 && len(r.Form["paramsList"]) > 0 {
			bid, err := strconv.Atoi(r.Form["bid"][0])
			if err != nil || bid < 1 {
				c.Error(w, "行为不存在", "")
				return
			}
			remark := "-"
			if len(r.Form["remark"]) > 0 {
				remark = r.Form["remark"][0]
			}
			bname := r.Form["bname"][0]
			//检查行为名称是否存在
			if models.Behavior.IsExistBehavior(bid, bname) {
				c.Error(w, "行为名已存在，请调整", "")
				return
			}
			models.Behavior.UpdateBehavior(bid, bname, []byte(r.Form["paramsList"][0]), remark)
			// c.Success(w, "行为更新成功", "/behavior/index")
		} else {
			c.Error(w, "数据不能为空", "")
		}
	} else {
		r.ParseForm()
		var bid int64
		id, err := strconv.Atoi(r.Form["bid"][0])
		if err != nil || id < 1 {
			c.Error(w, "行为不存在", "")
			return
		}
		bid = int64(id)
		behavior := models.Behavior.GetBehavior(bid)
		c.View(w, behavior)
	}
}

func (c *BehaviorController) Delete(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	var bid int64
	bint, _ := strconv.Atoi(r.Form["bid"][0])
	bid = int64(bint)
	if bid < 1 {
		c.Error(w, "参数不正确", "")
		return
	}
	//检查行为是否已被任务使用
	count := models.Task.GetBehaviorCountInTask(bid)
	if count > 0 {
		c.Error(w, "该行为已被任务使用", "")
	} else {
		if models.Behavior.DeleteBehavior(bid) {
			c.Success(w, "删除成功", "")
		} else {
			c.Error(w, "删除失败", "")
		}
	}

}
