package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
1046. 最后一块石头的重量
	https://leetcode-cn.com/problems/last-stone-weight/
题目描述
	有一堆石头，每块石头的重量都是正整数。
	每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
	如果 x == y，那么两块石头都会被完全粉碎；
	如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
	最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
示例：
	输入：[2,7,4,1,8,1]
	输出：1
	解释：
	先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
	再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
	接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
	最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
提示：
	1 <= stones.length <= 30
	1 <= stones[i] <= 1000
*/
func main() {
	stones := []int{3, 7, 8}
	fmt.Println(lastStoneWeight(stones))

	stones = []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))

	stones = []int{2, 7}
	fmt.Println(lastStoneWeight2(stones))

	stones = []int{2}
	fmt.Println(lastStoneWeight2(stones))
}

/*
将 stones 降序排列，而后模拟题目描述的操作
*/
func lastStoneWeight(stones []int) int {
	sort.Slice(stones, func(i, j int) bool { return stones[i] > stones[j] })
	for len(stones) > 2 {
		if stones[0] == stones[1] { //如果最大的两个石头重量相等，直接全部粉碎
			stones = stones[2:]
			continue
		}
		temp := stones[0] - stones[1]      //两个石头重量不等，都要粉碎，用temp记录粉碎后产生的小石头的重量
		stones = stones[1:]                //粉碎较大的那个石头，第二大的石头在放入新石头时顺便粉碎
		for i := 0; i < len(stones); i++ { //将新产生的小石头放到合适的位置，保证所有石头的重量仍是降序排列
			if i == len(stones)-1 || stones[i+1] <= temp {
				stones[i] = temp
				break
			} else {
				stones[i] = stones[i+1]
			}
		}
	}
	if len(stones) == 2 {
		return stones[0] - stones[1]
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

/*
最大堆
*/
func lastStoneWeight2(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}
