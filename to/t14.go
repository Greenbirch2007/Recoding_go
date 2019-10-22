package main

import "fmt"


func main(){
	var a int =21
	var c int
	c = a
	fmt.Printf("第１行　-= 运算符实例，c值为= %d\n",c)

	c += a
	fmt.Printf("第2行　+= 运算符实例,c值为=%d\n",c)
	c -= a
	fmt.Printf("第3行　-= 运算符实例,c值为=%d\n",c)
