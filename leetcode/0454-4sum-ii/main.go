package main
/*
454. 四数相加 II
	https://leetcode-cn.com/problems/4sum-ii/
题目描述：
	给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
	为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。
输入:
	A = [ 1, 2]
	B = [-2,-1]
	C = [-1, 2]
	D = [ 0, 2]
输出:
	2
解释:
	两个元组如下:
	1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
	2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/
import "fmt"

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}

/*
暴力解法
*/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	mapAB, length, ret := make(map[int]int, 0), len(A), 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			mapAB[A[i]+B[j]]++
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if cnt, ok := mapAB[-(C[i] + D[j])]; ok {
				ret += cnt
			}
		}
	}
	return ret
}
