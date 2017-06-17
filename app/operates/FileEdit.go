package operates

import (
	"flashCoder/app/kernel/file"
	"fmt"
	"io/ioutil"
	// "reflect"
	"strconv"
	"strings"
)

type FileEdit struct {
}

func (op FileEdit) Execute(params map[string]string, lastRes interface{}) interface{} {
	orgin, ok := params["params"]
	if !ok {
		return false
	}
	data := make(map[string]string)
	tmp := strings.Split(orgin, ";")
	for _, v := range tmp {
		tmp2 := strings.Split(v, "=")
		data[tmp2[0]] = tmp2[1]
	}

	path := data["path"]
	funcName := data["funcName"]
	offset, _ := strconv.Atoi(data["offset"])
	isBegin := false
	if strings.Trim(data["isBegin"], " ") == "1" {
		isBegin = true
	}

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return false
	}
	tmpFile := strings.Split(string(dat), "\n")
	content := lastRes.(string)
	fh := new(file.FlashFile)
	res := fh.AddFuncContent(tmpFile, funcName, content, isBegin, offset)
	if res != nil {
		newContent := strings.Join(res, "\n")
		go func() {
			ioutil.WriteFile(path, []byte(newContent), 0)
		}()
		return true
	} else {
		fmt.Print("failed")
		return false
	}

}
