package operates

import (
	"context"
	"flashCoder/utils"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type ParseTmpl struct {
	OperatesBase
}

func (op *ParseTmpl) Execute(ctx context.Context) map[string]interface{} {
	select {
	case <-ctx.Done():
		return nil
	default:
		parseParams(op, ctx)
		resolve := make(map[string]interface{})
		itemplate, ok1 := op.currentParams["itemplate"]
		orgin, ok2 := op.currentParams["params"]
		if !ok1 || !ok2 {
			return nil
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
			return nil
		}

		tmpl := template.Must(template.New(name).Parse(itemplate))
		tmpl.Execute(fh, data)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		resolve["content"] = string(content)
		return resolve
	}

}
