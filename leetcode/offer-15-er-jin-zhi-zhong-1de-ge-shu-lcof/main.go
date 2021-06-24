package main

import "fmt"

func main() {
	num := 00000000000000000000000000001011
	fmt.Println(hammingWeight(uint32(num)))
}

// 循环检查二进制位
func hammingWeight(num uint32) int {
	ret := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ret++
		}
	}
	return ret
}
