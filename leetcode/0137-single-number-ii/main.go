package main

import "fmt"

/*
137. 只出现一次的数字 II
	https://leetcode-cn.com/problems/single-number-ii/
题目描述
	给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
示例 1：
	输入：nums = [2,2,3,2]
	输出：3
示例 2：
	输入：nums = [0,1,0,1,0,1,99]
	输出：99
*/
func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))
}

// (所有元素(只计算一次)和 * 3 - 所有元素和)/2
func singleNumber(nums []int) int {
	t1, t2, m := 0, 0, make(map[int]bool)
	for _, num := range nums {
		t2 += num
		if _, exist := m[num]; !exist {
			t1 += num
			m[num] = true
		}
	}
	return (t1*3 - t2) / 2
}

// hash 直接统计
func singleNumber2(nums []int) int {
    freq := map[int]int{}
    for _, v := range nums {
        freq[v]++
    }
    for num, occ := range freq {
        if occ == 1 {
            return num
        }
    }
    return 0 // 不会发生，数据保证有一个元素仅出现一次
}

// 位运算
// 分别统计每个元素的二进制位第 i 位 1 的个数，出现 3 次的元素，个数除以 3 余数必为 0
func singleNumber3(nums []int) int {
    ans := int32(0)
    for i := 0; i < 32; i++ {
        total := int32(0)
        for _, num := range nums {
            total += int32(num) >> i & 1
        }
        if total%3 > 0 {
            ans |= 1 << i
        }
    }
    return int(ans)
}
