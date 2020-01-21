package main

type Printer interface{
	Print()
}

type S struct{}


func (s S) Print(){
	println("print")
}


func main(){
	var i Printer

	//没有初始化的接口调用其方法会产生panic
	//必须初始化
	i = S{}
	i.Print()
}