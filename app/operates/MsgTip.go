package operates

import (
	"fmt"
	"github.com/andlabs/ui"
)

type MsgTip struct {
}

func (op *MsgTip) Execute(params map[string]string, lastRes interface{}) interface{} {

	if msg, ok := params["msg"]; ok {
		fmt.Println(msg)
		go func() {
			err := ui.Main(func() {
				w := ui.NewWindow("消息提醒", 400, 300, false)
				w.OnClosing(func(*ui.Window) bool {
					ui.Quit()
					return true
				})
				ui.MsgBox(w, "消息提示", msg)
			})
			if err != nil {
				panic(err)
			}
		}()

		return true
	} else {
		return false
	}

}
