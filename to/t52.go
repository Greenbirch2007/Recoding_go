package main

import (
	"fmt"
	"os"
)


func test_read_file(filename string){
		fin ,err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		}
		defer fin.Close()
		buf := make([]byte,1024) // 开辟1024个字节的slice 作为缓冲
		for {	
			n, _ := fin.Read(buf)
			if n== 0{
			// o 表示到达EOF
				break
			}

			os.Stdout.Write(buf)
}
}

func main(){
	test_read_file("path")


同时使用os.Open和os.Create操作文件




