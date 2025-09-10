package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите строку для проверки")
	var str string
	fmt.Fscan(os.Stdin, &str)
	fmt.Println(checkUniqueLetters(str))
}

func checkUniqueLetters(str string) bool {
	str = strings.ToLower(str)
	arr := []rune(str)
	mp := make(map[rune]int)
	for _, v := range arr {
		_, ok := mp[v]
		if ok {
			return false
		}
		mp[v]++
	}
	return true
}
