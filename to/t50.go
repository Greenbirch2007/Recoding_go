package main

import "fmt"

func funcName(a interface{}) string{
	value,ok := a.(string)
	if !ok{
		fmt.Println("It is not ok for type string")
		return ""
	}

	fmt.Println("The value is",value)
	return value
}

func main(){
	// str := "123"
	var a int =10
	funcName(a)
}


