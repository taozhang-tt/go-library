package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 3, 1}
	fmt.Println(reversePairs(nums))

	nums = []int{2, 4, 3, 5, 1}
	fmt.Println(reversePairs2(nums))
}

/*
暴力解法
超时
*/
func reversePairs(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j]*2 {
				cnt++
			}
		}
	}
	return cnt
}
/*
归并排序
	对于数组 [i, j]，取其中点 m，将其分为两个子数组 [i, m] 和 [m+1, j]，分别对两个子数组排序，同时求得子数组的翻转对数目分别为 cnt1，cnt2
	对于子数组 [i, m] 中的每一个元素，求得其相对于子数组 [m+1, j] 的翻转对数量 cnt[k]，k取值[i, m]
	则 cnt1 + cnt2 + cnt[k] 即为 [i, j] 数组对翻转对数量
*/
func reversePairs2(arr []int) int {
	length := len(arr)
	if length < 2 {
		return 0
	}
	mid := length / 2
	left, right := append([]int{}, arr[0:mid]...), append([]int{}, arr[mid:]...)
	cnt := reversePairs2(left) + reversePairs2(right)	//对两个子数组排序，获取子数组的翻转对数量
	p := 0
	for _, value := range left {	//求子数组中1的每个元素相对于子数组2的翻转对数量，此时两个子数组都是有序的
		for p < len(right) && value > 2*right[p] {
			p++
		}
		cnt += p
	}
	// left 和 right 归并填入 nums，也就是归并排序的归并操作
	p1, p2 := 0, 0
	for i := range arr {
		if p1 < len(left) && (p2 == len(right) || left[p1] <= right[p2]) {
			arr[i] = left[p1]
			p1++
		} else {
			arr[i] = right[p2]
			p2++
		}
	}
	return cnt
}
