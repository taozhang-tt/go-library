package main

import "fmt"

func main() {
	nums := []int{9, 4, 3, 5, 6, 1, 2}
	BubbleSort(nums)
	fmt.Println(nums)
}

//1，比较相邻的两个元素，如果前一个元素比后一个元素大，则交换它们
//2，对每一对元素重复1，直到最后一对，这样一个循环下来，最大对元素会被移动到最后一个
//3，针对所有元素重复步骤 1、2，需要重复 len（数组）- 1 次
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}