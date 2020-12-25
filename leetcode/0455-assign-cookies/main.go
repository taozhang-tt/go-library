package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))

	g = []int{1, 2}
	s = []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))

	g = []int{}
	s = []int{1}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	index := 0
	for i := 0; i < len(s); i++ {
		if index == len(g) {
			break
		}
		if s[i] >= g[index] {
			ans++
			index++
		}
	}
	return
}
