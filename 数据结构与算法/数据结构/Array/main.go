package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [5]int{1,2,3,4,5}
	arr3 := [5]int{1,2,3}
	arr4 := [5]int{4,1}
	arr5 := [...]int{1,2,3,4,5}
	arr6 := [...]int{8:1}
	arr7 := [...]int{}
	fmt.Println(arr1,arr2,arr3,arr4,arr5,arr6,arr7)
}
