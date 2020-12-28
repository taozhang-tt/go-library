package main

import "fmt"
/*
46. 全排列
	https://leetcode-cn.com/problems/permutations/
题目描述
	给定一个 没有重复 数字的序列，返回其所有可能的全排列。
示例:
	输入: [1,2,3]
	输出:
	[
	[1,2,3],
	[1,3,2],
	[2,1,3],
	[2,3,1],
	[3,1,2],
	[3,2,1]
	]
*/
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))

	nums = []int{1, 2}
	fmt.Println(permute(nums))

	nums = []int{1, 2, 3, 4}
	fmt.Println(permute(nums))
}

func permute(nums []int) (ans [][]int) {
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
			ansItemNew := append(ansItem, residue[i])
			residueNew := append(residue[0:i:i], residue[i+1:]...)
			backtrace(ansItemNew, residueNew)
		}
	}
	for i := 0; i < len(nums); i++ {
		ansItem := []int{nums[i]}
		resuide := append(nums[0:i:i], nums[i+1:]...)
		backtrace(ansItem, resuide)
	}
	return
}
