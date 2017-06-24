package operates

import (
	"context"
)

type MsgTip struct {
	OperatesBase
}

func (op *MsgTip) Execute(ctx context.Context) map[string]interface{} {
	select {
	case <-ctx.Done():
		return nil
	default:
		parseParams(op, ctx)
		resolve := make(map[string]interface{})
		if msg, ok := op.currentParams["msg"]; ok {
			WinMsg <- []byte(msg)
			resolve["ret"] = "success"
			return resolve
		} else {
			return nil
		}
	}

}
