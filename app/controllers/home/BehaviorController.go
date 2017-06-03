package home

import (
	"encoding/json"
	"flashCoder/app/behaviors"
	"flashCoder/app/kernel/ctr"
	"flashCoder/app/kernel/html"
	"flashCoder/app/models"
	"flashCoder/utils"
	"fmt"
	"net/http"
	"reflect"
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
			utils.CheckError(err)
			remark := "-"
			if len(r.Form["remark"]) > 0 {
				remark = r.Form["remark"][0]
			}
			//检查行为是否存在

			//添加数据
			models.Behavior.AddBehavior(r.Form["bname"][0], opid, []byte(r.Form["paramsList"][0]), remark)
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
		if len(r.Form["bid"]) > 0 && len(r.Form["bname"]) > 0 && len(r.Form["operate"]) > 0 && len(r.Form["paramsList"]) > 0 {
			bid, err := strconv.Atoi(r.Form["bid"][0])
			if err != nil || bid < 1 {
				c.Error(w, "行为不存在", "")
				return
			}
			opid, err := strconv.Atoi(r.Form["operate"][0])
			utils.CheckError(err)
			remark := "-"
			if len(r.Form["remark"]) > 0 {
				remark = r.Form["remark"][0]
			}
			models.Behavior.UpdateBehavior(bid, r.Form["bname"][0], opid, []byte(r.Form["paramsList"][0]), remark)
			c.Success(w, "行为更新成功", "/behavior/index")
		} else {
			c.Error(w, "数据不能为空", "")
		}
	} else {
		r.ParseForm()
		bid, err := strconv.Atoi(r.Form["bid"][0])
		if err != nil || bid < 1 {
			c.Error(w, "行为不存在", "")
			return
		}
		behavior := models.Behavior.GetBehavior(bid)
		c.View(w, behavior)
	}
}

func (c *BehaviorController) Btest(r *http.Request, w http.ResponseWriter) {
	res := models.Behavior.BehaviorTest()
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
