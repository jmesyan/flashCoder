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
			ret := op.addContent("funcName", editType)
			resolve["ret"] = ret
			return resolve
		case "AddClassContent":
			ret := op.addContent("className", editType)
			resolve["ret"] = ret
			return resolve
		case "AddRouteGroupContent":
			ret := op.addContent("routeName", editType)
			resolve["ret"] = ret
			return resolve
		}
		resolve["ret"] = false
		return resolve
	}

}

func (op *FileEdit) getContent() string {
	content := ""
	if _, ok := op.resolveParams["content"]; ok {
		content = op.resolveParams["content"].(string)
	} else {
		content = op.globalParams["content"]
	}
	return content
}

func (op *FileEdit) addContent(ckey, editType string) bool {
	path := op.globalParams["path"]
	contains := op.currentParams[ckey]
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
	content := op.getContent()
	fh := &file.FlashFile{filepath.Ext(path)}
	var res []string
	switch editType {
	case "AddFuncContent":
		res = fh.AddFuncContent(tmpFile, contains, content, isBegin, offset)
	case "AddClassContent":
		res = fh.AddClassContent(tmpFile, contains, content, isBegin, offset)
	case "AddRouteGroupContent":
		res = fh.AddRouteGroupContent(tmpFile, contains, content, isBegin, offset)
	}

	if res != nil {
		newContent := strings.Join(res, "\n")
		go func() {
			ioutil.WriteFile(path, []byte(newContent), 0)
		}()
		utils.CheckError("info", editType+"添加内容成功")
		return true
	} else {
		utils.CheckError("error", editType+"添加内容失败")
		return false
	}
}
