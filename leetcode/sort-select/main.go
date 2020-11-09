package main

import "fmt"

func main() {
	var nums []int
	nums = []int{9, 4, 3, 5, 6, 1, 2}
	SelectSort(nums)
	fmt.Println(nums)

	nums = []int{9, 4, 3, 5, 6, 1, 2}
	selectSort(nums)
	fmt.Println(nums)
}

//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
//再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func selectSort(arr []int) {
	for i:=0; i<len(arr); i++ {
		minIndex := i
		for j:=i+1; j<len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}


