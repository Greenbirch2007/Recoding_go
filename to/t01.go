package main

import (
	"fmt"
	"reflect"
)

func main()  {
	//声明整型变量a并赋初值

	var a int = 1024

	ValueOfA := reflect.ValueOf(a)

	var getA int = ValueOfA.Interface().(int)
	var getA2 int = int(ValueOfA.Int())
	fmt.Println(getA,getA2)
}