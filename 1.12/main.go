package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var buf string
	var arr []string
	mp := make(map[string]int)

	fmt.Println("Сколько слов будете вводить?")
	fmt.Fscan(os.Stdin, &n)
	arr = make([]string, 0, n)
	list := make([]string, 0, n)
	fmt.Println("Введите слова")
	for range n {
		fmt.Fscan(os.Stdin, &buf)
		arr = append(arr, buf)
	}

	for _, v := range arr {
		mp[v]++
	}
	for i, _ := range mp {
		list = append(list, i)
	}
	fmt.Println(list)
}
