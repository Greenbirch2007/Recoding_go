package main

import "fmt"
import "runtime"

func main(){
	runtime.GOMAXPROCS(1)
	go func(){
		for i :=0;i <20;i++{
			fmt.Println(i)
		}
	}()

	for {}
}