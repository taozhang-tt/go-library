package main

import (
	"fmt"
	"strconv"
)

/*
738. 单调递增的数字
	https://leetcode-cn.com/problems/monotone-increasing-digits/
题目描述
	给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
	（当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
示例 1:
	输入: N = 10
	输出: 9
示例 2:
	输入: N = 1234
	输出: 1234
示例 3:
	输入: N = 332
	输出: 299
*/
func main() {
	N := 999998
	fmt.Println(monotoneIncreasingDigits(N))


	N = 1234
	fmt.Println(monotoneIncreasingDigits(N))

	N = 100
	fmt.Println(monotoneIncreasingDigits(N))


	N = 4321
	fmt.Println(monotoneIncreasingDigits2(N))
}

func monotoneIncreasingDigits(N int) int {
	nums, base := N, 1
	for nums > 9 {
		base *= 10
		nums /= 10
	}
	if ret, ok := backtrace(N, 0, base, false) ; ok {
		return ret
	}
	return 0
}

/*
回溯 + 剪枝
	从左向右，取第一位，余下的数字向后遍历
*/

func backtrace(num, min, base int, sign bool) (int, bool) {
	curr := num / base
	max := curr
	if sign {
		max = 9
	}
	if base == 1 && max >= min {
		return max, true
	}
	for i:=max; i>=min; i-- {
		if i != max {
			sign = true
		}
		if ret, ok := backtrace(num%base, i, base/10, sign); ok {
			return i * base + ret, true
		}
	}
	return 0, false
}

/*
尽可能让靠左的位置的数字与原数字相同
先找出递增子串
该子串从末尾向前，每一位减1，直到找到一位 i 满足，s[i]--以后还能满足单调递增，然后 [i:] 全部置为 9
*/
func monotoneIncreasingDigits2(n int) int {
    s := []byte(strconv.Itoa(n))
    i := 1
    for i < len(s) && s[i] >= s[i-1] {
        i++
    }
    if i < len(s) {
        for i > 0 && s[i] < s[i-1] {
            s[i-1]--
            i--
        }
        for i++; i < len(s); i++ {
            s[i] = '9'
        }
    }
    ans, _ := strconv.Atoi(string(s))
    return ans
}
