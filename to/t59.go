package main

import "fmt"
import "encoding/json"
import "log"



type Movie struct {
	Title string
	Year int 'json:released'
	Color bool 'json:"color.omitempty"'
	Actors []string
}


var movies =[]Movie{
	{Title:"战狼2",Year:2017,Color:true,
		Actors:[]string{"吴京","吴刚"}},
}


func main(){
	json_str, err := json.MarshalIndent(movies,"","  " )
	if err != nil {
		log.Fatal(err)
		return }

	fmt.Printf("%s\n",json_str)
}

