package main


import (
	"io"
	"log"
	"net"
	"os"
)


func dealRecvData(dst io.Writer,src io.Reader){
	//将读缓冲　拷贝到写缓冲
	if _,err := io.Copy(dst,src); err != nil{
		log.Fatal(err)
}
}


func main(){
	conn, err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
}
	defer conn.Close()
	// os.Stdout标准输出，作为写io设备，conn客户端套接字作为读io设备
	dealRecData(os.Stdout,conn)
}


运行这个客户端，依然可以向netcat一样执行


但是如果第一个客户端已经执行的情况下，再打开一个客户端


第二个客户端必须等待第一个客户端完成工作，这样服务端才能继续向后
执行；因为这里的服务器程序同一时间只能处理一个客户端连接。
在这里对服务端程序做一个小改动，使其支持并发：在handleConn
函数调用的地方增加go关键字，让每一次handleConn的调用
都进入一个独立的goroutine



