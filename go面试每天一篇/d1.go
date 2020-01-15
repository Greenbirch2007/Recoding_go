package main

import  "fmt"


type T struct{
	n int
}


func main(){
	m := make(map[int]T)
	m[0].n = 1
	fmt.Println(m[0].n)
}