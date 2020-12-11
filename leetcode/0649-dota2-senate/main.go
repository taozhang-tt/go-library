package main

import "fmt"

/*
649. Dota2 参议院（中等）
	https://leetcode-cn.com/problems/dota2-senate/
题目描述
	Dota2 的世界里有两个阵营：Radiant(天辉)和 Dire(夜魇)
	Dota2 参议院由来自两派的参议员组成。现在参议院希望对一个 Dota2 游戏里的改变作出决定。他们以一个基于轮为过程的投票进行。在每一轮中，每一位参议员都可以行使两项权利中的一项：
	禁止一名参议员的权利：
	参议员可以让另一位参议员在这一轮和随后的几轮中丧失所有的权利。
	宣布胜利：
	   如果参议员发现有权利投票的参议员都是同一个阵营的，他可以宣布胜利并决定在游戏中的有关变化。
	给定一个字符串代表每个参议员的阵营。字母 “R” 和 “D” 分别代表了 Radiant（天辉）和 Dire（夜魇）。然后，如果有 n 个参议员，给定字符串的大小将是 n。
	以轮为基础的过程从给定顺序的第一个参议员开始到最后一个参议员结束。这一过程将持续到投票结束。所有失去权利的参议员将在过程中被跳过。
	假设每一位参议员都足够聪明，会为自己的政党做出最好的策略，你需要预测哪一方最终会宣布胜利并在 Dota2 游戏中决定改变。输出应该是 Radiant 或 Dire。
示例 1：
	输入："RD"
	输出："Radiant"
	解释：第一个参议员来自 Radiant 阵营并且他可以使用第一项权利让第二个参议员失去权力，因此第二个参议员将被跳过因为他没有任何权利。然后在第二轮的时候，第一个参议员可以宣布胜利，因为他是唯一一个有投票权的人
示例 2：
	输入："RDD"
	输出："Dire"
	解释：
		第一轮中,第一个来自 Radiant 阵营的参议员可以使用第一项权利禁止第二个参议员的权利
		第二个来自 Dire 阵营的参议员会被跳过因为他的权利被禁止
		第三个来自 Dire 阵营的参议员可以使用他的第一项权利禁止第一个参议员的权利
		因此在第二轮只剩下第三个参议员拥有投票的权利,于是他可以宣布胜利
*/
func main() {
	senate := "RD"
	fmt.Println(predictPartyVictory(senate))

	senate = "RDD"
	fmt.Println(predictPartyVictory2(senate))

	senate = "R"
	fmt.Println(predictPartyVictory2(senate))
}

/*
暴力法，超时
*/
func predictPartyVictory(senate string) string {
	length, deleted, prev, next, ans := len(senate), map[int]bool{}, 0, 1, byte('R')
	for length-len(deleted) > 1 {
		prev %= length
		next %= length
		for deleted[prev] {
			prev++
			prev %= length
		}
		for deleted[next] {
			next++
			next %= length
		}
		if senate[next] == senate[prev] {
			prev = next
			next++
		} else {
			deleted[next] = true
			prev = next + 1
			next += 2
		}
	}
	for i := 0; i < length; i++ {
		if !deleted[i] {
			ans = senate[i]
		}
	}
	if ans == 'D' {
		return "Dire"
	}
	return "Radiant"
}

/*
贪心 + 模拟
	贪心：对于议员，越是排在前面的位置越好，这样可以早早行使权利，同样地，当一个议员行使权利时，肯定也想优先废除对方排在靠前的位置的议员权利
	做法：维护两个切片radiant、dire，分别记录两个阵营的议员在一轮投票中行使权利的优先级（其实就是议员的位置）
	过程：比较两个切片首部的元素的大小，较小者先行使权利，他肯定是要废除对方阵营的第一个议员的权利，行使外权利以后，该议员本轮不能再行使权利，但是下一轮还是可以的，
			所以将其放入切片的末尾，此时他行使权利的优先级应该是下一轮的第一个，所以其优先级应该是原来的值 + 所有议员的数量
		 循环这个过程，直到一个切片为空
*/
func predictPartyVictory2(senate string) string {
	radiant, dire := []int{}, []int{}	//初始化两个切片，分别记录两个阵营的议员行使权利的优先级
	for i, s := range senate {
		if s == 'D' {
			dire = append(dire, i)
		} else {
			radiant = append(radiant, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {	//模拟行使权利的过程
		if radiant[0] < dire[0] {	//如果 radian 阵营的议员先行使权利，他肯定要废除 dire 阵营的第一个能够行使权利的议员，从而保护己方议员，同样的逻辑对于dire阵营也是一样
			radiant = append(radiant, radiant[0]+len(senate))	//将这个行使权利的议员从新入队，他在下一轮仍可能有机会行使权利，只不过优先级不再是当前的优先级，而是顺延 len(senate)
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
