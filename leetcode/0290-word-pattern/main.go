package main

import (
	"fmt"
	"strings"
)

/*
290. 单词规律
	https://leetcode-cn.com/problems/word-pattern/
题目描述
	给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
	这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
示例1:
	输入: pattern = "abba", str = "dog cat cat dog"
	输出: true
示例 2:
	输入:pattern = "abba", str = "dog cat cat fish"
	输出: false
示例 3:
	输入: pattern = "aaaa", str = "dog cat cat dog"
	输出: false
示例 4:
	输入: pattern = "abba", str = "dog dog dog dog"
	输出: false
说明:
	你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
*/
func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))

	pattern = "abba"
	str = "dog cat cat fish"
	fmt.Println(wordPattern(pattern, str))

	pattern = "aaaa"
	str = "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))

	pattern = "abba"
	str = "dog dog dog dog"
	fmt.Println(wordPattern(pattern, str))

	pattern = "aaa"
	str = "aa aa aa aa"
	fmt.Println(wordPattern(pattern, str))
}

func wordPattern(pattern string, s string) bool {
	strs, hashPattern, hashS := strings.Split(s, " "), make(map[byte]string), make(map[string]byte)
	if len(strs) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		str := strs[i]
		if value, ok := hashPattern[pattern[i]]; ok && str != value {
			return false
		}
		if value, ok := hashS[str]; ok && pattern[i] != value {
			return false
		}
		hashPattern[pattern[i]] = str
		hashS[str] = pattern[i]
	}
	return true
}
