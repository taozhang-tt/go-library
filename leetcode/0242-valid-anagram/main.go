package main

import "fmt"

/*
242. 有效的字母异位词
	https://leetcode-cn.com/problems/valid-anagram/
题目描述
	给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
示例 1:
	输入: s = "anagram", t = "nagaram"
	输出: true
示例 2:
	输入: s = "rat", t = "car"
	输出: false
*/

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram2(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sum := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		sum[s[i]]++
		sum[t[i]]--
	}
	for _, total := range sum {
		if total != 0 {
			return false
		}
	}
	return true
}

/*
最快的解法
*/
func isAnagram2(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	table := make([]int, 26)
	for i := 0; i < m; i++ {
		table[s[i]-'a']++
	}
	for i := 0; i < n; i++ {
		table[t[i]-'a']--
		if table[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
