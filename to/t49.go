package main

import (
	"fmt"
	
	)

type Phone interface{
	call()
}


type NokiaPhone struct {
}


func (nokiaPhone NokiaPhone) call(){
	fmt.Println("I am Nokia, I can call you!")
}


type IPhone struct {
	}


func (iPhone IPhone) call(){
	fmt.Println("I am iPhone,I can call you!")
}


func main(){
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

上面的例子，定义了一个接口Phone,接口里面有一个方法call()
然后我们在main函数里面定义了一个Phone类型变量，并分贝为之赋值为NokiaPhone和IPhone

然后调用call()方法


类型断言

golang的语言中提供了断言的功能。
golan中的程序都是实现了interface{}的接口，这意味着，所有的类型如string,int,int64甚至是自定义的struct类型都就此拥有
了interface{}的接口，这种方法和java中的Object类型很类似

在一个数据通过func funcName(interface{})的方式传进来的时候，也就意味着这个参数被自动的
转为interface{}的类型


func funcName(a interface{}) string {
	return string(a)
}

编译器会返回


此时，意味着整个转化的过程需要类型断言。类型断言有以下几种形式：

1. 直接断言使用

var a interface{}
fmt.Println("where are you, Jonny?".a.(string))

但是如果断言失败一般会导致panic的发生。所以为了防止panic的发生，
我们需要在断言前进行一定判断



value, ok := a.(string)

如果断言失败，那么ok的值将会是false,但是如果断言成功ok的值将会是true,
同时value将会得到所期待的正确的值


value,ok := a.(string)
if !ok{
	fmt.Println("It's not ok for type string")
	return 
}

fmt.Println("The value is ",value)

