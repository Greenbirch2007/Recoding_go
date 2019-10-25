package main



import "fmt"


//定义结构体类型


type Books struct{
	title string
	author string
	subject string
	book_id int
}



func printBook(book,*Books){
	fmt.Printf("Book title:%s\n",book.title)
	fmt.Printf("Book author:%s\n",book.author)
	fmt.Printf("Book subject:%s\n",book.subject)
	fmt.Printf("Book book_id:%d\n",book.book_id)

}



func main(){
	var book1 Booka //声明book1为Book类型
	var book2 Books


	// book1的描述

	book1.title = "how~"
	book1.author = "how~"
	book1.subject = "how~"
	book1.book_id = 234


	printBook(&book1)

}

