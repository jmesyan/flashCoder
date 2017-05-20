package main

import (
	ctr "flashCoder/controllers"
	// "flashCoder/utils"
	"fmt"
	// "html"
	"html/template"
	"net/http"
	"reflect"
)

type handler struct {
	Welcome int
	hello   float64
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "golang")
	fmt.Fprintf(w, "jmesyan,%q", "nice to meet you")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("resources/views/test.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		data := make(map[string]string)
		data["helloworld"] = "this is jamesyan"
		t.Execute(w, data)
	} else {
		username := r.Form["username"]
		password := r.Form["password"]
		fmt.Fprintf(w, "username = %s, password = %s", username, password)
	}
}

type refInter interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func main() {
	con := ctr.Router["MainController"]
	res := reflect.ValueOf(con)
	if res.NumMethod() > 0 {
		for i := 0; i < res.NumMethod(); i++ {
			f := res.Method(i)
			if res.Type().Method(i).Name == "Welcome" {
				f.Call(nil)
			}
		}
	}
	// http.Handle("/hello", new(handler))
	// http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/login", login)
	// err := http.ListenAndServe(":8080", nil)
	// utils.CheckError(err)

}
