package main

import (
	"log"
	"sort"
)

//@ title 判断某个数是否存在
func main() {
	arr := []int{1, 3, 0, -9, 2, 100, -22}
	test := 0
	sort.Ints(arr)
	isExists := exist(arr, test)
	log.Println(isExists)
}

// @title 判断某个数是否存在
// @description 判断某个数是否存在
// @arr	待查询数组	[]int
// @num	数			int
func exist(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	L := 0
	R := len(arr) - 1
	mid := 0
	for L < R {
		mid = L + (R-L)>>1
		if arr[mid] == num {
			return true
		} else if arr[mid] > num {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return arr[L] == num
}
