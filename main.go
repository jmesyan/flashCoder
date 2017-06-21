package main

import (
	"flashCoder/app/controllers"
	"flashCoder/app/kernel/crontab"
	"flashCoder/utils"
	"github.com/go-ini/ini"
	"net/http"
	"time"
)

func main() {
	//启动事件监听
	crontab.Watching()
	//获取server配置
	config, err := ini.Load(utils.GetRootDirectory() + "/.env")
	utils.CheckError(err)
	servercf, err := config.GetSection("server")
	utils.CheckError(err)
	serverPort := "6339"
	if servercf.HasKey("port") {
		serverPort = servercf.Key("port").String()
	}
	//启动服务器
	var Handler http.Handler
	Handler = new(controllers.Controller)
	s := &http.Server{
		Addr:           ":" + serverPort,
		Handler:        Handler,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = s.ListenAndServe()
	utils.CheckError(err)
}
