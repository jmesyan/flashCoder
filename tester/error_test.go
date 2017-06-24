package tester

import (
	"errors"
	// "fmt"
	// "reflect"
	"flashCoder/utils"
	"testing"
)

func TestErr(t *testing.T) {
	utils.CheckError("debug", errors.New("what is your name"))
}
