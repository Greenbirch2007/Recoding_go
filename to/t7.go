package main

import (
	"io"
	"log"
	"net"
	"os"
)

func dealRecvData(dst io.Writer,src io.Reader){
	if _,err := io.Copy(dst,src);  err != nil {
		log.Fatal(err)
}
}

func main(){
	conn, err := net.Dial("tcp","127.0.0.1:8001")
	if err != nil{
		log.Fatal(err)
	}

	defer conn.Close()
	go dealRecvData(os.Stdout,conn)
	//　从客户端输入，将客户端标输入的数据发给客户端套接字
	dealRecvData(conn,os.Stdin)
}


在
