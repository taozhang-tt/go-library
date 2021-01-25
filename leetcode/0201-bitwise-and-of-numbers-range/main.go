package main

/*
201. 数字范围按位与
	https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/
题目描述
	给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
示例 1:
	输入: [5,7]
	输出: 4
示例 2:
	输入: [0,1]
	输出: 0
*/
import "fmt"

func main() {
	m := 5
	n := 7
	fmt.Println(rangeBitwiseAnd2(m, n))

	m = 0
	n = 1
	fmt.Println(rangeBitwiseAnd2(m, n))
}

/*
暴力法（超时）
*/
func rangeBitwiseAnd(m int, n int) int {
	ans := m
	for i := m + 1; i <= n; i++ {
		ans &= i
	}
	return ans
}

/*
位运算
	两个数按位与的结果
*/
func rangeBitwiseAnd2(m int, n int) int {
	shift := 0
	for m != n {
		m, n = m>>1, n>>1
		shift++
	}
	return m << shift
}
