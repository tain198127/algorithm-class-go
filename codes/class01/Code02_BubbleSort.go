package main

import "log"

//@ title 冒泡排序
func main() {
	arr := []int{1, 3, 0, -9, 2, 100, -22}
	bubbleSort(arr)
	log.Println(arr)
}

// @title 冒牌排序
// @description 使用冒泡排序
// @arr	待排序参数	[]int
func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	for e := N - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
}
func swap(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
