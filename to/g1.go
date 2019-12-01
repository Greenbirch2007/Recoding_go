package main

import "fmt"

//声明全局变量


//函数定义－两数字相加

func sum(a,b int) int  {
	fmt.Printf("sum()函数中a = %d\n",a)
	fmt.Printf("sum()函数中b = %d\n",b)
	return a +b
}

func main()  {
	//main函数中声明巨变
	var a int =10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中a=%d\n",a)
	c = sum(a,b)
	fmt.Printf("main()函数中c=%d\n",c)

}