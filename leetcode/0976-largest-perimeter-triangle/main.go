package main
/*
976. 三角形的最大周长
	https://leetcode-cn.com/problems/largest-perimeter-triangle/
题目描述
	给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。
	如果不能形成任何面积不为零的三角形，返回 0。
示例 1：
	输入：[2,1,2]
	输出：5
示例 2：
	输入：[1,2,1]
	输出：0
示例 3：
	输入：[3,2,3,4]
	输出：10
示例 4：
	输入：[3,6,2,3]
	输出：8
*/
import (
	"fmt"
	"sort"
)

func main() {
	A := []int{2, 1, 2}
	fmt.Println(largestPerimeter(A))

	A = []int{1, 2, 1}
	fmt.Println(largestPerimeter(A))

	A = []int{3, 2, 3, 4}
	fmt.Println(largestPerimeter(A))

	A = []int{3, 6, 2, 3}
	fmt.Println(largestPerimeter(A))
}

func largestPerimeter(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	for i := 0; i < len(A)-2; i++ {
		if A[i]-A[i+1] < A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
