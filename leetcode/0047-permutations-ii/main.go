package main

import (
	"fmt"
	"sort"
)
/*
47. 全排列 II
	https://leetcode-cn.com/problems/permutations-ii/
题目描述
	给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
示例 1：
	输入：nums = [1,1,2]
	输出：
	[[1,1,2],
	[1,2,1],
	[2,1,1]]
示例 2：
	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]	
提示：
	1 <= nums.length <= 8
	-10 <= nums[i] <= 10
*/
func main() {
	// nums := []int{1, 2, 1, 2, 3}
	// fmt.Println(permuteUnique(nums))

	nums := []int{-1, 2, -1, 2, 1, -1, 2, 1}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) (ans [][]int) {
	if len(nums) <= 0 {
		return
	}
	ans = make([][]int, 0)
	backtrace := func(ansItem, residue []int) {}
	backtrace = func(ansItem, residue []int) {
		if len(residue) == 0 {
			ans = append(ans, ansItem)
			return
		}
		for i := 0; i < len(residue); i++ {
			if i > 0 && residue[i] == residue[i-1] {
				continue
			}
			ansItemNew := append(ansItem[0:len(ansItem):len(ansItem)], residue[i])
			residueNew := append(residue[0:i:i], residue[i+1:]...)
			backtrace(ansItemNew, residueNew)
		}
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ansItem := []int{nums[i]}
		resuide := append(nums[0:i:i], nums[i+1:]...)
		backtrace(ansItem, resuide)
	}
	return
}
