
import (
	ke "flashCoder/app/kernel"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	// "reflect"
)
func (c *IndexController) Test(r Request, w Reponse) {
	var max int
	// c.chances = make([][2]int, 0)
	config := [...][2]int{{800, 1235}, {600, 1235}, {600, 1235}, {590, 1235}, {560, 1235}, {500, 1235}, {450, 1235}, {400, 1235}, {400, 1235}, {300, 1235}, {50, 1235}, {15, 1235}, {5, 1235}, {5, 1235}}

	max = 0
	for i := 0; i < 14; i++ {
		max += config[i][0]
	}
	fmt.Println("max is : ", max)

	begin := time.Now()
	for i := 1; i <= 100; i++ {
		// c.wg.Add(1)
		go func(i int) {
			for j := 1; j <= 1000; j++ {
				chance := rand.Intn(max) + 1
				match := false
				for k := 0; k < 14; k++ {
					if config[k][0] > chance {
						match = true
						// fmt.Println(true, config[k][0], chance)
						// fmt.Fprintln(w, true, config[k][0], chance)
						fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
						fd_content := "1111, " + strconv.Itoa(i) + "-" + strconv.Itoa(j) + "," + strconv.Itoa(config[k][0]) + "," + strconv.Itoa(chance) + "\n"
						buf := []byte(fd_content)
						fd.Write(buf)
						fd.Close()
						// break
					} else {
						chance -= config[k][0]
					}
				}
				if !match {
					// tmp := [2]int{0, chance}
					// fmt.Println(false, 0, chance)
					// fmt.Fprintln(w, false, 0, chance)

					fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
					fd_content := "2222, " + strconv.Itoa(i) + "-" + strconv.Itoa(j) + ", 0," + strconv.Itoa(chance) + "\n"
					buf := []byte(fd_content)
					fd.Write(buf)
					fd.Close()
				}
			}
		}(i)
	}
	// c.wg.Wait()
	// for _, v := range c.chances {
	// 	if v[0] > 0 {
	// 		fmt.Fprintln(w, true, v)
	// 	} else {
	// 		fmt.Fprintln(w, false, v)
	// 	}
	// }

	diff := time.Since(begin)
	fmt.Fprintf(w, "diff is %s", diff)

}

func (c *IndexController) create(config [14][2]int, max int) {
	defer c.wg.Done()
	for j := 0; j < 1000; j++ {
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