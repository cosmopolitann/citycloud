package main

import (
	"fmt"
	"sort"
)

func main(){
	//  map  就是 散列hash  引用类型  必须要初始化
	s:=map[int]int{}   //这是已经初始化了
	var s1 map[int]int//zhe这个 没有 初始化
	fmt.Println(s,s1)

   s[1]=1
   fmt.Println(s)

   l:=make(map[int]int,10)
   l[1]=1000

l[2]=2000
l[3]=3000
l[4]=4000
fmt.Println(l)
	var s123 []int

for k,v:=range l{
	fmt.Println(k,v)
	s123=append(s123,k)
}

sort.Ints(s123)
for _,v:=range s123{
	fmt.Println(v,l[v])
}






}
