package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1}
	fmt.Println(mergeSort(arr))

	arr = []int{4, 3, 2, 1}
	mergeSort2(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	left, right := arr[0:mid], arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(arr1, arr2 []int) (arr []int) {
	for len(arr1) != 0 && len(arr2) != 0 {
		if arr1[0] < arr2[0] {
			arr = append(arr, arr1[0])
			arr1 = arr1[1:]
		} else {
			arr = append(arr, arr2[0])
			arr2 = arr2[1:]
		}
	}
	if len(arr1) > 0 {
		arr = append(arr, arr1...)
	}
	if len(arr2) > 0 {
		arr = append(arr, arr2...)
	}
	return
}

func mergeSort2(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	mid := length / 2
	left, right := append([]int{}, arr[0:mid]...), append([]int{}, arr[mid:]...)
	mergeSort2(left)
	mergeSort2(right)
	// left 和 right 归并填入 nums
	p1, p2 := 0, 0
	for i := range arr {
		if p1 < len(left) && (p2 == len(right) || left[p1] <= right[p2]) {
			arr[i] = left[p1]
			p1++
		} else {
			arr[i] = right[p2]
			p2++
		}
	}
}
