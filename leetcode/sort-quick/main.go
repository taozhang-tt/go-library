package main

import "fmt"

func main() {
	var nums []int

	nums = []int{9, 4, 3, 5, 6, 1, 2}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)

	nums = []int{9, 4, 3, 5, 6, 1, 2}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

//1，选取一个基准（数组第一个元素 nums[l]）记为 mid
//2，从数组的右侧开始向前查找，找到一个比基准元素小的元素 nums[r]，因为基准元素已被保存到 mid 中，所以原来的位置 nums[l] 空出来一个坑，将 nums[r] 保存到 nums[l]
//3，从数组的左侧开始向后查找，找到一个比基准元素大的元素 nums[l]，因为步骤 2 中 nums[r] 已被保存到 nums[l] 中，所以又空出来一个坑，将新查找到的 nums[l] 保存到 nuns[r] 中
//4，重复步骤 2，3 直到 l == r，将 mid 元素 保存到 nums[l] 中，数组的坑被填满
//5，步骤 4 结束后，数组已被 mid 元素分为两部分，mid 前面的元素全部小于 mid，后面的元素全部大于mid，此时递归调用函数对两部分元素分别进行排序
func QuickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	l := left
	r := right
	mid := nums[l]
	for l < r {
		for l < r && nums[r] > mid {
			r--
		}
		if l < r {
			nums[l] = nums[r]
			l++
		}
		for l < r && nums[l] < mid {
			l++
		}
		if l < r {
			nums[r] = nums[l]
			r--
		}
	}
	nums[l] = mid
	QuickSort(nums, left, l-1)
	QuickSort(nums, l+1, right)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	l, r, mid := left, right, arr[left]
	for l < r {
		for l < r && arr[r] > mid {
			r--
		}
		if (l < r) {
			arr[l] = arr[r]
			l++
		}
		for l < r && arr[l] < mid {
			l++
		}
		if (l < r) {
			arr[r] = arr[l]
			r--
		}
	}
	arr[l] = mid
	quickSort(arr, left, l-1)
	quickSort(arr , l+1, right)
}