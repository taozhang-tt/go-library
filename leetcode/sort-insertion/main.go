package main

import "fmt"

func main() {
	var nums []int
	nums = []int{9, 4, 3, 5, 6, 1, 2}
	insertSort(nums)
	fmt.Println(nums)

	nums = []int{9, 4, 3, 5, 6, 1, 2}
	InsertSort(nums)
	fmt.Println(nums)
}

//1，从第一个元素开始，该元素可以认为已经被排序；
//2，取出下一个元素，在已经排序的元素序列中从后向前扫描；
//3，如果该元素（已排序）大于新元素，将该元素移到下一位置；
//4，重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
//5，将新元素插入到该位置后；
//6，重复步骤2~5
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		var j = i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
}

func insertSort(arr []int) {
	for i:=1; i<len(arr); i++ {
		curr, j := arr[i], i-1
		for j >= 0 && arr[j] > curr {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
	}
}