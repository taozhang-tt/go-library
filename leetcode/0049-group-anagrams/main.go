package main

import (
	"fmt"
	"sort"
)

/*
49. 字母异位词分组
	https://leetcode-cn.com/problems/group-anagrams/
题目描述
	给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
示例:
	输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
	输出:
	[
	["ate","eat","tea"],
	["nat","tan"],
	["bat"]
	]
说明：
	所有输入均为小写字母。
	不考虑答案输出的顺序。
*/
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams2(strs))
}

/*
排序+hash
	先对每个字符串排序，用排序好的字符串作为 map 的 key 进行统计
*/
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		sortedStr := string(bs)
		hash[sortedStr] = append(hash[sortedStr], str)
	}
	ans := make([][]string, 0, len(hash))
	for _, value := range hash {
		ans = append(ans, value)
	}
	return ans
}

/*
hash
	对每个单词开辟一个长度为 26 的数组统计每个字符出现的次数，用该数组作为 map 的key
*/
func groupAnagrams2(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for _, ch := range str {
			key[ch-'a'] ++
		}
		hash[key] = append(hash[key], str)
	}
	ans := make([][]string, 0)
	for _, value := range hash {
		ans = append(ans, value)
	}
	return ans
}