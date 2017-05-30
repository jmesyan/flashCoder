package behaviors

import (
	// "fmt"
	// "os"
	"reflect"
)

type WriteFile struct {
}

func (b *WriteFile) Execute(params map[string]string, lastRes interface{}) interface{} {

	if content, ok := params["content"]; ok {
		last := reflect.ValueOf(lastRes)
		if last.Kind() == reflect.Ptr {
			buf := []byte(content)
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(buf)
			last.MethodByName("Write").Call(in)
			return lastRes
		}
		return nil
	} else {
		return nil
	}
}
