package main

/*
327. 区间和的个数
	https://leetcode-cn.com/problems/count-of-range-sum/
题目描述
	给定一个整数数组 nums，返回区间和在 [lower, upper] 之间的个数，包含 lower 和 upper。
	区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
说明:
	最直观的算法复杂度是 O(n2) ，请在此基础上优化你的算法。
示例:
	输入: nums = [-2,5,-1], lower = -2, upper = 2,
	输出: 3
	解释: 3个区间分别是: [0,0], [2,2], [0,2]，它们表示的和分别为: -2, -1, 2。
*/
import (
	"fmt"
)

func main() {
	nums := []int{-2, 5, -1}
	lower := -2
	upper := 2
	fmt.Println(countRangeSum(nums, lower, upper))
	fmt.Println(countRangeSum2(nums, lower, upper))
}

/*
暴力解法，通过了
遍历每个元素 i，如果 [i, j]，（j 从 i 开始到最后一个元素止）的和满足条件，则结果 +1
*/
func countRangeSum2(nums []int, lower int, upper int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum <= upper && sum >= lower {
				cnt++
			}
		}
	}
	return cnt
}

/*
前缀和+归并排序
	前缀和记为 preSum
	区间和 lower <= S(i, j) <= upper 等价于 lower <= preSum[j] - preSum[i-1] <= upper
	假设有两个单调递增数组 n1, n2，寻找有多少个下标对 [i, j] 满足 lower <= n1[i] - n2[j] <= upper
*/
func countRangeSum(nums []int, lower, upper int) int {
    var mergeCount func([]int) int
    mergeCount = func(arr []int) int {
        n := len(arr)
        if n <= 1 {
            return 0
        }
        n1 := append([]int(nil), arr[:n/2]...)
        n2 := append([]int(nil), arr[n/2:]...)
        cnt := mergeCount(n1) + mergeCount(n2) // 递归完毕后，n1 和 n2 均为有序
        // 统计下标对的数量
        l, r := 0, 0
        for _, v := range n1 {
            for l < len(n2) && n2[l]-v < lower {
                l++
            }
            for r < len(n2) && n2[r]-v <= upper {
                r++
            }
            cnt += r - l //这里其实应该是 (r-1) - l + 1
        }
        // n1 和 n2 归并填入 arr
        p1, p2 := 0, 0
        for i := range arr {
            if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
                arr[i] = n1[p1]
                p1++
            } else {
                arr[i] = n2[p2]
                p2++
            }
        }
        return cnt
    }

    prefixSum := make([]int, len(nums)+1)
    for i, v := range nums {
        prefixSum[i+1] = prefixSum[i] + v
    }
    return mergeCount(prefixSum)
}