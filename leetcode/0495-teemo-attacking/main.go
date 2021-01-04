package main

import "fmt"
/*
495. 提莫攻击
	https://leetcode-cn.com/problems/teemo-attacking/
题目描述
	在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄，他的攻击可以让敌方英雄艾希（编者注：寒冰射手）进入中毒状态。现在，给出提莫对艾希的攻击时间序列和提莫攻击的中毒持续时间，你需要输出艾希的中毒状态总时长。
	你可以认为提莫在给定的时间点进行攻击，并立即使艾希处于中毒状态。
示例1:
	输入: [1,4], 2
	输出: 4
	原因: 第 1 秒初，提莫开始对艾希进行攻击并使其立即中毒。中毒状态会维持 2 秒钟，直到第 2 秒末结束。
	第 4 秒初，提莫再次攻击艾希，使得艾希获得另外 2 秒中毒时间。
	所以最终输出 4 秒。
示例2:
	输入: [1,2], 2
	输出: 3
	原因: 第 1 秒初，提莫开始对艾希进行攻击并使其立即中毒。中毒状态会维持 2 秒钟，直到第 2 秒末结束。
	但是第 2 秒初，提莫再次攻击了已经处于中毒状态的艾希。
	由于中毒状态不可叠加，提莫在第 2 秒初的这次攻击会在第 3 秒末结束。
	所以最终输出 3 。
提示：
	你可以假定时间序列数组的总长度不超过 10000。
	你可以假定提莫攻击时间序列中的数字和提莫攻击的中毒持续时间都是非负整数，并且不超过 10,000,000。
*/
func main() {
	timeSeries, duration := []int{1, 4}, 2
	fmt.Println(findPoisonedDuration(timeSeries, duration))

	timeSeries, duration = []int{1, 2}, 2
	fmt.Println(findPoisonedDuration(timeSeries, duration))

	timeSeries, duration = []int{1, 2}, 3
	fmt.Println(findPoisonedDuration2(timeSeries, duration))
}

/*
区间合并
*/
func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
	if len(timeSeries) == 0 {
		return 0
	}
	l, r := timeSeries[0], timeSeries[0]+duration-1
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] > r {
			ans += r - l + 1
			l = timeSeries[i]
		}
		r = timeSeries[i] + duration - 1
	}
	return ans + r - l + 1
}

/*
优化版
	中毒时间不能叠加，考虑相邻两个攻击点 t[i], t[i+1] 和持续时间 t 的关系
	1. 如果 t[i] + t <= t[i+1]，说明在下次攻击到来之前，中毒状态已经结束，中毒持续时间为 t
	2. 如果 t[i] + t > t[i+1]，说明下次攻击到来的时候，上次攻击的中毒状态还未结束，两次攻击之间中毒时间为 t[i+1] - t[i]
*/
func findPoisonedDuration2(timeSeries []int, duration int) (ans int) {
	if len(timeSeries) == 0 {
		return 0
	}
	for i := 0; i < len(timeSeries)-1; i++ {
		if (timeSeries[i+1] - timeSeries[i]) < duration {
			ans += timeSeries[i+1] - timeSeries[i]
		} else {
			ans += duration
		}
	}
	return ans + duration
}
