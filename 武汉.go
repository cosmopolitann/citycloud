package main

import "fmt"

func main(){
	//闭包

	F(1,2)
   func(){
	fmt.Println("hello world")
}()


}


var F= func(x,y int) int{
	s:=x+y
	fmt.Println(s)
	return  s
}