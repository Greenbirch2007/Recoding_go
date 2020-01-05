package main

import (
	"fmt"
	"p_go"
)

func main()  {
	p := p_go.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(200)
	fmt.Println(p)
	fmt.Println(p.Name,"age=",p.GetAge(),"sal=",p.GetSal())
}