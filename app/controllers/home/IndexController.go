package home

import (
	ke "flashCoder/app/kernel"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	// "reflect"
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

func (c *IndexController) Test(r Request, w Reponse) {
	var chance, max int
	config := [...][2]int{{800, 1235}, {600, 1235}, {600, 1235}, {590, 1235}, {560, 1235}, {500, 1235}, {450, 1235}, {400, 1235}, {400, 1235}, {300, 1235}, {50, 1235}, {15, 1235}, {5, 1235}, {5, 1235}}

	max = 0
	for i := 0; i < 14; i++ {
		max += config[i][0]
	}
	fmt.Println("max is : ", max)

	begin := time.Now()
	var match bool
	for j := 0; j < 100000; j++ {
		chance = rand.Intn(max)
		match = false
		for k := 0; k < 14; k++ {
			if config[k][0] > chance {
				match = true
				fmt.Fprintln(w, true, config[k][0], chance)
				break
			} else {
				chance -= config[k][0]
			}
		}
		if !match {
			fmt.Fprintln(w, false, chance)
		}

	}
	diff := time.Since(begin)

	fmt.Fprintf(w, "diff is %s", diff)

}
