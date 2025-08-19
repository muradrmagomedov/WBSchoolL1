package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	for i := range 10 {
		wg.Add(1)
		go func(ctx context.Context, i int, wg *sync.WaitGroup) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Goroutine %d: Stopping\n", i+1)
					wg.Done()
					return
				case <-time.After(500 * time.Millisecond):
					fmt.Printf("Goroutine %d: Working\n", i+1)
				}
			}
		}(ctx, i, &wg)
	}
	fmt.Println(<-sigs)
	cancel()
	wg.Wait()
	fmt.Println("All goroutines stoped...")
}
