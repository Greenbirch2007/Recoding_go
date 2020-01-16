package main

import "os"


func main(){
	defer func(){
		println("defer")
	}()
	println("func body")

	os.Exit(1)
}