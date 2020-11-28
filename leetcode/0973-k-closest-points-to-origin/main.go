package main

/*
973. 最接近原点的 K 个点
	https://leetcode-cn.com/problems/k-closest-points-to-origin/
题目描述：
	我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。（这里，平面上两点之间的距离是欧几里德距离。）
	你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。
示例 1：
	输入：points = [[1,3],[-2,2]], K = 1
	输出：[[-2,2]]
	解释：
	(1, 3) 和原点之间的距离为 sqrt(10)，
	(-2, 2) 和原点之间的距离为 sqrt(8)，
	由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
	我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
示例 2：
	输入：points = [[3,3],[5,-1],[-2,4]], K = 2
	输出：[[3,3],[-2,4]]
	（答案 [[-2,4],[3,3]] 也会被接受。）
*/

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var points = [][]int{}
	var K int
	points = [][]int{{3, 3}, {5, -1}, {-2, 4}}
	K = 1
	fmt.Println(kClosest2(points, K))

	points = [][]int{{68, 97}, {34, -84}, {60, 100}, {2, 31}, {-27, -38}, {-73, -74}, {-55, -39}, {62, 91}, {62, 92}, {-57, -67}}
	K = 5
	fmt.Println(kClosest2(points, K))
}

type Point struct {
	X int
	Y int
}

type DisHeap []Point

func (h DisHeap) Len() int {
	return len(h)
}

func (h DisHeap) Less(i, j int) bool {
	return (h[i].X*h[i].X + h[i].Y*h[i].Y) < (h[j].X*h[j].X + h[j].Y*h[j].Y)
}

func (h DisHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DisHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *DisHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
最小堆做法
*/
func kClosest(points [][]int, K int) [][]int {
	disHeap := &DisHeap{}
	ret := [][]int{}
	heap.Init(disHeap)
	for _, item := range points {
		heap.Push(disHeap, Point{
			X: item[0],
			Y: item[1],
		})
	}
	for i := 0; i < K; i++ {
		item := heap.Pop(disHeap)
		ret = append(ret, []int{item.(Point).X, item.(Point).Y})
	}
	return ret
}

/*
官方题解、直接排序
*/
func kClosest2(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:K]
}
