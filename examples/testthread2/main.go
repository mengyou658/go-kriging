package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var n = 1000
	var wg1 sync.WaitGroup
	predictCh1 := make(chan int, n)
	for i := 0; i < n; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			fmt.Println("i start " + strconv.Itoa(i))
			sub(i, &wg1, &predictCh1)
			predictCh1 <- 0
			fmt.Println("i end " + strconv.Itoa(i))
		}(i)
	}
	go func() {
		wg1.Wait()
		close(predictCh1)
	}()
	for range predictCh1 {
	}
}
func sub(idx int, wg1 *sync.WaitGroup, predictCh1 *chan int) {
	var n = 1000
	for i := 0; i < n; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			fmt.Println("sub start " + strconv.Itoa(i) + " " + strconv.Itoa(idx))
			*predictCh1 <- 0
			fmt.Println("sub end " + strconv.Itoa(i) + " " + strconv.Itoa(idx))
		}(i)
	}
}
