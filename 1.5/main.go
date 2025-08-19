package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Ввелите время работы программы в секундах:")
	var n int
	fmt.Fscan(os.Stdin, &n)
	ch := make(chan int)
	defer close(ch)
	var count int
	go func() {
		for {
			count++
			ch <- count
		}
	}()
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	time.Sleep(time.Second * time.Duration(n))
}
