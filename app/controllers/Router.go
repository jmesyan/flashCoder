package controllers

import h "flashCoder/app/controllers/home"
import api "flashCoder/app/controllers/api"

var Router = map[string]interface{}{
	"/": map[string]interface{}{
		"/":        new(h.IndexController),
		"index":    new(h.IndexController),
		"behavior": new(h.BehaviorController),
		"operate":  new(h.OperateController),
		"task":     new(h.TaskController),
		"cron":     new(h.CronController),
	},
	"api": map[string]interface{}{
		"/":        new(api.IndexController),
		"index":    new(api.IndexController),
		"behavior": new(api.BehaviorController),
		"operate":  new(api.OperateController),
		"task":     new(api.TaskController),
		"cron":     new(api.CronController),
	},
}
