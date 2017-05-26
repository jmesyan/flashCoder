package home

import (
	ke "flashCoder/app/kernel"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
	// "reflect"
)

type Reponse http.ResponseWriter
type Request *http.Request

type IndexController struct {
	ke.BaseController
	chances [][2]int
	wg      sync.WaitGroup
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
	var max int
	c.chances = make([][2]int, 0)
	config := [...][2]int{{800, 1235}, {600, 1235}, {600, 1235}, {590, 1235}, {560, 1235}, {500, 1235}, {450, 1235}, {400, 1235}, {400, 1235}, {300, 1235}, {50, 1235}, {15, 1235}, {5, 1235}, {5, 1235}}

	max = 0
	for i := 0; i < 14; i++ {
		max += config[i][0]
	}
	fmt.Println("max is : ", max)

	begin := time.Now()
	for j := 0; j < 10; j++ {
		c.wg.Add(1)
		go c.create(config, max)
	}
	c.wg.Wait()
	for _, v := range c.chances {
		if v[0] > 0 {
			fmt.Fprintln(w, true, v)
		} else {
			fmt.Fprintln(w, false, v)
		}
	}

	diff := time.Since(begin)
	fmt.Fprintf(w, "diff is %s", diff)

}

func (c *IndexController) create(config [14][2]int, max int) {
	defer c.wg.Done()
	for j := 0; j < 10000; j++ {
		chance := rand.Intn(max-1) + 1
		match := false
		for k := 0; k < 14; k++ {
			if config[k][0] >= chance {
				match = true
				c.chances = append(c.chances, config[k])
				break
			} else {
				chance -= config[k][0]
			}
		}
		if !match {
			tmp := [2]int{0, chance}
			c.chances = append(c.chances, tmp)
		}
	}
}
