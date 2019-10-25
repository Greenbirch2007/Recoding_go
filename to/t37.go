package main

import "fmt"


//定义结构体类型

type Books struct {
	title string
	author string
	subject string
	book_id int
}


func printBook(book Books){
	fmt.Printf("Book title :%s \n",book.title)
	fmt.Printf("Book author: %s \n",book.author)
	fmt.Printf("Book subject:%s \n",book.subject)
	fmt.Printf("Book book_id: %d\n",book.book_id)
}



func main(){
	var book1 Books //声明book1位Books类型
	var book2 Books


	// book1的描述


	book1.title = "how do go lang"
	book1.author= "sdf"
	book1.subject ="go-incat"
	book1.book_id =100002

	//book2的描述

	book2.title = "ptyho"
	book2.author= "sadf"
	book2.subject = "s"
	book2.book_id = 22
	printBook(book1)
	printBook(book2)
}
