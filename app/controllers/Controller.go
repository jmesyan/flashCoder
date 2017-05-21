package controllers

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type Controller struct {
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	match := false
	mp, cp, fn := c.urlParse(r.URL.Path)
	if _, ok := Router[mp]; ok {
		mr := reflect.ValueOf(Router[mp])
		if mr.Kind() == reflect.Map {
			mma := mr.Interface().(map[string]interface{})
			if _, ok = mma[cp]; ok {
				ce := mma[cp]
				res := reflect.ValueOf(ce)
				if res.Kind() == reflect.Ptr {
					cen := res.NumMethod()
					if cen > 0 {
						for i := 0; i < cen; i++ {
							f := res.Method(i)
							if strings.ToLower(res.Type().Method(i).Name) == strings.ToLower(fn) {
								match = true
								in := make([]reflect.Value, 2)
								in[0], in[1] = reflect.ValueOf(r), reflect.ValueOf(w)
								f.Call(in)
							}
						}
					}
				}
			}
		}

	}

	if !match {
		http.NotFound(w, r)
	}
}

func (c *Controller) urlParse(up string) (mp, cp, fn string) {
	ua := strings.Replace(up, "\\", "/", -1)
	rep := regexp.MustCompile("/{2,}").ReplaceAllString(ua, "/")
	sp := strings.Split(strings.Trim(rep, "/"), "/")
	lsp := len(sp)
	if lsp == 1 && sp[0] != "" {
		mp, cp, fn = "/", "/", sp[0]
		return
	} else if lsp == 2 {
		mp, cp, fn = "/", sp[0], sp[1]
		return
	} else if lsp == 3 {
		mp, cp, fn = sp[0], sp[1], sp[2]
		return
	}

	mp, cp, fn = "/", "/", "index"
	return
}
