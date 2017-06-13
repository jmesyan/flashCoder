package operates

import (
	"fmt"
	"github.com/andlabs/ui"
)

var WinMsg chan []byte
var Operates = map[string]interface{}{
	"OpenFile":  new(OpenFile),
	"WriteFile": new(WriteFile),
	"CloseFile": new(CloseFile),
	"MsgTip":    new(MsgTip),
}

func init() {
	WinMsg = make(chan []byte, 255)
	go func() {
		err := ui.Main(func() {
			w := ui.NewWindow("消息提醒", 400, 300, false)
			defer ui.Quit()
			for {
				select {
				case msg := <-WinMsg:
					fmt.Println(msg)
					ui.MsgBox(w, "消息提示", string(msg))
				}
			}

		})
		if err != nil {
			panic(err)
		}
	}()
}
