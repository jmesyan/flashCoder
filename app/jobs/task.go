package jobs

import (
	"context"
	"encoding/json"
	"flashCoder/app/models"
	"flashCoder/app/operates"
)

func TaskExecute(taskDetail models.FlashTask) {
	taskBehavior := models.Task.GetTaskBehavior(taskDetail.Tid, taskDetail.Tcate)
	task, _ := context.WithCancel(context.Background())
	global := context.WithValue(task, operates.ParamsGlobal, map[string]string{})
	resolve := make(map[string]interface{})
	for _, v := range taskBehavior {
		bv := models.Behavior.GetBehavior(v.Bid)
		optag := models.Operate.GetOperateTagById(bv.Opid)
		var params []models.OperateParams
		json.Unmarshal([]byte(v.Paramsin), &params)
		if optag == "ParamsGlobal" {
			if pa := global.Value(operates.ParamsGlobal).(map[string]string); pa != nil {
				for _, param := range params {
					pa[param.Name] = param.Value
				}
				global = context.WithValue(task, operates.ParamsGlobal, pa)

			}
		} else {
			current := make(map[string]string)
			for _, param := range params {
				current[param.Name] = param.Value
			}
			val := map[string]interface{}{
				"current": current,
				"resolve": resolve,
			}
			curres := context.WithValue(global, operates.ParamsCurRes, val)
			if operate, ok := operates.Operates[optag]; ok {
				resolve = operate.Execute(curres)
			}
		}

	}
}
