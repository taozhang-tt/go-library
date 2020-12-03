package main

import "fmt"

/*
204. 计数质数
	https://leetcode-cn.com/problems/count-primes/
题目描述
	统计所有小于非负整数 n 的质数的数量。
示例 1：
	输入：n = 10
	输出：4
	解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
示例 2：
	输入：n = 0
	输出：0
示例 3：
	输入：n = 1
	输出：0
*/
func main() {
	n:= 2
	fmt.Println(countPrimes(n))

	n = 5
	fmt.Println(countPrimes(n))

	n = 1
	fmt.Println(countPrimes(n))

	n = 10
	fmt.Println(countPrimes(n))

	n = 11
	fmt.Println(countPrimes(n))

	n = 13
	fmt.Println(countPrimes2(n))
}

/*
枚举法，超时
*/
func isPrime(x int) bool {
    for i := 2; i*i <= x; i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}

func countPrimes(n int) (cnt int) {
    for i := 2; i < n; i++ {
        if isPrime(i) {
            cnt++
        }
    }
    return
}

/*
埃氏筛
	如果 x 是质数，则 2x, 3x, 4x...肯定是合数
*/
func countPrimes2(n int) (cnt int) {
	isPrime := make([]bool, n)
    for i := range isPrime {
        isPrime[i] = true
    }
	for i := 2; i < n; i++ {
        if isPrime[i] {
            cnt++
            for j := 2 * i; j < n; j += i {
                isPrime[j] = false
            }
        }
    }
    return
}