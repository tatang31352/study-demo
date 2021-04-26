package main

import (
	"fmt"
)
var ch2 =  make(chan bool)

func main() {
	ch := make(chan int)
	go write(ch)
	go read(ch)
	fmt.Println(<-ch2)
}


func write(ch chan int){
	for i:= 0;i < 10;i++{
		ch <- i
	}
	close(ch)
}

func read(ch chan int) {

	for{
		v,ok := <-ch
		if ok{
			fmt.Println(v)
		}else{
			ch2 <- true
		}
	}

}