/*
283. 移动零
	https://leetcode-cn.com/problems/move-zeroes/
tag:
	数组
题目描述
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
示例:
	输入: [0,1,0,3,12]
	输出: [1,3,12,0,0]
说明:
	必须在原数组上操作，不能拷贝额外的数组。
	尽量减少操作次数。
*/
package main

import "fmt"

func main() {
	var nums = []int{0,1,0,3,12, 0}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0,1,0,3,12, 0}
	moveZeroes2(nums)
	fmt.Println(nums)
}

/*
暴力法
	指针 i 遍历每个元素
		if (nums[i] != 0)
			继续遍历下一个
		else
			指针 j 从 i+1 个位置开始遍历剩下的元素，寻找一个不为 0 的元素与 nums[i] 交换
*/
func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

/*
指针 i 遍历每个元素
指针 j 记录为 0 的元素位置
如果元素 i 不为 0，交换 元素 j、i，
i++, j++
*/
func moveZeroes2(nums []int) {
	j := 0;
	for i:=0; i<len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
