package operates

import (
	"context"
	"fmt"
	"github.com/andlabs/ui"
)

type ParamsType int

const (
	_ ParamsType = iota
	ParamsGlobal
	ParamsCurrent
	ParamsResolve
	ParamsCurRes
)

var WinMsg chan []byte
var Operates = map[string]OperatesIn{
	// "OpenFile":  new(OpenFile),
	// "WriteFile": new(WriteFile),
	// "CloseFile": new(CloseFile),
	"MsgTip":    new(MsgTip),
	"ParseTmpl": new(ParseTmpl),
	"FileEdit":  new(FileEdit),
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
					fmt.Println("WinMsg")
					ui.MsgBox(w, "消息提示", string(msg))
				}
			}

		})
		if err != nil {
			panic(err)
		}
	}()
}

type OperatesIn interface {
	Execute(ctx context.Context) map[string]string
	setParams(t ParamsType, val map[string]string)
}

func parseParams(op OperatesIn, ctx context.Context) {
	if global := ctx.Value(ParamsGlobal).(map[string]string); global != nil {
		op.setParams(ParamsGlobal, global)
	}

	if curres := ctx.Value(ParamsCurRes).(map[string]map[string]string); curres != nil {
		op.setParams(ParamsCurrent, curres["current"])
		op.setParams(ParamsResolve, curres["resolve"])
	}

}

type OperatesBase struct {
	globalParams  map[string]string
	currentParams map[string]string
	resolveParams map[string]string
}

func (op *OperatesBase) setParams(t ParamsType, val map[string]string) {
	switch t {
	case ParamsGlobal:
		op.globalParams = val
	case ParamsCurrent:
		op.currentParams = val
	case ParamsResolve:
		op.resolveParams = val
	}
}
