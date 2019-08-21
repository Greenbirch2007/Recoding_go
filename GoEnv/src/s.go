package main

import (
	"fmt"
	"sort"
)

type StuScore struct {
	name string
	score int

}

type StuScores []StuScore

func (s StuScores) Len() int  {
	return len(s)
}

func (s StuScores)Less(i,j int) bool  {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i,j int )  {
	s[i],s[j] = s[j],s[i]
}

func main()  {
	stus　:= StuScores{
		{"alan",95},
		{"alan",95},
		{"alan",95}
	}
	fmt.Println("Default")
	for _,v := range stus{
		fmt.Println(v.name,":",v.score)
	}
	fmt.Println()

	sort.Sort(stus)
	fmt.Println("Sorted:")
	for _,v := range stud{
		fmt.Println(v.name,":",v.score)
	}
	fmt.Println("IS Sorted?",sort.IsSorted(stud))
}