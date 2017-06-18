package operates

import (
	"flashCoder/app/kernel/file"
	"fmt"
	"io/ioutil"
	// "reflect"
	"context"
	"strconv"
	"strings"
)

type FileEdit struct {
	globalParams  map[string]string
	currentParams map[string]string
	resolveParams map[string]string
}

func (op *FileEdit) setParams(t ParamsType, val map[string]string) {
	switch t {
	case ParamsGlobal:
		op.globalParams = val
	case ParamsCurrent:
		op.currentParams = val
	case ParamsResolve:
		op.resolveParams = val
	}
}

func (op *FileEdit) Execute(ctx context.Context) map[string]string {
	select {
	case <-ctx.Done():
		return nil
	default:
		parseParams(op, ctx)
		resolve := make(map[string]string)
		if _, ok := op.currentParams["params"]; !ok {
			return nil
		}
		origin := op.currentParams["params"]
		data := make(map[string]string)
		tmp := strings.Split(origin, ";")
		for _, v := range tmp {
			tmp2 := strings.Split(v, "=")
			data[tmp2[0]] = tmp2[1]
		}

		path := op.globalParams["path"]
		funcName := data["funcName"]
		offset, _ := strconv.Atoi(data["offset"])
		isBegin := false
		if strings.Trim(data["isBegin"], " ") == "1" {
			isBegin = true
		}

		dat, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
			return nil
		}
		tmpFile := strings.Split(string(dat), "\n")
		content := op.resolveParams["content"]
		fh := new(file.FlashFile)
		res := fh.AddFuncContent(tmpFile, funcName, content, isBegin, offset)
		if res != nil {
			newContent := strings.Join(res, "\n")
			go func() {
				ioutil.WriteFile(path, []byte(newContent), 0)
			}()
			resolve["result"] = "success"
			return resolve
		} else {
			fmt.Print("failed")
			return nil
		}
	}

}
