package main

import "fmt"

func main() {
	s := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 15}
	fmt.Println(s)

	s1 := s[:]

s2:=append(s1[0:1],s1[2:]...)
fmt.Println(s2)


}
