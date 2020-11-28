package main

import (
	"fmt"
	"sort"
)

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}

/*
官方题解：从高到底考虑
	将 people 按其高度从 高->低 排序，高度想同时按第二个维度从 小->大 排序，第i个人高度记为 hi，ki为第二个维度
	对于排序后，假设 0 ~ i 个人已经在最后的队列中找到了恰当的位置，那么对于第 i+1 个元素，将其放在任何位置，都不会影响前 i 个人位置第正确性
	所以可以采用插入法
*/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	ans := [][]int{}
	for _, item := range people {
		idx := item[1]                                                  //前面有几个比我大的元素？
		ans = append(ans[:idx], append([][]int{item}, ans[idx:]...)...) //比我大多少我就在前面空出多少个位置
	}
	return ans
}
