package main

import  "fmt"

type data struct{
	*sync.Mutex // *Mutex
}


func (d data) test(s string){
	d.Lock()
	defer d.Unlock()


	for i:=0;i<5;i++{
		fmt.Println(s,i)
		time.Sleep(time.Second)
	}
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	d := data{new(sync.Mutex)} // 初始化
	go func(){
		defer wg.Done()
		d.test("read")
	}()
	go func(){
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
}