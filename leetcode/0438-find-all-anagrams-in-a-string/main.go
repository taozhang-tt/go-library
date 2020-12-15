package main

import (
	"fmt"
)

/*
438. 找到字符串中所有字母异位词
	https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
题目描述
	给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
	字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
	说明：
	字母异位词指字母相同，但排列不同的字符串。
	不考虑答案输出的顺序。
示例 1:
	输入:
	s: "cbaebabacd" p: "abc"
	输出:
	[0, 6]
	解释:
	起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
	起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 示例 2:
	输入:
	s: "abab" p: "ab"
	输出:
	[0, 1, 2]
	解释:
	起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
	起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
	起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/
func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))

	s = "baa"
	p = "aa"
	fmt.Println(findAnagrams(s, p))
}

/*
双指针 + 计数的方法
*/
func findAnagrams(s string, p string) []int {
	aimP, ans := [26]int{}, []int{}	//aimP 记录 p 每个字符出现的次数
	for _, item := range p {
		aimP[item-'a']++
	}
	first, last, currS := 0, 0, [26]int{}	//双指针遍历数组，currS 记录 从 first到last 这个区间内每个字符出现的次数
	for last < len(s) {
		if aimP[s[last]-'a'] == 0 {	//该字符不存在 p 中
			first, last, currS = last+1, last+1, [26]int{}
			continue
		}
		currS[s[last]-'a']++
		if last - first + 1 == len(p) {
			if currS == aimP {
				ans = append(ans, first)
			}
			currS[s[first]-'a']--
			first++
		}
		last++
	}
	return ans
}
