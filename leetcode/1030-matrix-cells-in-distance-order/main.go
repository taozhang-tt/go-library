package main

import (
	"fmt"
	"math"
	"sort"
)
/**
1030. 距离顺序排列矩阵单元格
	https://leetcode-cn.com/problems/matrix-cells-in-distance-order/
题目描述：
	给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。
	另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。
	返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）
示例 1：
	输入：R = 1, C = 2, r0 = 0, c0 = 0
	输出：[[0,0],[0,1]]
	解释：从 (r0, c0) 到其他单元格的距离为：[0,1]
示例 2：
	输入：R = 2, C = 2, r0 = 0, c0 = 1
	输出：[[0,1],[0,0],[1,1],[1,0]]
	解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2]
	[[0,1],[1,1],[0,0],[1,0]] 也会被视作正确答案。
*/
func main() {
	R := 1
	C := 2
	r0 := 0
	c0 := 0
	fmt.Println(allCellsDistOrder(R, C, r0, c0))

	R = 2
	C = 2
	r0 = 0
	c0 = 1
	fmt.Println(allCellsDistOrder2(R, C, r0, c0))
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	points := [][]int{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			points = append(points, []int{i, j})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return math.Abs(float64(points[i][0]-r0))+math.Abs(float64(points[i][1]-c0)) < math.Abs(float64(points[j][0]-r0))+math.Abs(float64(points[j][1]-c0))
	})
	return points
}

func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	queue := [][]int{{r0, c0}}
	visited := map[int]bool{r0*C + c0: true}
	ret := [][]int{}
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //上、右、下、左
	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]
		ret = append(ret, []int{r, c})
		for i := 0; i < 4; i++ {
			rr := r + dir[i][0]
			cc := c + dir[i][1]
			if rr >= 0 && rr < R && cc >= 0 && cc < C && !visited[rr*C+cc] {
				queue = append(queue, []int{rr, cc})
				visited[rr*C+cc] = true
			}
		}
	}
	return ret
}
