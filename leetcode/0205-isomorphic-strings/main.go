package main

import "fmt"

/*
205. 同构字符串
	https://leetcode-cn.com/problems/isomorphic-strings/
题目描述
	给定两个字符串 s 和 t，判断它们是否是同构的。
	如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
	所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
示例 1:
	输入: s = "egg", t = "add"
	输出: true
示例 2:
	输入: s = "foo", t = "bar"
	输出: false
示例 3:
	输入: s = "paper", t = "title"
	输出: true
	说明:
	你可以假设 s 和 t 具有相同的长度。
*/
func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))

	s = "foo"
	t = "tar"
	fmt.Println(isIsomorphic(s, t))

	s = "paper"
	t = "title"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	hashS, hashT := make(map[byte]byte), make(map[byte]byte)
	for i:=0; i<len(s); i++ {
		if value, ok := hashS[s[i]]; ok && value != t[i] {
			return false
		}
		if value, ok := hashT[t[i]]; ok && value != s[i] {
			return false
		}
		hashS[s[i]] = t[i]
		hashT[t[i]] = s[i]
	}
	return true
}