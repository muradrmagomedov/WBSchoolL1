package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Введите кол-во воркеров")
	var n int
	fmt.Fscan(os.Stdin, &n)
	ch := make(chan int)
	defer close(ch)
	for i := range n {
		go func(i int) {
			for {
				str := <-ch
				fmt.Printf("Горутина %d вывела: %d\n", i, str)
				time.Sleep(time.Second)
			}
		}(i)
	}
	var count int
	for {
		count++
		ch <- count
	}
}
