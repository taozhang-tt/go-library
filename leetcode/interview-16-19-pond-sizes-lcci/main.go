package main

import (
	"fmt"
	"sort"
)

/*
面试题 16.19. 水域大小
	https://leetcode-cn.com/problems/pond-sizes-lcci/
题目描述
	你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
示例：
	输入：
	[
	[0,2,1,0],
	[0,1,0,1],
	[1,1,0,1],
	[0,1,0,1]
	]
	输出： [1,2,4]
*/
func main() {
	land := [][]int{

		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}
	fmt.Println(pondSizes(land))
}

var direction = [][]int{ //上、右上、右、右下、下、左下、左、左上
	{-1, 0},  //上
	{-1, 1},  //右上
	{0, 1},   //右
	{1, 1},   //右下
	{1, 0},   //下
	{1, -1},  //左下
	{0, -1},  //左
	{-1, -1}, //左上
}

func pondSizes(land [][]int) []int {
	ans := make([]int, 0)
	rows := len(land)
	if rows == 0 {
		return ans
	}
	cols := len(land[0])
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if land[x][y] == 0 {
				if cnt := depth(land, x, y); cnt > 0 {
					ans = append(ans, cnt)
				}
				fmt.Println(land)
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func depth(land [][]int, x, y int) int {
	cnt := 0
	if land[x][y] == 0 {
		cnt++
		land[x][y] = -1
		for i := 0; i < 8; i++ {
			xx, yy := x+direction[i][0], y+direction[i][1]
			if xx >= 0 && xx < len(land) && yy >= 0 && yy < len(land[0]) {
				cnt += depth(land, xx, yy)
			}
		}
	} else {
		land[x][y] = -1
	}
	return cnt
}
