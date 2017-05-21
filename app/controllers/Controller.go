package controllers

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

var StaticMap map[string]string

type Controller struct {
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sli := strings.Split(r.URL.Path, "/")
	prefix := "/" + sli[1]
	if localdir, ok := StaticMap[prefix]; ok != false {
		file := localdir + r.URL.Path[len(prefix):]
		http.ServeFile(w, r, file)
		return
	}

	mp, cp, fn := c.urlParse(r.URL.Path)
	lfn := strings.ToLower(fn)
	if lfn == "setview" || lfn == "setbase" || lfn == "view" {
		http.Error(w, "can't find method", 500)
		return
	}

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
							mmn := strings.ToLower(res.Type().Method(i).Name)
							if mmn == lfn {
								inb := make([]reflect.Value, 3)
								inb[0], inb[1], inb[2] = reflect.ValueOf(mp), reflect.ValueOf(cp), reflect.ValueOf(fn)
								res.MethodByName("SetBase").Call(inb)
								inf := make([]reflect.Value, 2)
								inf[0], inf[1] = reflect.ValueOf(r), reflect.ValueOf(w)
								f.Call(inf)
								return
							}
						}
					}
				}
			}
		}

	}

	http.NotFound(w, r)
	return
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

func AddstaticMap(webdir, localdir string) {
	StaticMap[webdir] = localdir
}
