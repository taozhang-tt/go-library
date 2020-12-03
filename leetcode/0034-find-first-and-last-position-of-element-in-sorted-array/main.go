package main
/*
34. 在排序数组中查找元素的第一个和最后一个位置
	https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
题目描述
	给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
	如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：
	你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
示例 1：
	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]
示例 2：
	输入：nums = [5,7,7,8,8,10], target = 6
	输出：[-1,-1]
示例 3：
	输入：nums = [], target = 0
	输出：[-1,-1]
提示：
	0 <= nums.length <= 105
	-109 <= nums[i] <= 109
	nums 是一个非递减数组
	-109 <= target <= 109
*/
import (
	"fmt"
)

func main() {

	nums := []int{1}
	target := 1
	fmt.Println(searchRange(nums, target))

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8
	fmt.Println(searchRange(nums, target))

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println(searchRange(nums, target))

	nums = []int{}
	target = 0
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) == 0 {
		return ans
	}
	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			ans = []int{mid, mid}
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l > r {
		return ans
	} else {
		l, r = mid, mid
		for l >= 0 && nums[l] == target {
			l--
		}
		ans[0] = l + 1
		for r < len(nums) && nums[r] == target {
			r++
		}
		ans[1] = r - 1
	}
	return ans
}
