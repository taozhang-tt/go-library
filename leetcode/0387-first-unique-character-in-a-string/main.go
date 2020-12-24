package main

import "fmt"

/*
387. 字符串中的第一个唯一字符
	https://leetcode-cn.com/problems/first-unique-character-in-a-string/
题目描述
	给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
示例：
	s = "leetcode"
	返回 0
	s = "loveleetcode"
	返回 2
*/
func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))

	s = "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

//暴力
func firstUniqChar(s string) int {
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
