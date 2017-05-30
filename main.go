package main

import (
	ctr "flashCoder/app/controllers"
	"log"
	"net/http"
	"time"
)

func main() {
	//静态资源
	ctr.StaticMap = make(map[string]string)
	ctr.AddstaticMap("/assets", "./resources/assets")
	ctr.AddstaticMap("/components", "./resources/assets/js/components")
	ctr.AddstaticMap("/favicon.ico", "./resources/assets/images/flash.ico")
	//启动服务器
	var Handler http.Handler
	Handler = new(ctr.Controller)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        Handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
