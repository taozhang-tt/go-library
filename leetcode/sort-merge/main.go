package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1}
	fmt.Println(mergeSort(arr))
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