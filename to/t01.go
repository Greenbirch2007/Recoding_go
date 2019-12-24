package main

import (
	"errors"
)

var errByZero = errors.New("除数不能为０")
func div(a,b int)(r int,err error){
	if b== 0{
		return 0,errByZero
	}

	return a/b,nil
}