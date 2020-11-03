/**
941. 有效的山脉数组
	https://leetcode-cn.com/problems/valid-mountain-array/
题目描述：
	给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。
	让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
	A.length >= 3
	在 0 < i < A.length - 1 条件下，存在 i 使得：
	A[0] < A[1] < ... A[i-1] < A[i]
	A[i] > A[i+1] > ... > A[A.length - 1]
示例 1：
	输入：[2,1]
	输出：false
示例 2：
	输入：[3,5,5]
	输出：false
示例 3：
	输入：[0,3,2,1]
	输出：true
示例 4：
	输入：[0, 1, 2, 3, 4, 8, 9, 10, 11, 12, 11]
	输出：true
示例 5：
	输入：[0,1,2,3]
	输出：false
提示：
	0 <= A.length <= 10000
	0 <= A[i] <= 10000 
*/
package main

import "fmt"


func main() {
	var A []int
	A = []int{2, 1}
	fmt.Println(validMountainArray(A))

	A = []int{3, 5, 5}
	fmt.Println(validMountainArray(A))

	A = []int{0, 1, 2, 3, 4, 8, 9, 10, 11, 12, 11}
	fmt.Println(validMountainArray(A))

	A = []int{0, 1, 2, 3, 4}
	fmt.Println(validMountainArray2(A))
}

/**
双指针做法
*/
func validMountainArray(A []int) bool {
	if len(A) < 3 {	//特判
		return false
	}
	i, j := 1, len(A)-2	//初始化两个指针分别为 第 1 个和倒数第 2 个；[0-i]递增，[j-最后一个]递减
	for i <= j {	//循环向内侧挪动指针
		if A[i] <= A[i-1] || A[j] <= A[j+1] {	
			return false
		}
		if A[i] < A[j] {	//如果元素 i 更小，i指针向内侧挪动一个位置，否则 j 指针向右挪动一个位置
			i++
		} else {
			j--
		}
	}
	return true
}

/**
官方题解代码
*/
func validMountainArray2(a []int) bool {
    i, n := 0, len(a)
    // 递增扫描
    for ; i+1 < n && a[i] < a[i+1]; i++ {
    }
    // 最高点不能是数组的第一个位置或最后一个位置
    if i == 0 || i == n-1 {
        return false
    }
    // 递减扫描
    for ; i+1 < n && a[i] > a[i+1]; i++ {
    }
    return i == n-1
}
