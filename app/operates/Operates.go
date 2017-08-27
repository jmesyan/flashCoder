package operates

import (
	"context"
	"flashCoder/utils"
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
	"MsgTip":    new(MsgTip),
	"ParseTmpl": new(ParseTmpl),
	"FileEdit":  new(FileEdit),
	"Spider":new(Spider),
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
					utils.CheckError("info", "WinMsg:"+string(msg))
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
	Execute(ctx context.Context) map[string]interface{}
	setParams(t ParamsType, val interface{})
}

func parseParams(op OperatesIn, ctx context.Context) {
	if global := ctx.Value(ParamsGlobal); global != nil {
		op.setParams(ParamsGlobal, global)
	}

	if curres := ctx.Value(ParamsCurRes).(map[string]interface{}); curres != nil {
		op.setParams(ParamsCurrent, curres["current"])
		op.setParams(ParamsResolve, curres["resolve"])
	}

}

type OperatesBase struct {
	globalParams  map[string]string
	currentParams map[string]string
	resolveParams map[string]interface{}
}

func (op *OperatesBase) setParams(t ParamsType, val interface{}) {
	switch t {
	case ParamsGlobal:
		op.globalParams = val.(map[string]string)
	case ParamsCurrent:
		op.currentParams = val.(map[string]string)
	case ParamsResolve:
		op.resolveParams = val.(map[string]interface{})
	}
}
