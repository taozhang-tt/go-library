package main

import "fmt"

/*
389. 找不同(简单)
	https://leetcode-cn.com/problems/find-the-difference/
题目描述
	给定两个字符串 s 和 t，它们只包含小写字母。
	字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
	请找出在 t 中被添加的字母。
示例 1：
	输入：s = "abcd", t = "abcde"
	输出："e"
	解释：'e' 是那个被添加的字母。
示例 2：
	输入：s = "", t = "y"
	输出："y"
示例 3：
	输入：s = "a", t = "aa"
	输出："a"
示例 4：
	输入：s = "ae", t = "aea"
	输出："a"
提示：
	0 <= s.length <= 1000
	t.length == s.length + 1
	s 和 t 只包含小写字母
*/
func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(findTheDifference(s, t))
	fmt.Println(findTheDifference2(s, t))
	fmt.Println(findTheDifference3(s, t))
}

/*
计数
*/
func findTheDifference(s string, t string) byte {
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		cnt[t[i]-'a']--
		if cnt[t[i]-'a'] < 0 {
			return t[i]
		}
	}
	return t[0]
}

/*
求和
*/
func findTheDifference2(s string, t string) byte {
	cntS, cntT := 0, 0
	for i := 0; i < len(s); i++ {
		cntS += int(s[i])
		cntT += int(t[i])
	}
	cntT += int(t[len(t)-1])
	return byte(cntT - cntS)
}

/*
位运算
	异或运算规律
		a ^ 0 = a
		a ^ a = 0
		a ^ b ^ a = a ^ a ^ b = (a ^ a) ^ b = 0 ^ b = b
*/
func findTheDifference3(s string, t string) (diff byte) {
	for i:= range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}