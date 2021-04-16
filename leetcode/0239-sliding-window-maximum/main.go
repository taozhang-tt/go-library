package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
	dequeue, ans := make([]int, 0), make([]int, 0)
	for idx := 0; idx < len(nums); idx++ {
		if len(dequeue) > 0 && idx-dequeue[0] >= k { //窗口左侧的元素已经划出
			dequeue = dequeue[1:]
		}
		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] < nums[idx] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, idx)
		if idx >= k-1 {
			ans = append(ans, nums[dequeue[0]])
		}
	}
	return ans
}
