package main

import "fmt"

func main() {
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	fmt.Println(p)
    m:=*p
    fmt.Println(m)

    a:=new(int)
    *a=100
    fmt.Println(*a)

    var  a1 *int
  //  *a1=10
    fmt.Println(a1 ==nil)


    //  new  和  make 的区别
    /*
    都是用来初始化类型  申请内存的  返回对应的类型指针

    make  初始化 chan  slice  map    返回的是类型本身
     */

}
