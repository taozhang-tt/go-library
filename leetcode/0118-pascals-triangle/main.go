package main

import "fmt"

/*
118. 杨辉三角
	https://leetcode-cn.com/problems/pascals-triangle/
题目描述
	给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
示例:
	输入: 5
	输出:
	[
		 [1],
		[1,1],
	   [1,2,1],
	  [1,3,3,1],
	 [1,4,6,4,1]
	]
 */
func main()  {
	numRows := 5
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i:=0; i<numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1
		for j:=1; j < i; j++ {
			ans[i][j] = ans[i-1][j]+ans[i-1][j-1]
		}
	}
	return ans
}