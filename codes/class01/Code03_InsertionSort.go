package main

import "log"

//@ title 冒泡排序
func main() {
	arr := []int{1, 3, 0, -9, 2, 100, -22}
	insertionSort(arr)
	log.Println(arr)
}

// @title 插入排序
// @description 使用插入排序
// @arr	待排序参数	[]int
func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 不止1个数
	N := len(arr)
	for i := 1; i < N; i++ { // 0~i有序
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

// @title 要保证i和j不是一个位置，否则出错
func swap(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
