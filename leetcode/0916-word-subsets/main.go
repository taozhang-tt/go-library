package main

import "fmt"

/*
916. 单词子集
	https://leetcode-cn.com/problems/word-subsets/
题目描述
	我们给出两个单词数组 A 和 B。每个单词都是一串小写字母。
	现在，如果 b 中的每个字母都出现在 a 中，包括重复出现的字母，那么称单词 b 是单词 a 的子集。 例如，“wrr” 是 “warrior” 的子集，但不是 “world” 的子集。
	如果对 B 中的每一个单词 b，b 都是 a 的子集，那么我们称 A 中的单词 a 是通用的。
	你可以按任意顺序以列表形式返回 A 中所有的通用单词。
示例 1：
	输入：A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
	输出：["facebook","google","leetcode"]
示例 2：
	输入：A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
	输出：["apple","google","leetcode"]
示例 3：
	输入：A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
	输出：["facebook","google"]
示例 4：
	输入：A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
	输出：["google","leetcode"]
示例 5：
	输入：A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
	输出：["facebook","leetcode"]
提示：
	1 <= A.length, B.length <= 10000
	1 <= A[i].length, B[i].length <= 10
	A[i] 和 B[i] 只由小写字母组成。
	A[i] 中所有的单词都是独一无二的，也就是说不存在 i != j 使得 A[i] == A[j]。
*/
func main() {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"e", "o"}
	fmt.Println(wordSubsets2(A, B))

	A = []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B = []string{"l", "e"}
	fmt.Println(wordSubsets2(A, B))

	A = []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B = []string{"e", "oo"}
	fmt.Println(wordSubsets2(A, B))

	A = []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B = []string{"lo", "eo"}
	fmt.Println(wordSubsets2(A, B))

	A = []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B = []string{"ec", "oc", "ceo"}
	fmt.Println(wordSubsets(A, B))
}

/*
暴力法（通过）
*/
func wordSubsets(A []string, B []string) []string {
	hashB, ans := make(map[byte]int), make([]string, 0) //先统计B的要求，有哪些字符需要出现，每个字符至少出现几次
	for _, word := range B {
		hashWord := make(map[byte]int)
		for i := 0; i < len(word); i++ {
			b := word[i]
			hashWord[b]++
			if hashWord[b] > hashB[b] {
				hashB[b] = hashWord[b]
			}
		}
	}
	for _, word := range A {
		hashWord, sign := make(map[byte]int), true
		for i := 0; i < len(word); i++ {
			b := word[i]
			hashWord[b]++
		}
		for b, cnt := range hashB {
			if cnt > hashWord[b] {
				sign = false
				break
			}
		}
		if sign {
			ans = append(ans, word)
		}
	}
	return ans
}

/*
用 slice 计数，比 map 计数要快
*/
func wordSubsets2(A []string, B []string) []string {
	aim, ans := [26]int{}, []string{}
	for _, word := range B {
		tmp := [26]int{}
		for i := range word {
			b := word[i] - 'a'
			tmp[b]++
			if tmp[b] > aim[b] {
				aim[b] = tmp[b]
			}
		}
	}
	for _, word := range A {
		tmp := [26]int{}
		for i := range word {
			b := word[i] - 'a'
			tmp[b]++
		}
		flag := true
		for i, v := range tmp {
			if v < aim[i] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, word)
		}
	}
	return ans
}
