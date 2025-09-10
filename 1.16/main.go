package main

import "fmt"

func main() {
	sl := []int{5, 12, 8, 1, 0, 6, 9, 12, 34, 11, 18, 5, 4, 0, 20}
	sl = quickSort(sl)
	fmt.Println(sl)
}

func quickSort(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}
	var left, right = 0, len(sl) - 1
	pivot := sl[len(sl)/2]
	for left <= right {
		for sl[left] < pivot {
			left++
		}
		for pivot < sl[right] {
			right--
		}
		if left <= right {
			sl[left], sl[right] = sl[right], sl[left]
			left++
			right--
		}
	}
	quickSort(sl[:right+1])
	quickSort(sl[left:])
	return sl
}
