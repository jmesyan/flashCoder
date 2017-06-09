package main

import (
	ctr "flashCoder/app/controllers"
	"flashCoder/app/kernel/cron"
	"flashCoder/utils"
	// "log"
	// "fmt"
	"net/http"
	"time"
)

func main() {
	//静态资源
	ctr.StaticMap = make(map[string]string)
	ctr.AddstaticMap("/assets", "./resources/assets")
	ctr.AddstaticMap("/components", "./resources/assets/js/components")
	ctr.AddstaticMap("/favicon.ico", "./resources/assets/images/flash.ico")
	//启动事件监听
	crons := new(cron.CronWatcher)
	crons.Watching()
	//启动服务器
	var Handler http.Handler
	Handler = new(ctr.Controller)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        Handler,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	utils.CheckError(err)
}
