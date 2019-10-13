package main

import (
	"fmt"
)

type S struct {
	name string
	age  int

}


func main(){

	fmt.Println("nihao")
  //  beego.Error("nihao")
s:=&S{
	name: "123",
	age:  1,
}
s.name="90"

fmt.Println(s)


}
