package main

/*
15. 三数之和
	https://leetcode-cn.com/problems/3sum/
题目描述
	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
示例：
	给定数组 nums = [-1, 0, 1, 2, -1, -4]，
	满足要求的三元组集合为：
	[
	[-1, 0, 1],
	[-1, -1, 2]
	]
*/
import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSum2(nums))
	fmt.Println(threeSum3(nums))
	fmt.Println(threeSum4(nums))
}

/*
官方题解
*/
func threeSum4(nums []int) [][]int {
	ret, length := make([][]int, 0), len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < length; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			k := length - 1
			for j := i + 1; j < length; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					for j < k && nums[i]+nums[j]+nums[k] > 0 {
						k--
					}
					if j == k {
						break
					}
					if nums[i]+nums[j]+nums[k] == 0 {
						ret = append(ret, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return ret
}

/*
排序
	如果抛开题目中”不可以包含重复的三元组“的要求，直接枚举三元组（a, b, c）即可
	那「不重复」的本质是什么？我们保持三重循环的大框架不变，只需要保证：
		第二重循环枚举到的元素不小于当前第一重循环枚举到的元素；
		第三重循环枚举到的元素不小于当前第二重循环枚举到的元素。
	也就是说，对于我们枚举的三元组（a, b, c），如果 a <= b <= c，则只有（a, b, c）会被枚举到，除此之外的（a, c, b），（b, a, c）等都不会枚举到
	为了保证上述条件被满足，我们可以先对数组进行排序，再三重循环进行枚举
	同时，对于每一层循环来说，两次枚举的值不能相同，否则会出现元组相同的情况，比如枚举（-1, 1, 0, 0, 0），会出现多个 (-1, 1, 0)
	以上所述翻译成代码如下
*/
func threeSum(nums []int) [][]int {
	length, ret := len(nums), make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < length-2; first++ {
		if first != 0 && nums[first] == nums[first-1] { //避免同一层的枚举两次枚举的值相同
			continue
		}
		for second := first + 1; second < length-1; second++ {
			if second != first+1 && nums[second] == nums[second-1] {
				continue
			}
			for three := second + 1; three < length; three++ {
				if three != second+1 && nums[three] == nums[three-1] {
					continue
				}
				if nums[first]+nums[second]+nums[three] == 0 {
					ret = append(ret, []int{nums[first], nums[second], nums[three]})
				}
			}
		}
	}
	return ret
}

/*
排序 + 双指针
	接着上述思路，将第二层和第三层同时遍历，初始化两个指针l, r，l从左向右遍历，r从右向左遍历，思路如下
	因为数组已经排序，所以对于第一层遍历到到数字 a，第二层的遍历记为 b，第三层的记录为 c
	如果 a + b + c > 0，说明 b + c 太大了，那应该让 c 小一点（r指针向左移动）
	如果 a + b + c < 0，说明 b + c 太小了，那应该让 b 大一点（l指针向右移动）
	如果 a + b + c = 0，则找到一个解
*/
func threeSum2(nums []int) [][]int {
	length, ret := len(nums), make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < length-2; first++ {
		if first != 0 && nums[first] == nums[first-1] { //避免同一层的枚举两次枚举的值相同
			continue
		}
		l, r := first+1, len(nums)-1
		for l < r {
			if nums[first]+nums[l]+nums[r] == 0 {
				ret = append(ret, []int{nums[first], nums[l], nums[r]})
				l++
				r--
			} else if nums[first]+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}

func threeSum3(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1 //双指针
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			}
			if sum > 0 {
				r--
			}
			if sum == 0 {
				if r == len(nums)-1 || nums[r] != nums[r+1] {
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
				}
				l++
				r--
			}
		}
	}
	return ret
}
