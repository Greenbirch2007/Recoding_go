package main

import "fmt"


//定义一个结构体

type T  struct {
	name string

}

func (t T) method1() {
	t.name = "new name1"
}


func (t *T) method2(){
	t.name = "new name2"
}

func main(){
	t := T{"old name"}
	fmt.Println("method1 调用前",t.name)
	t.method1()
	fmt.Println("method2 调用后",t.name)
	t.method2()
	fmt.Println("method2 调用后",t.name)
}


当调用t.method1()时相当于method1(t),实参和行参都是类型T,可以接受。
此时method1()中的t只是参数t的值拷贝，所以method1()的修改影响不到main中的t变量


当调用 t.method2() => method2(t),这是将T类型传给了*T类型，go可能会取t的地址传进去:
method2(&t) 所以 method()的修改可以影响t


T类型的变这两个方法都是拥有的

方法值和方法表达式


方法值


我们经常选择一个方法，并且在同一个表达式里执行，比如p.Distance()形式，实际上将其分成两步来执行也是可能的。
p.Distance叫做"选择器",选择器会返回一个方法"值" 一个将方法(Point.Distance)绑定到特定接收器变量的函数。

这个函数可以不通过指定其接收器即可被调用；即调用时不需要指定接收器，只要传入函数的参数即可

p110
	
