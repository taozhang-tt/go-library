package main

import "fmt"

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}

func restoreString(s string, indices []int) string {
	ret := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ret[indices[i]] = s[i]
	}
	return string(ret)
}