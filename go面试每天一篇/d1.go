package main

import "fmt"

func main(){
	s := []int{1,2,3,4,5}
	for i,v := range s{
		fmt.Println("v:%d,v_addr:%p,elem_addr:%p\n",v,&v,&s[i])
	}
}