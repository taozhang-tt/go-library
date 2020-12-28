package main

import "fmt"
/*
剑指 Offer 10- I. 斐波那契数列
	https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
题目描述
	写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：
	F(0) = 0,   F(1) = 1
	F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
	斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
示例 1：
	输入：n = 2
	输出：1
示例 2：
	输入：n = 5
	输出：5
提示：
	0 <= n <= 100
*/
func main() {
	n := 45
	fmt.Println(fib(n))

	n = 95
	fmt.Println(fib(n))

	n = 2
	fmt.Println(fib(n))

	n = 6
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	f0, f1, f2 := 0, 1, 0
	for i := 2; i <= n; i++ {
		f2 = (f0 + f1)%1000000007
		f0, f1 = f1, f2
	}
	return f2
}
