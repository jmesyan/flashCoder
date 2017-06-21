package main

import (
	"flashCoder/app/controllers"
	"flashCoder/app/kernel/crontab"
	"flashCoder/utils"
	// "log"
	// "fmt"
	"net/http"
	"time"
)

func main() {
	//启动事件监听
	crontab.Watching()
	//启动服务器
	var Handler http.Handler
	Handler = new(controllers.Controller)
	s := &http.Server{
		Addr:           ":6339",
		Handler:        Handler,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	utils.CheckError(err)
}
