package main


import (
	"io/ioutil"
)


func main(){
	b,err := ioutil.ReadFile("input.txt") //读文件
	if err != nil { panic(err)}
	
	err  = ioutil.WriteFile("output.txt",b,0644) //写文件
	if err != nil {panic(err)}
}



