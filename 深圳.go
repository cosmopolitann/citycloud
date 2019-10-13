package main

import "fmt"

func main() {
	array := [...]int{1, 3, 5, 7, 8}
	var sum int
	for _, v := range array {
		sum += v
		//  sum = sum +v
	}
	fmt.Println(sum)

	for i:=0;i<len(array) ;i++  {

		for j:=0;j<len(array) ;j++  {
			if array[i]+array[j]==8{
				fmt.Println(i,j)
			}

		}

	}
	var l []int
	fmt.Print(l==nil)

fmt.Println(l)

}
