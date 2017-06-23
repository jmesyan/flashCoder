package vendor

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

type renderer struct {
	*otto.Otto
}

func NewRenderer(files []string) *renderer {
	r := &renderer{otto.New()}
	r.runFile("assets/global.js")
	r.runFile("assets/react.js")
	r.runFiles(files)
	return r
}

func (r *renderer) RunCmd(cmd string) otto.Value {
	v, err := r.Run(cmd)
	if err != nil {
		panic(err)
	}
	return v
}

func (r *renderer) runFiles(files []string) {
	for _, file := range files {
		r.runFile(file)
	}
}

func (r *renderer) runFile(path string) otto.Value {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result, err := r.Run(data)
	if err != nil {
		panic(err)
	}

	return result
}
