package main

import "fmt"

/*
321. 拼接最大数
	https://leetcode-cn.com/problems/create-maximum-number/
题目描述
	给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
	求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
	说明: 请尽可能地优化你算法的时间和空间复杂度。
示例 1:
	输入:
	nums1 = [3, 4, 6, 5]
	nums2 = [9, 1, 2, 5, 8, 3]
	k = 5
	输出:
	[9, 8, 6, 5, 3]
示例 2:
	输入:
	nums1 = [6, 7]
	nums2 = [6, 0, 4]
	k = 5
	输出:
	[6, 7, 6, 0, 4]
*/
func main() {
	nums1 := []int{3, 4, 6, 5}
	nums2 := []int{9, 1, 2, 5, 8, 3}
	k := 5
	fmt.Println(maxNumber(nums1, nums2, k))

}

/*
暴力解法
	1. 统计对于数组 nums， 在区间[idx, len(nums)-1]范围内的最大值是多少
	2. 统计数组每个元素值对应的索引 map[int][]int
*/
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	len1, len2, ans := len(nums1), len(nums2), make([]int, 0)
	nums1Max, nums2Max, nums1Map, nums2Map := make([]int, len1), make([]int, len2), make(map[int][]int), make(map[int][]int)
	nums1Max[len1-1], nums2Max[len2-1], nums1Map[nums1[len1-1]], nums2Map[nums2[len2-1]] = nums1[len1-1], nums2[len2-1], append(nums1Map[nums1[len1-1]], len1-1), append(nums2Map[nums2[len2-1]], len2-1)
	for i := len1 - 2; i >= 0; i-- { //统计nums1 在区间 [i, len(nums)-1] 区间的最大值
		nums1Map[nums1[i]] = append(nums1Map[nums1[i]], i) //统计nums1每个值出现的位置
		if nums1[i] > nums1Max[i+1] {
			nums1Max[i] = nums1[i]
		} else {
			nums1Max[i] = nums1Max[i+1]
		}
	}
	for i := len2 - 2; i >= 0; i-- { //统计nums2 在区间 [i, len(nums)-1] 区间的最大值
		nums2Map[nums2[i]] = append(nums2Map[nums2[i]], i) //统计nums2每个值出现的位置
		if nums2[i] > nums2Max[i+1] {
			nums2Max[i] = nums2[i]
		} else {
			nums2Max[i] = nums2Max[i+1]
		}
	}
	idx1, idx2 := 0, 0
	for i := 1; i <= k; i++ {
		
	}
	return ans
}
