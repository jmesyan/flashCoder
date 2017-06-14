package operates

import (
	"fmt"
	"reflect"
)

type CloseFile struct {
}

func (b *CloseFile) Execute(params map[string]string, lastRes interface{}) interface{} {
	last := reflect.ValueOf(lastRes)
	if last.Kind() == reflect.Ptr {
		fmt.Println("good")
		last.MethodByName("Close").Call(nil)
	}
	return true
}
