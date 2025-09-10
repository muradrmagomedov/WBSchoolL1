package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 5, 7, 8, 10, 11, 12, 15, 17}
	fmt.Println(binSearch(arr, 17))
}

func binSearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == num {
			return mid
		} else if num > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
