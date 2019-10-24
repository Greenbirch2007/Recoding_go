package main

import "fmt"


func swap(x,y string) (string , string){
	return y,x
}

func main(){
	a,b := swap("Mathe","lsfa")
	fmt.Println(a,b)
}



