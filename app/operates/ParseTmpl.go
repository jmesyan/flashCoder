package operates

import (
	"flashCoder/utils"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type ParseTmpl struct {
}

func (op *ParseTmpl) Execute(params map[string]string, lastRes interface{}) interface{} {
	itemplate, ok1 := params["itemplate"]
	orgin, ok2 := params["params"]
	if !ok1 || !ok2 {
		return false
	}
	data := make(map[string]string)
	tmp := strings.Split(orgin, ";")
	for _, v := range tmp {
		tmp2 := strings.Split(v, "=")
		data[tmp2[0]] = tmp2[1]
	}
	name := utils.MD5("parse")
	path := "./" + name
	os.Remove(path)
	defer os.Remove(path)
	m := new(sync.RWMutex)
	m.Lock()
	defer m.Unlock()
	fh, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	defer fh.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}

	tmpl := template.Must(template.New(name).Parse(itemplate))
	tmpl.Execute(fh, data)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return string(contents)
}
