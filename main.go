package main

import (
	ctr "flashCoder/app/controllers"
	"log"
	"net/http"
	"time"
)

func main() {
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
