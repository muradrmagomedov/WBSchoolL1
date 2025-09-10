package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Println("Введите строку для реверса")
	fmt.Fscan(os.Stdin, &str)
	buffer := []rune(str)
	swap := []rune{}
	for i := len(buffer) - 1; i >= 0; i-- {
		swap = append(swap, buffer[i])
	}
	newStr := string(swap)
	fmt.Printf("Новая строка: %s", newStr)
}
