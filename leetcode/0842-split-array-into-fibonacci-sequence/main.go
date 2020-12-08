package main

import "fmt"

/*
842. 将数组拆分成斐波那契序列
	https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/
题目描述
	给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。
	形式上，斐波那契式序列是一个非负整数列表 F，且满足：
	0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
	F.length >= 3；
	对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。
	另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。
	返回从 S 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []。
示例 1：
	输入："123456579"
	输出：[123,456,579]
示例 2：
	输入: "11235813"
	输出: [1,1,2,3,5,8,13]
示例 3：
	输入: "112358130"
	输出: []
	解释: 这项任务无法完成。
示例 4：
	输入："0123"
	输出：[]
	解释：每个块的数字不能以零开头，因此 "01"，"2"，"3" 不是有效答案。
示例 5：
	输入: "1101111"
	输出: [110, 1, 111]
	解释: 输出 [11,0,11,11] 也同样被接受。

提示：

	1 <= S.length <= 200
	字符串 S 中只含有数字。
*/
func main() {
	S := "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"
	fmt.Println(splitIntoFibonacci(S))

	S = "1230"
	fmt.Println(splitIntoFibonacci(S))

	S = "123456579"
	fmt.Println(splitIntoFibonacci(S))

	S = "11235813"
	fmt.Println(splitIntoFibonacci(S))

	S = "112358130"
	fmt.Println(splitIntoFibonacci(S))

	S = "0123"
	fmt.Println(splitIntoFibonacci(S))

	S = "123"
	fmt.Println(splitIntoFibonacci(S))
}

/*
回溯+剪枝
	回溯：拆分元素组成数字存入结果集，结果集元素数大于 3 时，判断最后三个元素是否满足斐波那契数列要求
	1. 如果第三个元素太小，则再从S中取一个字符，与第三个元素组成新的数字
	2. 如果第三个元素太大，则向前回溯，从 S 中取出一个元素和第二个元素组成新的数字当作新的第二个元素，再从S中取出一个元素作为结果集的第三个元素，继续判断是否满足
	3. 如果满足斐波那契要求，则直接从S中取出一个元素加入到结果集中，检查最后三个元素是否满足斐波那契数列要求
	剪枝：
		1. 元素组成的数字超过最大值
		2. 元素组成的数字以0开头
		3. 元素组成的数字大于结果集中它前面的两个元素
*/
func splitIntoFibonacci(S string) []int {
	curr := []int{}
	ans, ret := traceback(S, 0, curr)
	if ret {
		return ans
	}
	return []int{}
}

func traceback(S string, idx int, curr []int) (ans []int, ret bool) {
	length := len(curr)
	if length > 2 && curr[length-1] != curr[length-2]+curr[length-3] {	//剪枝 1
		return curr, false
	}
	if idx == len(S) && length > 2 {
		return curr, true
	}
	num := 0
	for i := idx; i < len(S); i++ {
		num = num*10 + int(S[i]-'0')
		if num > 2147483647 { //剪枝2 超过最大值，题目限制最大值为 2^32 -1
			return curr, false
		}
		curr = append(curr, num)
		if ans, ret = traceback(S, i+1, curr); ret {
			return
		} else {
			curr = curr[0 : len(curr)-1]	//回溯
			if num == 0 {	//剪枝3 如果数字为0，不能作为前缀，直接返回失败
				return curr, false
			}
		}
	}
	return
}
