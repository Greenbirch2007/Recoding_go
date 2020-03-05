package main



import "fmt"


type Person struct{
    age int
}


func (p Person) Elegance() int{
    return p.age
}


func (p *Person) GetAge(){
    p.age += 1

}


func main(){
    //p1 是值类型
    
    p := Person{age:18}

    //值类型　调用接收者也是值类型的方法


    fmt.Println(p.howOld())

    //值类型　　调用接收者是指针类型的方法

    p.GetAge()

    fmt.Println(p.GetAge())

    //p2 是指针类型

    //指针类型　调用接收者是值类型的方法

    fmt.Println(p2.GetAge())

    //指针类型　调用接收者也是指针类型的方法


    p2.GetAge()
    fmt.Println(p2.GetAge())
}