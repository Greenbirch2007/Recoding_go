package main

import (
	"io"
	"os"
)


func main(){
	fi,err := os.Open("path") // 打开输入 *File
	if err != nil {panic(err)}
	defer fi.Close()

	fo,err := os.Create("path") //创建输出*File

	if err != nil { panic(err)}
	defer fo.Close()
	buf := make([]byte,1024)

	for {
		n,err := fi.Read(buf) //从input.txt读取
		if err != nil && err != io.EOF{panic(err)}
		if n== 0{break}
		
		if n2, err := fo.Write(buf[:n];err != nil){
		// 写入output.txt ，直到错误
			panic(err)
		} else if n2 != n{
			panic("error in writing")
}
}
}


用bufio库


