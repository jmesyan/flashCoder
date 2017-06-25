package operates

import (
	"context"
	"flashCoder/supplier/file"
	"flashCoder/utils"
	"io/ioutil"
	"path/filepath"
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

		if _, ok := op.currentParams["type"]; !ok {
			utils.CheckError("error", "没有指定文件编辑类型")
			return nil
		}

		editType := op.currentParams["type"]

		switch editType {
		case "AddFuncContent":
			ret := op.addFuncContent()
			resolve["ret"] = ret
			return resolve
		}
		resolve["ret"] = false
		return resolve
	}

}

//在函数中添加代码
func (op *FileEdit) addFuncContent() bool {
	path := op.globalParams["path"]
	funcName := op.currentParams["funcName"]
	offset, err := strconv.Atoi(op.currentParams["offset"])
	utils.CheckError("error", err)
	isBegin := false
	if strings.TrimSpace(op.currentParams["isBegin"]) == "1" {
		isBegin = true
	}

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		utils.CheckError("error", err)
		return false
	}
	tmpFile := strings.Split(string(dat), "\n")
	var content string
	if _, ok := op.resolveParams["content"]; ok {
		content = op.resolveParams["content"].(string)
	} else {
		content = op.globalParams["content"]
	}
	fh := &file.FlashFile{filepath.Ext(path)}
	res := fh.AddFuncContent(tmpFile, funcName, content, isBegin, offset)
	if res != nil {
		newContent := strings.Join(res, "\n")
		go func() {
			ioutil.WriteFile(path, []byte(newContent), 0)
		}()
		utils.CheckError("info", "添加内容到文件成功")
		return true
	} else {
		utils.CheckError("error", "添加内容到文件失败")
		return false
	}
}
