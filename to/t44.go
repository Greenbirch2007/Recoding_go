package main

import "fmt"
import "math"


type Point struct{X,Y float64 }

//传统的方法

func Distance(p,q Point) float64 {
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}

//这是给struct Point类型定义一个方法


func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}


type Path []Point //定义一个Path切片类型，代表一个线段

func (path Path) Distance() float64 {
		sum := 0.0
	for i := range path{
		if if > 0{
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}


func main() {
	p := Point{1,2}
	q := Point{4,6}
	fmt.Println(Distance(p,q)) //调用传统的方法
	fmt.Println(p.Distance(q)) //调用p调用方法
	fmt.Println("------------")

	perim := Path{
		{1,1},
		{5,1},	
		{5,4},
		{1,1},
}

	fmt.Println(perim.Distance())

}


Path是一个命名的slice类型，而不是Point那样的struct类型，然而我们依然可以
为它定义方法。在能够给任意类型定义方法这一点上，
go和很多其他它的面向对象的语言不太一样。


可以给同一包内的任意命名类型定义方法，只要这个命名类型的底层类型(底层类型是指[]Point这个slice,
Path就是命名类型)不是指针或interface


两个Distance方法有不同的类型。他们两个方法之间没有任何关系，尽管Path的Distance方法会在
内部调用Point.Distance方法来计算每个连接邻接点的线段的长度


main函数最后计算三角形的周长


go基于指针对象的方法

假设有两个方法，一个方法的接收者是指针类型，一个方法的接收者是值类型


1.对于值类型的变量和指针类型的变量，两者的区别：
2.如果这两个方法是为了实现一个接口，
3.如果方法是嵌入到其他结构体中的，

