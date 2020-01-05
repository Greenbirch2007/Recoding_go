
package main

// 基于无缓存的channe的main 和goroutine的同步


import (
	"io"
	"log"
	"net"
	"os"
)


func main(){
	conn,err := net.Dial("tcp","127.0.0.1:8001")
	if err != nil{
		log.Fatal(err)
}
	done := make(chan string)
	go func(){
		io.Copy(os.Stdout,conn)
		log.Println("groutine:done!")
		done <- "I am done"
	} ()
		//从客户端输入，将客户端输入的数据发送给客户端套接字
	io.Copy(conn,os.Stdin)
	conn.Close() // 此时main要主动关闭conn,否则goroutine
		///里面的io.Copy()会一直阻塞等待conn
	log.Println("main wait goroutine..")
	<-done
	log.Println("main:done!")

	// 这样我们就保证了"main::done!" 打印之前一定先打印"groutine:done!"
}




