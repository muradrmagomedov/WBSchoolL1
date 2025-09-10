package main

import (
	"fmt"
	"sync"
	"time"
)

type Count struct {
	Value int
}

func main() {
	count := Count{0}
	mx := sync.Mutex{}
	for range 5 {
		go func(mx *sync.Mutex) {
			for range 10 {
				mx.Lock()
				count.Value++
				mx.Unlock()
			}
		}(&mx)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(count.Value)
}
