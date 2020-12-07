package main

import (
	"fmt"
)

/*
861. 翻转矩阵后的得分
	https://leetcode-cn.com/problems/score-after-flipping-matrix/
题目描述
	有一个二维矩阵 A 其中每个元素的值为 0 或 1 。
	移动是指选择任一行或列，并转换该行或列中的每一个值：将所有 0 都更改为 1，将所有 1 都更改为 0。
	在做出任意次数的移动后，将该矩阵的每一行都按照二进制数来解释，矩阵的得分就是这些数字的总和。
	返回尽可能高的分数。
示例：
	输入：[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
	输出：39
	解释：
	转换为 [[1,1,1,1],[1,0,0,1],[1,1,1,1]]
	0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
*/
func main() {
	A := [][]int{
		{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0},
	}
	fmt.Println(matrixScore(A))

	A = [][]int{
		{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0},
	}
	fmt.Println(matrixScore2(A))
}

/*
贪心算法
	要想使每一行表示的二进制数最大，就要让每一行尽量靠左的数字为1
	先进行行的翻转，对于每一行，若第一个数字不为 0，则翻转该行
	再进行列的翻转，对于每一列，如果 0 的个数大于 1，则翻转
*/
func matrixScore(A [][]int) int {
	rowLen, colLen, ans := len(A), len(A[0]), 0
	for i := 0; i < rowLen; i++ { //进行行翻转
		if A[i][0] == 0 {
			for j := 0; j < len(A[i]); j++ {
				A[i][j] ^= 1
			}
		}
	}
	for j := 0; j < colLen; j++ { //进行列翻转
		cnt := 0 //统计1的个数
		for i := 0; i < rowLen; i++ {
			cnt += A[i][j] & 1
		}
		if cnt*2 < rowLen {
			for i := 0; i < rowLen; i++ {
				A[i][j] ^= 1
			}
		}
	}
	for i := 0; i < rowLen; i++ {
		rowSum := 0
		for j := 0; j < colLen; j++ {
			rowSum = rowSum*2 + A[i][j]
		}
		ans += rowSum
	}
	return ans
}

/*
官方题解
	接着上面的思路，其实不用真的翻转每一行每一列，而只需要计算每一个元素的最终结果的贡献即可
	假设矩阵有 m 行、n 列
	对于最左侧的那一列，最终的结果肯定是令其每个元素都为1，其贡献为  m * 2 的 （n-1）次幂
	对于其它的任一列，分别统计 0 和 1 的个数（要注意为了使每一行的第一个元素为1而进行过的翻转），数量较多的那个记为 k 即为最终的翻转结果，对结果贡献为 k * 2 的 (n-j-1) 次幂
*/
func matrixScore2(a [][]int) int {
	m, n := len(a), len(a[0])
	ans := 1 << (n - 1) * m //计算第一列的共享
	for j := 1; j < n; j++ {
		ones := 0	//统计 0 或 1 的个数（视第一个元素而定）
		for _, row := range a {
			if row[j] == row[0] {
				ones++
			}
		}
		if ones < m-ones {
			ones = m - ones
		}
		ans += 1 << (n - 1 - j) * ones
	}
	return ans
}
