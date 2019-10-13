package main

import "fmt"

func main(){
	 for i:=0;i<10	 ;i++  {
		 if i ==5 || i==9{

		 	//break
		 	continue
		 }
		 fmt.Println("i的值是",i)

	 }
 	fmt.Println("已经跳出来for循环了===over")

 }