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
		// 直接向客户端套接字，写当前系统时间
		_,err := io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil{
			log.Print(err)
			return //由于defer 语法，c.Close()会在return 之前被调用
			
}
	time.Sleep(time.Second *1)
}


}

func main(){
	listener,err := net.Listen("tcp","localhost:8000")
	if err != nil {
		//err != nil 说明err有值，说明有错
			log.Fatal(err)
}
	// for 语句如果没有任何条件，表示一直循环，直到遇见break退出
	for {
		// 得到新的客户端链接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			}
		// 处理客户端链接
		handleConn(conn)
}
}



