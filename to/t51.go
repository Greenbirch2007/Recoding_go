package main

import ("fmt"
	"os"
)


func main(){
	myFile := "path"
	fout,err := os.Create(myFile)
	// fout,err := os.OpenFile(myFile,os.O_CREATE,0644)
	if err != nil {
		fmt.Println(err)
		return 
}
	for i := 0; i <10 ;i++ {
		outstr := fmt.Sprintf("%s:%d\n","Hello world",i)
		fout.WriteString(outstr)
		fout.Write([]byte("abcd\n"))}
		fout.Close()
}



