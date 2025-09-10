package main

import (
	"fmt"
	"os"
)

func main() {
	sl1 := make([]int, 0, 8)
	sl2 := make([]int, 0, 8)
	list := make([]int, 0, 8)
	mp := make(map[int]int)
	var i, v int
	fmt.Println("Введите размер первого слайса")
	fmt.Fscan(os.Stdin, &i)
	fmt.Println("Введите числа для первого слайса")
	for range i {
		fmt.Fscan(os.Stdin, &v)
		sl1 = append(sl1, v)
	}
	fmt.Println("Введите размер второго слайса")
	fmt.Fscan(os.Stdin, &i)
	fmt.Println("Введите числа для второго слайса")
	for range i {
		fmt.Fscan(os.Stdin, &v)
		sl2 = append(sl2, v)
	}
	for _, v := range sl1 {
		mp[v]++
	}
	for _, v := range sl2 {
		_, ok := mp[v]
		if ok {
			list = append(list, v)
			delete(mp, v)
		}
	}
	fmt.Println(list)
}
