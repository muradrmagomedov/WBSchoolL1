package main

import (
	"fmt"
	"os"
)

func main() {
	var n, i, b int64
	fmt.Println("Введите число:")
	fmt.Fscan(os.Stdin, &n)
	fmt.Println("Введите номер бита для замены")
	fmt.Fscan(os.Stdin, &i)
	fmt.Println("Введите бит, на который заменяем: 1 или 0")
	fmt.Fscan(os.Stdin, &b)
	fmt.Println(n, i, b)

	if b == 1 {
		n = n | (1 << i)
	}
	if b == 0 {
		n = n &^ (1 << i)
	}
	fmt.Println("Получилось число", n)
}
