package home

import (
	ke "flashCoder/app/kernel"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Reponse http.ResponseWriter
type Request *http.Request

type IndexController struct {
	ke.BaseController
}

func (c *IndexController) Index(r Request, w Reponse) {
	data := map[string]interface{}{
		"hello": "world",
		"good":  123,
		"yes":   "yes is good",
		"Title": "这是header测试",
	}
	c.View(w, data)
}

func (c *IndexController) test(r Request, w Reponse) {
	var chance, max, keys int
	config := map[int]int{
		800: 1235,
		600: 1235,
		600: 1235,
		590: 1235,
		560: 1235,
		500: 1235,
		450: 1235,
		400: 1235,
		400: 1235,
		300: 1235,
		50:  1235,
		15:  1235,
		5:   1235,
		5:   1235,
	}
	max = 0
	k, v :=  range config {
		max += k
		keys =append(keys, k)
	}
	fmt.Println("max is : ", max)

	begin := time.Now()
	var match bool
	for i := 0; i < 10000; i++ {
		chance = rand.Intn(max)
		match = false
		for _, key := range keys {
			if key >= chance {
				match = true
				fmt.Println(true, key, config[key])
			} else {
				chance -= key
			}
		}
		if !match {
			fmt.Println(false, chance)
		}

	}
	diff := time.Since(begin)

	fmt.Fprintf(w, "diff is %s", diff)

}
