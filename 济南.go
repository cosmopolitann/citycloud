package main

import "fmt"

func main(){


	var s =make([]map[int]int,2)
	s[0]=make(map[int]int,0)
	s[0][1]=1
	s[0][2]=2
	s[0][3]=3
	fmt.Println(s)
fmt.Println(s[1])

}
