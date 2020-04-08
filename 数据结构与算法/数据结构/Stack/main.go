package main

import "fmt"

func main() {

	type life []interface{}

	data := []int{1,23,4}
	fmt.Println(data)

	data2 := &life{
		"s",
		2,3,
	}

	fmt.Println(data2)
	fmt.Println(*data2)
}
