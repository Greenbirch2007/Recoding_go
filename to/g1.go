package main

import "fmt"

func counter() (func()int){
	var i int = 0
	//匿名函数引用了一个外部变量i
	return func() int{
		i++
		return i
	}
}

func main(){
	//创建一个计数器
	count := counter()
	//调用计数器，每次调用都会增加１
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}
