package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 5}
	target := 4
	fmt.Println(binary(nums, target))
	fmt.Println(binaryLeft(nums, 2))
	fmt.Println(binaryRight(nums, 2))
}

//朴素的二分查找
func binary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 //等同于 (left+right)/2 但是可以防止 left 和 right 太大，相加导致越界
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//返回最左侧的索引
func binaryLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

//返回最右侧的索引
func binaryRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= len(nums) || nums[right] != target {
		return -1
	}
	return right
}
