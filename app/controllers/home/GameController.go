package home

import (
	ke "flashCoder/app/kernel"
	// "fmt"
)

type GameController struct {
	ke.BaseController
}

func (c *GameController) Index(r Request, w Reponse) {
	data := map[string]interface{}{
		"hello": "this is game index",
	}
	c.View(w, data)
}
