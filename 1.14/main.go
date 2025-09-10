package main

import "fmt"

func main() {
	var n interface{}

	switch n.(type) {
	case bool:
		fmt.Println("Type is bool")
	case string:
		fmt.Println("Type is string")
	case int:
		fmt.Println("Type is int")
	case chan int:
		fmt.Println("Type is chan")
	}
}
