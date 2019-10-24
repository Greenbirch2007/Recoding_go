package main

import "fmt"

func main(){
	//局部变量定义
	var a int =100;
	// 判断布尔表达式

	if a < 20 {
		// 如果条件为true 则执行以下语句
		fmt.Printf("a　小于20 \n");
	} else{
		// 如果条件为false　则执行以下语句
		fmt.Printf("a　不小于　20\n");
	}
	fmt.Print("a 的值为: %d\n",a);
}
