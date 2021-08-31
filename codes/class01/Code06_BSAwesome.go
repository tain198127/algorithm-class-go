package main

import (
	"log"
)

func main() {
	arr := []int{1, 3, 0, 0, -9, 2, 100, -22, 99}
	ret := getLessIndex(arr)
	log.Println(ret)
}

// @title 找到最小值的IDX
// @arr 数组	[]int
// @value		int
// @return		int
func getLessIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	N := len(arr)
	if N == 1 || arr[0] < arr[1] {
		return 0
	}
	if arr[N-1] < arr[N-2] {
		return N - 1
	}
	left := 1
	right := N - 2
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
