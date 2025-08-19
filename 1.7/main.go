package main

import "sync"

func main() {
	mp := make(map[int]int)
	mx := sync.Mutex{}
	for range 10 {
		go func(mx *sync.Mutex) {
			map
		}()
	}
}
