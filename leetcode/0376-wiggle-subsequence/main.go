package main

import (
	"fmt"
)

/*
376. 摆动序列
	https://leetcode-cn.com/problems/wiggle-subsequence/
题目描述
	如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。
	例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。相反, [1,4,7,2,5] 和 [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
	给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。
示例 1:
	输入: [1,7,4,9,2,5]
	输出: 6
	解释: 整个序列均为摆动序列。
示例 2:
	输入: [1,17,5,10,13,15,10,5,16,8]
	输出: 7
	解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。
示例 3:
	输入: [1,2,3,4,5,6,7,8,9]
	输出: 2
*/
func main() {
	nums := []int{10, 9, 6, 3, 4}
	fmt.Println(wiggleMaxLength(nums))

	nums = []int{1, 7, 4, 9, 2, 5}
	fmt.Println(wiggleMaxLength2(nums))

	nums = []int{1, 7, 4, 9, 2, 5}
	fmt.Println(wiggleMaxLength3(nums))
}
/*
动态规划
	状态定义：
		up[i] 表示 [0, i] 这个区间内最长的摆动序列，且最后两个元素是上升的
		down[i] 表示 [0, j] 这个区间内最长的摆动序列，且最后两个元素是下降的
	状态转换：
		if nums[i] > nums[i-1]
*/
func wiggleMaxLength(nums []int) int {
    n := len(nums)
    if n < 2 {
        return n
    }
    up := make([]int, n)
    down := make([]int, n)
    up[0] = 1
    down[0] = 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            up[i] = max(up[i-1], down[i-1]+1)
            down[i] = down[i-1]
        } else if nums[i] < nums[i-1] {
            up[i] = up[i-1]
            down[i] = max(up[i-1]+1, down[i-1])
        } else {
            up[i] = up[i-1]
            down[i] = down[i-1]
        }
    }
    return max(up[n-1], down[n-1])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func wiggleMaxLength2(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	up, down := 1, 1
	for i:=1; i<length; i ++ {
		if nums[i] > nums[i-1] {
			up = down+1
		} else if nums[i] < nums[i-1] {
			down = up+1
		}
	}
	if up > down {
		return up
	}
	return down
}

/*
波峰波谷法
	寻找有多少个波峰和波谷就好了！！！
*/
func wiggleMaxLength3(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	ans, trend, first := 1, false, true
	for i:=1; i<length; i++ {
		if nums[i] != nums[i-1] {
			if first {
				ans++
				trend = nums[i] > nums[i-1]
				first=false
				continue
			}
			if nums[i] > nums[i-1] != trend {
				ans++
				trend = nums[i] > nums[i-1]
			}
		}
	}
	return ans
}