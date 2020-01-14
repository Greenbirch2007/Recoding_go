package main

import "fmt"


type User struct{
	name string
	age int
}


func main(){
	ma := make(map[int]User)
andes := User{
	name:"andes",
	age:18,

}

ma[1] = andes
// ma[1].age = 19 

andes.age = 19

ma[1] = andes //必须整体替换value

fmt.Printf("%v\n",ma)
}