package main
/*
151. 翻转字符串里的单词（简单）
	https://leetcode-cn.com/problems/reverse-words-in-a-string/
题目描述
	给定一个字符串，逐个翻转字符串中的每个单词。
说明：
	无空格字符构成一个 单词 。
	输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
	如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
示例 1：
	输入："the sky is blue"
	输出："blue is sky the"
示例 2：
	输入："  hello world!  "
	输出："world! hello"
	解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：
	输入："a good   example"
	输出："example good a"
	解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
示例 4：
	输入：s = "  Bob    Loves  Alice   "
	输出："Alice Loves Bob"
示例 5：
	输入：s = "Alice does not even like bob"
	输出："bob like even not does Alice"
提示：
	1 <= s.length <= 104
	s 包含英文大小写字母、数字和空格 ' '
	s 中 至少存在一个 单词
*/
import (
	"fmt"
)

func main() {
	s := " "
	fmt.Println(reverseWords(s))

	s = "  hello world!  "
	fmt.Println(reverseWords(s))

	s = "  "
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	word, ans := []byte{}, []byte{}
	for i:=0; i<len(s); i++ {
		b := s[i]
		if b == ' ' {
			if len(word) > 0 {
				word = append(word, b)
				ans = append(word, ans...)
				word = []byte{}
			}
		} else {
			word = append(word, b)
		}
	}
	if len(word) > 0 {
		word = append(word, ' ')
		ans = append(word, ans...)
	}
	if len(ans) == 0 {
		return ""
	}
	return string(ans[:len(ans)-1])
}