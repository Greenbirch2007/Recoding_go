package main

import "fmt"

type Steretype int


const (
	TypicalNoob Steretype = iota //0
	TypicalHipster  // 1
	TypicalUnixWizard // 2
	TypicalStartupFounder //  3
)




/*

@brief 测试　const形参匹配类型

@param int

@returns

	string 类型


*/



func CountAllTheThings(i int) string{
	// 返回一个字符串　　Sprintf是创建一个字符串
	return fmt.Sprint("there are %d things",i)
}

func SoSaethThe(character Stereotype) string {
	var s string
	switch character{
	case TypicalNoob:
		s = "I'm a confused ninja rockstar."
	case TypicalHipster:
		s = "~~~~~~"
	case TypicalUnixWizard:
		s = "sudo grep awk sed %!#"
	case TypicalStartupFounder:
		s = "exploit compelling convergence to ~~"

}
	return s
}


func main()
{
	n := TypicalHipster
	// 虽然自定义类型，但是传参也是会报错...
	fmt.Println(CountAllTheThings(n))
	i := 2
	fmt.Println(SoSayethThe(i))
}



