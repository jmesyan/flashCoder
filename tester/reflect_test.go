package tester

func TestReflect() {
	type myFloat float64
	var x myFloat = 3.4
	// v := reflect.ValueOf(x)
	v := reflect.ValueOf(&x)
	p := v.Elem()
	fmt.Println(p)
	fmt.Println("value--", v.Type(), "--", v.Kind(), "--", v.Kind() == reflect.Float64, "--", v.Pointer(), p.Type(), p.Kind(), p.Float(), p.CanSet())
	p.SetFloat(7.2)
	fmt.Println(p.Interface().(myFloat) == 7.2)
	hand := new(handler)
	var refhand refInter
	refhand = hand
	v = reflect.ValueOf(refhand)
	fmt.Println("value--", v.Type(), "--", v.Kind(), "--", v.Kind() == reflect.Interface, v.Interface().(*handler))
	for j := 0; j < v.NumMethod(); j++ {
		m := v.Method(j)
		fmt.Println(1212, v.Type().Method(j).Name, m.Type(), m, m.Call(in))
	}
	abc := handler{9, 6.3}
	v = reflect.ValueOf(&abc).Elem()
	fmt.Println("value--", v.Type(), "--", v.Kind(), "--", v.Kind() == reflect.Struct, v.Interface().(handler), v.CanSet())
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Println(f, f.Type(), v.Type().Field(i).Name, f.CanSet())
	}
	v.Field(0).SetInt(88)
	if v.Field(1).CanSet() {
		v.Field(1).SetFloat(63.3)
	} else {
		fmt.Println("can't set field 1")
	}
	fmt.Println(abc, v.FieldByName("hello").CanSet())
}
