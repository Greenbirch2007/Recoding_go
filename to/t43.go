package main


import "fmt"
import "math"


type Point struct{X,Y float64}


//传统的方法

func Distance(p,q Point) float64{
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}




//这是给struct Point类型定义一个方法

func (p Point) Distance(q Point) float64{
	return math.Hypot(q.X-p.X,q.Y-p.Y)

}


func main(){
	p := Point{1,2}
	q := Point{4,6}
	fmt.Println(Distance(p,q)) // 调用传统的方法
	fmt.Println(p.Distance(q)) // 通过p调用方法

}



上面代码附加的参数p,叫做方法的接收器(receiver),

早期的面向对象语言留下的遗产将调用一个方法称为”向一个对象发送消息“


在go语言中，我们并不会像其它语言那样用this或self作为接收其



