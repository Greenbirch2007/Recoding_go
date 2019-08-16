package main

import "fmt"

type Books struct {
	title stirng
	author string
	subject string
	book_id int
}

func main()  {
	var Book1 Books  //声明Book1 为Books类型

	Book1.title = "go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "go语言"
	Book1.book_id = 999

	printBook(Book1)
}

func printBook(book Books)  {
	fmt.Printf(book.title);
}