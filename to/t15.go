package main

import "fmt"


func main(){
	var a int =4
	var b int32
	var c float32
	var ptr *int
	
	// 运算符实例
	fmt.Printf("第1行　a变量类型为 =%T\n",a);

	//　＆　和　* 运算符实例
	ptr = &a // ptr包含了ａ变量的地址
	fmt.Printf("a的值为 %d\n",a)
	fmt.Printf("ptr 为%d\n",*ptr);
}
