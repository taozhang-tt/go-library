package main

import (
	"fmt"
	"math"
	"sort"
)

/*
1122. 数组的相对排序
	https://leetcode-cn.com/problems/relative-sort-array/
题目描述：
	给你两个数组，arr1 和 arr2，

	arr2 中的元素各不相同
	arr2 中的每个元素都出现在 arr1 中
	对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
示例：
	输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
	输出：[2,2,2,1,4,3,3,9,6,7,19]
提示：
	arr1.length, arr2.length <= 1000
	0 <= arr1[i], arr2[i] <= 1000
	arr2 中的元素 arr2[i] 各不相同
	arr2 中的每个元素 arr2[i] 都出现在 arr1 中
*/
func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	ret := relativeSortArray(arr1, arr2)

	arr1 = []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 = []int{2, 1, 4, 3, 9, 6}
	ret = relativeSortArray2(arr1, arr2)
	fmt.Println(ret)
}

/*
选择排序
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	sortMap := map[int]int{}
	for index, value := range arr2 {
		sortMap[value] = index
	}
	for _, value := range arr1 {
		if _, ok := sortMap[value]; !ok {
			sortMap[value] = math.MaxInt64
		}
	}
	for i := 0; i < len(arr1); i++ {
		minIndex := i
		for j := i + 1; j < len(arr1); j++ {
			if (sortMap[arr1[j]] < sortMap[arr1[minIndex]]) || (sortMap[arr1[j]] == sortMap[arr1[minIndex]] && arr1[j] < arr1[minIndex]) {
				minIndex = j
			}
		}
		arr1[i], arr1[minIndex] = arr1[minIndex], arr1[i]
	}
	return arr1
}

/*
快排
*/
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	sortMap := map[int]int{}
	for index, value := range arr2 {
		sortMap[value] = index
	}
	for _, value := range arr1 {
		if _, ok := sortMap[value]; !ok {
			sortMap[value] = math.MaxInt64
		}
	}
	sort.Slice(arr1, func(i, j int) bool {
		return (sortMap[arr1[i]] < sortMap[arr1[j]]) || (sortMap[arr1[i]] == sortMap[arr1[j]] && arr1[i] < arr1[j])
	})
	return arr1
}
