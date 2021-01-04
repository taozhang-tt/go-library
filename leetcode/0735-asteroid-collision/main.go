package main

import "fmt"

func main() {
	asteroids := []int{1, -1, -2, -2}
	fmt.Println(asteroidCollision(asteroids))

	asteroids = []int{5, 10, -5}
	fmt.Println(asteroidCollision(asteroids))

	asteroids = []int{8, -8}
	fmt.Println(asteroidCollision(asteroids))

	asteroids = []int{-2, -1, 1, 2}
	fmt.Println(asteroidCollision(asteroids))
}	

/*
暴力法，直接模拟
*/
func asteroidCollision(asteroids []int) (ans []int) {
	if len(asteroids) == 0 {
		return
	}
	for _, item := range asteroids {
		if len(ans) == 0 || (ans[len(ans)-1] < 0 && item > 0) || (ans[len(ans)-1] * item > 0){	//不会碰撞
			ans = append(ans, item)
			continue
		}
		for j:=len(ans)-1; j>=0; j-- {	//开始碰撞
			if (ans[j] < 0 && item > 0) || (ans[j] * item > 0){	//不会碰撞
				ans = append(ans, item)
				break
			}
			absJ, absA := abs(ans[j]), abs(item)
			if absJ == absA {	//一样大小，都爆炸
				ans = ans[:len(ans)-1]
				break
			} 
			if absA < absJ {	//没有前一个大，当前的爆炸
				break
			}
			if absA > absJ {	//比前一个大，前一个爆炸
				ans = ans[:len(ans)-1]
				if len(ans) == 0 {
					ans = append(ans, item)
					break
				}
			}
		}
	}
	return
}

func abs(value int) int {
	if value > 0 {
		return value
	}
	return -value
}