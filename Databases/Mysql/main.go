package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(50000000)
	for i := 0; i < 10000;i++{
		go func() {
			for i := 0;i < 5000;i++{
					http.Get("http://test.com/index/index/update")
					wg.Done()
			}
		}()
	}

	wg.Wait()


	fmt.Println("结束所有")
}
