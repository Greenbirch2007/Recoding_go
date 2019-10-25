package main


import "fmt"


func main(){
	var n [10]int //n是一个长度为１０的数组
	var i,j int


	// 为数组 n初始化元素


	for i=0;i<10;i++{
		n[i] = i+100 //设置元素为 i + 100
	}
	
	//输出每个数组元素的值


	for j =0;j=10; j++{
		fmt.Printf("Element[%d] = %d \n",j,n[j])

}
}

多维数组


go语言支持多维数组，以下为常用的多维数组声明方式


var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type


以下实例声明了三维的整型数组


var threedim [5][10][4]int


a = [3][4]int{
	{0,1,2,3}, //第一行索引为0
	{4,5,6,7}, // 第二行索引为１
	{8,9,10,11} // 第三行索引为2
}



