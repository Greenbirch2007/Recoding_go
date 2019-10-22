package main


import "fmt"


func main(){
	var a int =21
	var b int =19

	if (a==b){
		fmt.Printf("第一行 -a 等于b \n")
	}else{
		fmt.Printf("第一行 -a 不等于b \n")
	}
	if(a <b ){
		fmt.Printf("第二行　-a 小于b\n")
	}else{
		fmt.Printf("第二行　-a 不小于b\n")
	}
	if (a>b){
		fmt.Printf("第三行　-a 大于b\n")
	}else{
		fmt.Printf("第三行　-a 不大于b \n")
	}
	// let's change value of a and b

	a = 5
	b = 20
	if (a<=b){
		fmt.Printf("第四行　-a 小于等于b \n")
	} 
	if (b>=a){
		fmt.Printf("第五行 - b 大于等于ａ \n")
}
}


