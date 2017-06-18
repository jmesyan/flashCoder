package tester

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	// "time"
)

// type Operates interface {
// 	Execute(params )
// }

type OperateParams struct {
	global  context.Context
	current context.Context
	resolve context.Context
}

func TestOperate(t *testing.T) {
	type favContextKey string
	parent, cancel := context.WithCancel(context.Background())
	defer cancel()
	paramsKey := favContextKey("params")
	global := context.WithValue(context.Background(), paramsKey, map[string]string{"global1": "key1"})
	current := context.WithValue(context.Background(), paramsKey, map[string]string{"current1": "current1"})
	resolve := context.WithValue(context.Background(), paramsKey, map[string]string{"resolve1": "resolve1"})
	val := OperateParams{global, current, resolve}
	c1 := context.WithValue(parent, paramsKey, val)
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	f(c1, paramsKey)
	if v := c1.Value(paramsKey); v != nil {
		v1 := v.(OperateParams)
		f(v1.global, paramsKey)
		f(v1.current, paramsKey)
		f(v1.resolve, paramsKey)
	}

	global = context.WithValue(context.Background(), paramsKey, map[string]string{"global2": "key2"})
	current = context.WithValue(context.Background(), paramsKey, map[string]string{"current2": "current2"})
	resolve = context.WithValue(context.Background(), paramsKey, map[string]string{"resolve2": "resolve2"})
	val = OperateParams{global, current, resolve}
	c2 := context.WithValue(parent, paramsKey, val)
	f(c2, paramsKey)
	if v := c2.Value(paramsKey); v != nil {
		v1 := v.(OperateParams)
		f(v1.global, paramsKey)
		f(v1.current, paramsKey)
		f(v1.resolve, paramsKey)
	}

	f(c1, paramsKey)
	if v := c1.Value(paramsKey); v != nil {
		v1 := v.(OperateParams)
		f(v1.global, paramsKey)
		f(v1.current, paramsKey)
		f(v1.resolve, paramsKey)
	}
}

type test struct {
}

func (t *test) Execute(params context.Context) (resolve map[string]string) {
	// fmt.Println(params)
	resolve = make(map[string]string)
	if pa := params.Value(ParamsGlobal).(map[string]string); pa != nil {
		num := pa["num"]
		resolve["resolve"+num] = "hello" + num
	}

	return
}

type ParamsType int

const (
	_ ParamsType = iota
	ParamsGlobal
	ParamsCurRes
)

func Test123(t *testing.T) {
	task, cancel := context.WithCancel(context.Background())
	// globalVal := make(map[string]string)
	global := context.WithValue(task, ParamsGlobal, map[string]string{})
	current := map[string]string{"current1": "1"}
	resolve := map[string]string{}
	for i := 1; i < 10; i++ {
		opType := rand.Intn(2)
		num := strconv.Itoa(i)
		if opType == 1 {
			//传递全局变量
			if pa := global.Value(ParamsGlobal).(map[string]string); pa != nil {
				pa["num"] = num
				pa["global"+num] = num
				global = context.WithValue(task, ParamsGlobal, pa)

			}
		} else {

			//将当前和resolve变量结合传参给content
			current = map[string]string{"current" + num: num}
			val := map[string]map[string]string{
				"current" + num: current,
				"resolve" + num: resolve,
			}
			curres := context.WithValue(global, ParamsCurRes, val)
			go func() {

				select {
				case <-curres.Done():
					fmt.Println(curres.Value(ParamsGlobal).(map[string]string), curres.Value(ParamsCurRes).(map[string]map[string]string))
				default:
					abc := new(test)
					resolve = abc.Execute(curres)
					fmt.Println(global, val)
				}
			}()

		}

	}

	cancel()
	// time.Sleep(100 * time.Millisecond)

}
