package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	wg.Add(len(nums))
	for _, v := range nums {
		go func(v int) {
			defer wg.Done()
			fmt.Println(v * v)
		}(v)
	}
	wg.Wait()
}
