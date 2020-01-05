//使用package 声明一个包叫做test


package test

import "fmt"

//函数名首字母必须大写，否则包外无法访问

func init()  {
	fmt.Println("test init run")
}


