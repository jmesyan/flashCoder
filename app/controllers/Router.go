package controllers

import h "flashCoder/app/controllers/home"

var Router = map[string]interface{}{
	"/": map[string]interface{}{
		"/":     new(h.IndexController),
		"index": new(h.IndexController),
		"game":  new(h.GameController),
	},
}
