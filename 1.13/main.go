package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Println("Введите два числа")
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)

	fmt.Printf("a=%d, b=%d\n\n", a, b)
	fmt.Println("Свапаем....")
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("После свапа a=%d, b=%d", a, b)
}
