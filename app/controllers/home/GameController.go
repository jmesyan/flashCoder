package home

import (
	ke "flashCoder/app/kernel"
	"fmt"
)

type GameController struct {
	ke.BaseController
}

func (c *GameController) Index(r Request, w Reponse) {
	fmt.Println(r.Method)
	fmt.Println("this is home GameController index")
	w.Header().Set("Server", "golang")
	fmt.Fprintf(w, "jmesyan,%q", "location in GameController index")
}

func (c *GameController) Welcome(r Request, w Reponse) {
	fmt.Fprintf(w, "jmesyan,%q", "location in GameController welcome")
}
