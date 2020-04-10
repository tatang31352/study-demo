package main

import "fmt"

func main() {
	
	data := []int{2,3,4,6,7,3,5,6,7,3}
	
	
	fmt.Println(sort(data))
}

func sort(data []int) []int {

	for i := 0;i < len(data) ; i++  {
		for j := i + 1; j < len(data);j++{
			if data[i] < data[j]{
				data[j],data[i] = data[i],data[j]
			}
		}
	}
	return data
}


