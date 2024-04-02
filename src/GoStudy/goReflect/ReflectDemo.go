package goReflect

import "reflect"

func reflectArg(arg interface{}) {

	println(reflect.TypeOf(arg))
	println(reflect.ValueOf(arg))
}

func main() {
	var num float32 = 1.23223423
	reflectArg(num)

	//
	//1、通过varType := reflect.TypeOf(var) 获取var 的type，type有 varType.name()，varType.NumField() // type有多少个属性
	//2、得到每一个field 数据类型
	//3、通过field有一个Interface() 方法 得到对应value
	var reflectDemo = DemoType{1, "reflectDemo test"}
	demoType := reflect.TypeOf(reflectDemo)
	demoValue := reflect.ValueOf(reflectDemo)

	for i := 0; i < demoType.NumField(); i++ {
		field := demoType.Field(i)
		println(field.Name)
		println(field.Type)
		value := demoValue.Field(i).Interface()
		println("field [", i, "] = ", field, ",value = ", value)
	}

	// 通过type 获取方法
	for i := 0; i < demoType.NumMethod(); i++ {
		method := demoType.Method(i)

		println(method.Name)
		println(method.Func)
	}
}

type DemoType struct {
	DemoId   int
	DemoName string
}

func (this *DemoType) ReflectFuncDemo() {
	println(this.DemoId)
}
