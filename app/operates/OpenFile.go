package operates

import (
	"flashCoder/utils"
	// "fmt"
	"os"
	// "reflect"
)

type OpenFile struct {
}

func (b *OpenFile) Execute(params map[string]string, lastRes interface{}) interface{} {

	if path, ok := params["path"]; ok {
		fd, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		utils.LogError("error", err)
		return fd
	} else {
		return false
	}

}
