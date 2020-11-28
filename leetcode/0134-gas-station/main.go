package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))

	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println(canCompleteCircuit2(gas, cost))

	gas = []int{1, 2, 3}
	cost = []int{3, 1, 2}
	fmt.Println(canCompleteCircuit3(gas, cost))
}

/*
思路是对的，复杂度也是 n 的平方，但是比官方题解慢很多啊
*/
func canCompleteCircuit(gas []int, cost []int) int {
	temp := []int{}
	long := len(gas)
	for i := 0; i < long; i++ {
		temp = append(temp, gas[i]-cost[i])
	}
	circleJudge := func(arr []int, begin int) bool {
		goal := arr[begin]
		for curr := (begin + 1) % len(arr); curr != begin; {
			goal += arr[curr]
			if goal < 0 {
				return false
			}
			curr = (curr + 1) % len(arr)
		}
		return true
	}
	for i := 0; i < long; i++ {
		if temp[i] < 0 { //为负数的不可能作为起点，因为无法到达下一个加油站
			continue
		}
		if circleJudge(temp, i) { //判断以 i 为起点，能否走一圈
			return i
		}
	}
	return -1
}

/*
再写一遍，暴力暴力，和官方很类似，但是依然慢，因为没有加速
*/
func canCompleteCircuit2(gas []int, cost []int) int {
	for i, long := 0, len(gas); i < long; i++ {
		if gas[i] < cost[i] {
			continue
		}
		sumGas, costGas, cnt := 0, 0, 0
		for cnt < long {
			j := (i + cnt) % long
			sumGas += gas[j]
			costGas += cost[j]
			if sumGas < costGas {
				break
			}
			cnt++
		}
		if cnt == long {
			return i
		}
	}
	return -1
}

/*
官方题解，加速版
*/
func canCompleteCircuit3(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
