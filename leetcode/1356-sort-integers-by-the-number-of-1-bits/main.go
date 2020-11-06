package main

/**
1356. 根据数字二进制下 1 的数目排序
	https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/
题目描述：
	给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。
	如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。
	请你返回排序后的数组。
示例 1：
	输入：arr = [0,1,2,3,4,5,6,7,8]
	输出：[0,1,2,4,8,3,5,6,7]
	解释：[0] 是唯一一个有 0 个 1 的数。
	[1,2,4,8] 都有 1 个 1 。
	[3,5,6] 有 2 个 1 。
	[7] 有 3 个 1 。
	按照 1 的个数排序得到的结果数组为 [0,1,2,4,8,3,5,6,7]
*/

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int
	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	sortByBits(arr)
	fmt.Println((arr))

	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	sortByBits2(arr)
	fmt.Println((arr))

	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	sortByBits3(arr)
	fmt.Println((arr))

	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	sortByBits4(arr)
	fmt.Println((arr))
}

/**
暴力解法
*/
func getOneNum (item int) (num int) {	//获取整数含有多少个二进制位的1
	for item > 0 {
		num += item%2
		item = item / 2
	}
	return num
}
func sortByBits(arr []int) []int {
	m := map[int]int{}
	insertSort := func (nums []int) {
		for i := 1; i < len(nums); i++ {
			temp := nums[i]
			var j = i - 1
			for j >= 0 && (m[nums[j]] > m[temp] || (m[nums[j]] == m[temp] && nums[j] > temp)) {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = temp
		}
	}
	for _, item := range arr {
		m[item] = getOneNum(item)
	}
	insertSort(arr)
	return arr
}

/**
同样是暴力法，使用库函数排序
*/
func sortByBits2(arr []int) []int {
	m := map[int]int{}
	for _, item := range arr {
		m[item] = getOneNum(item)
	}
	sort.Slice(arr, func(i, j int) bool {
        x, y := arr[i], arr[j]
        cx, cy := m[x], m[y]
        return cx < cy || cx == cy && x < y
	})
	return arr
}

/**
利用位运算获取 1 的个数
*/
func getOneByBit(num int) (total int) {
	for (num > 0) {
		total += num & 1
		num >>= 1
	}
	return total
}

func sortByBits3(arr []int) []int {
	m := map[int]int{}
	for _, item := range arr {
		m[item] = getOneByBit(item)
	}
	sort.Slice(arr, func(i, j int) bool {
        x, y := arr[i], arr[j]
        cx, cy := m[x], m[y]
        return cx < cy || cx == cy && x < y
	})
	return arr
}


/**
官方题解2：递推预处理
	我们定义 bit[i]，bit[i] 为数字 i 二进制表示下数字 1 的个数，则可以列出递推式：
	bit[i]=bit[i>>1]+(i&1)
	所以我们线性预处理 bitbit 数组然后去排序即可。

*/
var bit = [1e4 + 1]int{}

func init() {
    for i := 1; i <= 1e4; i++ {
        bit[i] = bit[i>>1] + i&1
    }
}

func sortByBits4(a []int) []int {
    sort.Slice(a, func(i, j int) bool {
        x, y := a[i], a[j]
        cx, cy := bit[x], bit[y]
        return cx < cy || cx == cy && x < y
    })
    return a
}


