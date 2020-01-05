package main

import "fmt"

func main()
{
	var a uint = 60 // 60 = 0011 1100
	var b uint = 13 // 13 = 0000 1101
	var c uint = 0
	c = a&b // 12 = 000 1100
	fmt.Printf("第一行　-c 的值为%d \n",c)
	c = a|b // 61 =0011 1101
	fmt.Printf("第二行 -c 的值为%d\n",c)

	c = a^b
	fmt.Printf("第三行　-c 的值为%d\n",c)

	c = a << 2
