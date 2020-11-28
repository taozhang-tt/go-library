package main

/*
9. 回文数
*/

import "fmt"

func main() {
	x := 0
	fmt.Println(isPalindrome(x))

	x = 10
	fmt.Println(isPalindrome(x))

	x = -10
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	before, after := x, 0
	for x > 0 {
		after = after*10 + x%10
		x /= 10
	}
	return before == after
}
