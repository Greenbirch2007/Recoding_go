package main

import (
	"fmt"
)

func main()  {
	//定义一个数组
	var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	//youArray := [6]int{6,6,6,6,6,6}
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for k, v := range mySlice {
		fmt.Print(k,v)
	}
	fmt.Println()
}


