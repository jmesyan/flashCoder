package tester

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for i := 0; i < 10; i++ {
		time := <-ticker.C
		fmt.Println(time.String())
	}
	fmt.Println("it is end")
}
