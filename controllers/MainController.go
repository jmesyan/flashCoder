package controllers

import (
	"fmt"
)

var Router = make(map[string]interface{})

type MainController struct {
	Hello int
}

func (m *MainController) Welcome() {
	fmt.Println("this is maincontroller welcome function")
}

func init() {
	Router["MainController"] = new(MainController)
}
