package main

func throwsPanic(f func())(b bool)  {
defer func() {
	if x := recover();x != nil{
		b = true
	}
}()
f()
return 
}