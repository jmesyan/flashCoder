package operates

import (
	"flashCoder/supplier/file"
	"fmt"
	"io/ioutil"
	// "reflect"
	"context"
	"strconv"
	"strings"
)

type FileEdit struct {
	OperatesBase
}

func (op *FileEdit) Execute(ctx context.Context) map[string]interface{} {
	select {
	case <-ctx.Done():
		return nil
	default:
		parseParams(op, ctx)
		resolve := make(map[string]interface{})
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
		content := op.resolveParams["content"].(string)
		fh := new(file.FlashFile)
		res := fh.AddFuncContent(tmpFile, funcName, content, isBegin, offset)
		if res != nil {
			newContent := strings.Join(res, "\n")
			go func() {
				ioutil.WriteFile(path, []byte(newContent), 0)
			}()
			resolve["ret"] = "success"
			return resolve
		} else {
			fmt.Print("failed")
			return nil
		}
	}

}
