package operates

import (
	"fmt"
	// "github.com/andlabs/ui"
	"log"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

type MsgTip struct {
}

func (op *MsgTip) Execute(params map[string]string, lastRes interface{}) interface{} {

	if msg, ok := params["msg"]; ok {
		fmt.Println(msg)
		go func() {
			w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG|sciter.SW_POPUP, nil)
			if err != nil {
				log.Fatal("Create Window Error: ", err)
			}
			w.LoadFile("index.html")
			// setEventHandler(w)
			w.Show()
			w.Run()
		}()

		// js.Global.Call("alert", msg)
		// go func() {
		// 	err := ui.Main(func() {
		// 		w := ui.NewWindow("消息提醒", 400, 300, false)
		// 		w.OnClosing(func(*ui.Window) bool {
		// 			ui.Quit()
		// 			return true
		// 		})
		// 		ui.MsgBox(w, "消息提示", msg)
		// 	})
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }()

		return true
	} else {
		return false
	}

}
