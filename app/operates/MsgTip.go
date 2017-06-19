package operates

import (
	"context"
	"fmt"
)

type MsgTip struct {
	OperatesBase
}

func (op *MsgTip) Execute(ctx context.Context) map[string]string {
	select {
	case <-ctx.Done():
		return nil
	default:
		parseParams(op, ctx)
		resolve := make(map[string]string)
		if msg, ok := op.currentParams["msg"]; ok {
			fmt.Println(msg)
			WinMsg <- []byte(msg)
			resolve["ret"] = "success"
			return resolve
		} else {
			return nil
		}
	}

}
