package main

import (
	"fmt"
	"sort"
)

func main(){

	s:=make(map[string]int,10)
	s["lily"]=1
	s["kitty"]=2
	s["milliy"]=3
	s["kitlina"]=4
	s["angela"]=5
	s["lida"]=6
	s["nick"]=7
fmt.Println(s)
 var  a []string
for k,_:=range s{
	a=append(a,k)
}
sort.Strings(a)
for _,v:=range a{
	fmt.Println(v,s[v])


}
  delete(s,"lily")




}
