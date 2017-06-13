package operates

import (
	"fmt"
)

type MsgTip struct {
}

func (op *MsgTip) Execute(params map[string]string, lastRes interface{}) interface{} {
	if msg, ok := params["msg"]; ok {
		fmt.Println(msg)
		WinMsg <- []byte(msg)

		return true
	} else {
		return false
	}

}
