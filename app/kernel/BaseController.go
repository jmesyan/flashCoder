package kernel

import (
	"fmt"
	"net/http"
)

type Reponse http.ResponseWriter
type Request *http.Request

type BaseController struct {
	Hello int
}

func (c *BaseController) Welcome(r Request, w Reponse) {
	fmt.Println("this is BaseController welcome function")
}
