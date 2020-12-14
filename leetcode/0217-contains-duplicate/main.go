package main

import (
	"fmt"
	"sort"
)

/*
217. 存在重复元素
	https://leetcode-cn.com/problems/contains-duplicate/
题目描述：
	给定一个整数数组，判断是否存在重复元素。
	如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
示例 1:
	输入: [1,2,3,1]
	输出: true
示例 2:
	输入: [1,2,3,4]
	输出: false
示例 3:
	输入: [1,1,1,3,3,4,3,2,4,2]
	输出: true
*/
func main() {
	nums := []int{1,2,3,1}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1,2,3,4}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate(nums))


	nums = []int{1,2,3,4}
	fmt.Println(containsDuplicate2(nums))

	nums = []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate3(nums))

}

//hash表
func containsDuplicate(nums []int) bool {
	cnt := make(map[int]int)
	for _, item := range nums {
		if cnt[item] > 0 {
			return true
		}
		cnt[item]++
	}
	return false
}

//排序，练习一下快排
func containsDuplicate2(nums []int) bool {
	quickSort(nums, 0, len(nums)-1)
	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

//快排
func quickSort(nums []int, left, right int) {
	if left >  right {
		return
	}
	l, r, mid := left, right, nums[left]
	for l < r {
		for l<r && nums[r] > mid {
			r--
		}
		if l<r {
			nums[l] = nums[r]
			l++
		}
		for l<r && nums[l] < mid {
			l++
		}
		if l<r {
			nums[r] = nums[l]
			r--
		}
	}
	nums[l] = mid
	quickSort(nums, left, l-1)
	quickSort(nums, l+1, right)
}

//使用sort包的排序方法，为什么比我的快排快那么多啊！！
func containsDuplicate3(nums []int) bool {
	sort.Ints(nums)
	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}