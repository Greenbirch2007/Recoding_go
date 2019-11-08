package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)


func main(){
	listener,err := net.Listen("tcp","127.0.0.1:8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		//处理客户端业务
		handleConn(conn)
		}
}


func echo(c net.Conn,outstr string ,delay time.Duration){
		fmt.Fprintln(c,strings.ToUpper(outstr))
		time.Sleep(delay)
		fmt.Fprintln(c,outstr)
		time.Sleep(delay)
		fmt.Fprintln(c, strings.ToLower(outstr))
}



func handleConn(c net.Conn){
		//input 是一个Scanner类型，该类型可以通过scan方法
	// 依次迭代从io设备中读取数据，直到遇到eof为止
	input := bufio.NewScanner(c)
	//　scan方法　如果缓冲有数据会返回true,否则返回false
	for input.Scan(){
	 	// 如果有数据　　input.Text()可以取出
			echo(c,input.Text(),1*time.Second)
}

	c.Close()

}



