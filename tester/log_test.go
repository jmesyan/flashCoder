package tester

import (
	// "fmt"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	fileName := "Info_First.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	debugLog := log.New(logFile, "[Info]", log.LstdFlags|log.Llongfile)
	debugLog.Println("A Info message here")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("A Debug Message here ")
}
