package main

import "log"

func main() {
	arr := []int{6, 3, -3, 9, 6}

	selectionSort(arr)
	log.Println(arr)
}

// @title 选择排序
// @description 进行选择排序
// @ arr	待排序数组	[]int
func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	//定义数组长度
	N := len(arr)
	for i := 0; i < N-1; i++ {
		minIdx := i
		for j := i + 1; j < N; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		swap(arr, i, minIdx)
	}
}
func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
