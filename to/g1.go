package main

func GetValue() int  {
	return 1
}

func main()  {
	i := GetValue()
	switch i.(type){
	case int:
	}
}