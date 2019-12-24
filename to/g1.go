package main

import (
	"fmt"
	"net/http"
	"time"

func main()  {
	fmt.Println("*******服务器启动，监听端口 8080*")
	mux := http.NewServeMux()

	//处理静态资源文件
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/",http.StripPrefix("/static",files)
	mux.HandleFunc("/",index)

	//配置服务器
	server := &http.Server{
		Addr:"0.0.0.0:8080",
		Handler: mux, //设置多路复用器
		ReadTimeout:time.Duration(10*int64(time.Second)),
		WriteTimeout: time.Duration(200*int64(time.Second)),
		MaxHeaderBytes:1 <<20, //左移运算，等同于1*2^20,高性能乘法运算
		}
		server.ListenAndServe()
}
func index(writer http.ResponseWriter,request *http.Request)  {
	fmt.Fprintf(writer,"hello world!")
})