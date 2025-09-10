package main

import (
	"fmt"
	"os"
)

func main() {
	temp := make([]float64, 0, 8)
	var buf float64
	mp := make(map[int][]float64)
	fmt.Println("Введите 8 чисел")
	for range 8 {
		fmt.Fscan(os.Stdin, &buf)
		temp = append(temp, buf)
	}
	fmt.Println(temp)
	for _, num := range temp {
		key := int(num/10) * 10
		sl, ok := mp[key]
		if !ok {
			mp[key] = []float64{num}
		} else {
			sl = append(sl, num)
			mp[key] = sl
		}
	}
	fmt.Println(mp)
}
