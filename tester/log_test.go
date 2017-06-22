package tester

import (
	// "fmt"
	"flashCoder/utils"
	// "log"
	// "os"
	"testing"
)

func TestLog(t *testing.T) {
	utils.Loger.Debug("hello world")
	utils.Loger.Fatal("dfdfdvery good", new(utils.LogHandler))
	utils.Loger.Info("very good")
}
