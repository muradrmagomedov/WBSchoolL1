package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(cut(arr, 1))
}

func cut(arr []int, i int) []int {
	if i >= len(arr) || i < 0 {
		return arr
	}
	result := make([]int, len(arr)-1)
	result = append(result, arr[:i]...)
	result = append(result, arr[i+1:]...)
	return result
}
