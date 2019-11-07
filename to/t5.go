package main


import (
	"io"
	"log"
	"net"
	"time"
)


func handleConn(c net.Conn){
	defer c.Close()
	for {
	// z直接向客户端套接字，写当前系统时间
	_, er := io.WriteString(c,time.Now().Format("15:04"05\n"))
	if err != nil{
		log.Print(err)
		return //由于defer语法，c.Close()会在
			//return 之前被调动
	}
	time.Sleep(time.Second *1)

}

func main(){
	listener,err := net.Listent("tcp","localhost:8000")
	if err != nil{
		// err != nil 说明err 有值，说明有错误
		log.Fatal(err)
	}

	// for语句如果没有任何条件，表示一直循环，直到遇见break退出
	conn ,err := listener.Accept()
	if err != nil {
		log.Print(err)
		continue
	}

	// 处理客户端链接
	go handleConn(conn)
}
}

现在两个客户端就可以同时连接服务器了


案例:并发的echo服务

p158
