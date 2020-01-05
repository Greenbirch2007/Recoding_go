package main

import "fmt"


func main(){
	var x interface{}
	
	switch i := x.(type){
		case nil:
			fmt.Printf("x的类型: %T",i)
		case int:
			fmt.Printf("x　是 int型")
		case float64:
			fmt.Printf("x　是 float64型")
		case func(int) float64:
			fmt.Print("x 是func(int)型")
		case bool,string:
			fmt.Printf("x是bool或string型")
		default:
			fmt.Printf("未知型")
}
｝


