package main


import (
	"fmt"
	"reflect"

)


func main(){
	var x float64 = 3.4
	var y string = "abc"
	var a uint8 =10
	var b uint16= 10
	var c int =100

	fmt.Println("x's type:",reflect.TypeOf(x))
	fmt.Println("y's type:",reflect.TypeOf(y))
	fmt.Println("a's type:",reflect.TypeOf(a))
	fmt.Println("b's type:",reflect.TypeOf(b))
	fmt.Println("c's type:",reflect.TypeOf(c))
}


