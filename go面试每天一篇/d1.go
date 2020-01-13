package main

import "fmt"

func main(){
	str := "hello ,世界" // 通过字符串字面量初始化一个字符串str
	a := []byte(str) // 将字符串转换为[]byte类型切片
	b := []rune(str) // 将字符串转换为[]rune 类型切片
	fmt.Println(a)
	fmt.Println(b)
}