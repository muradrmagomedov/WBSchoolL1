package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ch1, ch2 chan int) {
		defer wg.Done()
		defer close(ch2)
		for i := range ch1 {
			ch2 <- i * 2
		}
	}(ch1, ch2)
	go func(ch2 chan int) {
		defer wg.Done()
		for i := range ch2 {
			fmt.Println(i)
		}
	}(ch2)
	for i := range 10 {
		ch1 <- i
	}
	close(ch1)
	wg.Wait()
}
