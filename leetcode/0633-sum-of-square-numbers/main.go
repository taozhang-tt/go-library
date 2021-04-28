package main

import (
	"fmt"
	"math"
)

/*
633. 平方数之和
	https://leetcode-cn.com/problems/sum-of-square-numbers/
题目描述
	给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
示例 1：
	输入：c = 5
	输出：true
	解释：1 * 1 + 2 * 2 = 5
示例 2：
	输入：c = 3
	输出：false
示例 3：
	输入：c = 4
	输出：true
*/

func main() {
	c := 1
	fmt.Printf("%v sum-of-square-numbers is %v\n", c, judgeSquareSum(c))

	c = 2
	fmt.Printf("%v sum-of-square-numbers is %v\n", c, judgeSquareSum(c))
}

func judgeSquareSum(c int) bool {
	maxF := math.Sqrt(float64(c))
	min, max, aim := int64(0), int64(maxF), int64(c)
	for min <= max {
		ret := min*min + max*max
		if ret == aim {
			return true
		} else if ret < aim {
			min++
		} else {
			max--
		}
	}
	return false
}
