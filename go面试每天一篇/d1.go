package main

import "fmt"
import "flag"


var ip string
var port int


func init(){
	
}


func main(){
	flag.Parse()
	fmt.Printf("%s:%d",ip,port)
}