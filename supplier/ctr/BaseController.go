package ctr

import (
	"encoding/json"
	"flashCoder/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

const RESOURCEPATH = "resources/"
const VIEWPATH = RESOURCEPATH + "views/"
const TPLEXT = ".html"

type BaseController struct {
	Mn          string
	Cn          string
	Fn          string
	DefaultView string
}

func (c *BaseController) SetBase(mn, cn, fn string) {
	c.Mn, c.Cn, c.Fn = strings.ToLower(mn), strings.ToLower(cn), strings.ToLower(fn)
	c.DefaultView = c.SetView(c.Mn, c.Cn, c.Fn)
}

func (c *BaseController) SetView(mn, cn, fn string) string {
	var mp, cp, fp string
	if mn == "/" {
		mp = VIEWPATH + "home/"
	} else {
		if len(mn) > 0 {
			mp = VIEWPATH + mn + "/"
		} else {
			mp = ""
		}

	}

	if cn == "/" {
		cp = "index/"
	} else {
		if len(cn) > 0 {
			cp = cn + "/"
		} else {
			cp = ""
		}

	}

	fp = fn
	return mp + cp + fp + TPLEXT
}

func (c *BaseController) View(w http.ResponseWriter, data interface{}) {
	headerView := c.SetView(c.Mn, "public", "header")
	footerView := c.SetView(c.Mn, "public", "footer")
	if c.DefaultView != "" {
		tplName := c.Fn + TPLEXT
		t := template.New(tplName).Funcs(funcMaps)
		t, err := t.ParseFiles(c.DefaultView, headerView, footerView)
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		err = t.Execute(w, data)
		utils.CheckError("error", err)
	} else {
		utils.CheckError("info", "template is empty")
	}
	return
}

func (c *BaseController) Jsonp(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	res, _ := json.Marshal(data)
	fmt.Fprint(w, string(res))
}

func (c *BaseController) Success(w http.ResponseWriter, message, jumpUrl string) {
	c.dispatchJump(w, message, 1, jumpUrl)
	return
}

func (c *BaseController) Error(w http.ResponseWriter, message, jumpUrl string) {
	c.dispatchJump(w, message, 0, jumpUrl)
	return
}

func (c *BaseController) dispatchJump(w http.ResponseWriter, message string, status int, jumpUrl string) {
	data := make(map[string]interface{})
	data["jumpUrl"] = jumpUrl
	data["jumplen"] = len(jumpUrl)
	data["status"] = status
	if status > 0 { //发送成功信息
		data["message"] = message // 提示信息
		data["error"] = ""        // 提示信息
		data["waitSecond"] = 1    // 成功操作后默认停留1秒
	} else {
		data["error"] = message // 提示信息
		data["message"] = ""    // 提示信息
		//发生错误时候默认停留3秒
		data["waitSecond"] = 3 // 成功操作后默认停留3秒
	}
	data["msglen"] = len(data["message"].(string))
	disView := c.SetView("public", "", "dispatch_jump")
	tplName := "dispatch_jump.html"
	t := template.New(tplName).Funcs(funcMaps)
	t, err := t.ParseFiles(disView)
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	err = t.Execute(w, data)
	utils.CheckError("error", err)
	return
}

func (c *BaseController) ParsePage(r *http.Request) int {
	r.ParseForm()
	p := r.Form.Get("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		return 1
	} else {
		return page
	}

}
