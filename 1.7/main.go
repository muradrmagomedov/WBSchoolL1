package main

import (
	"fmt"
	"sync"
)

func main() {
	mp := make(map[int]int)
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	for range 10 {
		wg.Add(1)
		go func(mx *sync.Mutex) {
			defer wg.Done()
			mx.Lock()
			for range 5 {
				mp[1]++
			}
			mx.Unlock()
		}(&mx)
	}
	wg.Wait()
	fmt.Println(mp[1])
}
