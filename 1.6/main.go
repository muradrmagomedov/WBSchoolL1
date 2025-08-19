package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"runtime"
	"time"
)

func main() {
	//выход по условию
	go func() {
		done := false
		for !done {
			time.Sleep(500 * time.Millisecond)
			check := rand.IntN(5)
			if check == 2 {
				done = true
			}
		}
	}()

	//Выход через контекст
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("working")
			}
		}
	}()
	cancel()

	//через GoExit
	go func() {
		defer fmt.Println("Stoping...")
		runtime.Goexit()
	}()

	//через канал уведомлений
	done := make(chan int)
	go func(done chan int) {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("Working")
			}

		}
	}(done)
	done <- 1
}
