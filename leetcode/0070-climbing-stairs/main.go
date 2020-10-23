/*
70. 爬楼梯
	https://leetcode-cn.com/problems/climbing-stairs/
题目描述
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数。
示例 1：
	输入： 2
	输出： 2
	解释： 有两种方法可以爬到楼顶。
	1.  1 阶 + 1 阶
	2.  2 阶
示例 2：
	输入： 3
	输出： 3
	解释： 有三种方法可以爬到楼顶。
	1.  1 阶 + 1 阶 + 1 阶
	2.  1 阶 + 2 阶
	3.  2 阶 + 1 阶
*/

package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))

	fmt.Println(climbStairs2(1))
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs2(3))
	fmt.Println(climbStairs2(4))
}

/*
斐波那契数列，递归做法，超时
思路：
	如果想要跳到第 n 阶，可以选择从第 n-1 阶跳一步或者从第 n-2 阶跳两步
	由此可知，跳到第 n 阶的方式 f(n) = f(n-1) + f(n-2)
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	} else {
		return climbStairs(n-1) + climbStairs(n-2)
	}
}

/*
递推做法，动态规划
	f(n) = f(n-1) + f(n-2)
*/
func climbStairs2(n int) int {
	s1, s2, s3 := 0, 0, 1
	for i := 1; i <= n; i++ {
		s1, s2 = s2, s3
		s3 = s1 + s2
	}
	return s3
}
