package main

import "fmt"

/*
714. 买卖股票的最佳时机含手续费
	https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
题目描述
	给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
	你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
	返回获得利润的最大值。
	注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
示例 1:
	输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
	输出: 8
	解释: 能够达到的最大利润:
	在此处买入 prices[0] = 1
	在此处卖出 prices[3] = 8
	在此处买入 prices[4] = 4
	在此处卖出 prices[5] = 9
	总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
注意:
	0 < prices.length <= 50000.
	0 < prices[i] < 50000.
	0 <= fee < 50000.
*/
func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit2(prices, fee))
}

/*
动态规划
	dp[i] 第 i 天第情况
	dp[i][0] 第 i 天手里没有股票的情况
	dp[i][1] 第 i 天手里拥有一只股票的情况
	dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
	dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
*/
func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i:=1; i<len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	}
	return dp[len(prices)-1][0]
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

/*
优化版动态规划
	不需要保留所有 dp 状态，只需要保留上一个即可
*/
func maxProfit2(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp_0, dp_1 := 0, -prices[0]
	for i:=1; i<len(prices); i++ {
		dp_0, dp_1 = max(dp_0, dp_1+prices[i]-fee), max(dp_1, dp_0 - prices[i])
	}
	return dp_0
}