package main

/*
18. 四数之和
	https://leetcode-cn.com/problems/4sum/
题目描述
	给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
	注意：答案中不可以包含重复的四元组。
示例：
	给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
	满足要求的四元组集合为：
	[
	[-1,  0, 0, 1],
	[-2, -1, 1, 2],
	[-2,  0, 0, 2]
	]
*/
import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	target := -11
	fmt.Println(fourSum(nums, target))

	nums = []int{0, 0, 0, 0}
	target = -0
	fmt.Println(fourSum(nums, target))
}

/*
排序 + 双指针
具体思路参考 0015
*/
func fourSum(nums []int, target int) [][]int {
	length, ret := len(nums), make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, length-1 //双指针
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					if r == length-1 || nums[r] != nums[r+1] {
						ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					}
					l++
					r--
				}
			}
		}
	}
	return ret
}
