package main

import "fmt"

/*
62. 不同路径
	https://leetcode-cn.com/problems/unique-paths/
题目描述
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	问总共有多少条不同的路径？
示例 1:
	输入: m = 3, n = 2
	输出: 3
	解释:
	从左上角开始，总共有 3 条路径可以到达右下角。
	1. 向右 -> 向右 -> 向下
	2. 向右 -> 向下 -> 向右
	3. 向下 -> 向右 -> 向右
示例 2:
	输入: m = 7, n = 3
	输出: 28
*/
func main() {
	m := 3
	n := 2
	fmt.Println(uniquePaths(m, n))

	m = 7
	n = 3
	fmt.Println(uniquePaths(m, n))

	m = 1
	n = 1
	fmt.Println(uniquePaths(m, n))
}

/*
动态规划
	要想到达 [i, j] 位置，可以从 [i-1, j] 或 [i, j-1] 处移动一步到达
	定义状态 dp[i][j]
	状态转移方程 dp[i][j] = dp[i-1][j] + dp[i][j-1]
*/
func uniquePaths(m int, n int) int {	//m列，n行
	dp := make([][]int, n)
	for i:=0; i<n; i++ {
		dp[i] = make([]int, m)
	}
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if i==0 && j==0 {
				dp[i][j] = 1
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1]
}