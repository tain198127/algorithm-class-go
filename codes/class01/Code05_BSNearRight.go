package main

import (
	"log"
	"sort"
)

func main() {
	arr := []int{1, 3, 0, 0, -9, 2, 100, -22}
	sort.Ints(arr)
	test := 3
	ret := nearestIndex(arr, test)
	log.Println(ret)
}

// @title 找到最右位置
// @arr 数组	[]int
// @value		int
// @return		int
func nearestIndex(arr []int, val int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	L := 0
	R := len(arr) - 1
	index := -1 //记录最优的对号
	for L <= R {
		// 其实就是 (R+L)/2 由于担心R+L越界 因此写成R/2+L/2，而L/2可以写成 L-L/2
		//从而写成R/2+L-L/2，转换后变成L+(R/2-L/2)=>L+(R-L)/2，由于除法比较慢，可以写成>>1
		//最终变成L+((R-L)>>1)
		// 主要是用来装逼
		mid := L + ((R - L) >> 1)
		if arr[mid] <= val {
			index = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return index
}
