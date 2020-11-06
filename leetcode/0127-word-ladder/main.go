package main

import (
	"fmt"
	"math"
)

/**
127. 单词接龙
	https://leetcode-cn.com/problems/word-ladder/
题目描述：
	给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
	每次转换只能改变一个字母。
	转换过程中的中间单词必须是字典中的单词。
说明:
	如果不存在这样的转换序列，返回 0。
	所有单词具有相同的长度。
	所有单词只由小写字母组成。
	字典中不存在重复的单词。
	你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:
	输入:
	beginWord = "hit",
	endWord = "cog",
	wordList = ["hot","dot","dog","lot","log","cog"]

	输出: 5

	解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
		返回它的长度 5。
示例 2:
	输入:
	beginWord = "hit"
	endWord = "cog"
	wordList = ["hot","dot","dog","lot","log"]

	输出: 0

	解释: endWord "cog" 不在字典中，所以无法进行转换。
相似题目
	https://leetcode-cn.com/problems/word-ladder-ii/
	https://leetcode-cn.com/problems/minimum-genetic-mutation/
*/
func main() {
	var (
		beginWord, endWord string
		wordList           []string
	)
	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
	fmt.Println(ladderLength2(beginWord, endWord, wordList))
	fmt.Println(ladderLength3(beginWord, endWord, wordList))
	fmt.Println(ladderLength4(beginWord, endWord, wordList))
}

/**
官方题解
	本题要求的是最短转换序列的长度，看到最短首先想到的就是广度优先搜索。想到广度优先搜索自然而然的就能想到图，但是本题并没有直截了当的给出图的模型，因此我们需要把它抽象成图的模型。
	我们可以把每个单词都抽象为一个点，如果两个单词可以只改变一个字母进行转换，那么说明他们之间有一条双向边。因此我们只需要把满足转换条件的点相连，就形成了一张图。
	基于该图，我们以 beginWord 为图的起点，以 endWord 为终点进行广度优先搜索，寻找 beginWord 到 endWord 的最短路径。
算法
	基于上面的思路我们考虑如何编程实现。
	首先为了方便表示，我们先给每一个单词标号，即给每个单词分配一个 id。创建一个由单词 word 到 id 对应的映射 wordId，并将 beginWord 与 wordList 中所有的单词都加入这个映射中。之后我们检查 endWord 是否在该映射内，若不存在，则输入无解。我们可以使用哈希表实现上面的映射关系。
	然后我们需要建图，依据朴素的思路，我们可以枚举每一对单词的组合，判断它们是否恰好相差一个字符，以判断这两个单词对应的节点是否能够相连。但是这样效率太低，我们可以优化建图。
	具体地，我们可以创建虚拟节点。对于单词 hit，我们创建三个虚拟节点 *it、h*t、hi*，并让 hit 向这三个虚拟节点分别连一条边即可。如果一个单词能够转化为 hit，那么该单词必然会连接到这三个虚拟节点之一。对于每一个单词，我们枚举它连接到的虚拟节点，把该单词对应的 id 与这些虚拟节点对应的 id 相连即可。
	最后我们将起点加入队列开始广度优先搜索，当搜索到终点时，我们就找到了最短路径的长度。注意因为添加了虚拟节点，所以我们得到的距离为实际最短路径长度的两倍。同时我们并未计算起点对答案的贡献，所以我们应当返回距离的一半再加一的结果。
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{} //存储单词到id的映射
	graph := [][]int{}         //存储单词间的转换关系
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

/**
看完官方题解后自己梳理一遍
	建模，利用图表示出所有单词间的关系，为了方便，添加虚拟节点，例如对于单词 dot，额外添加 *ot, d*t, do*
	使用邻接表存储
	wordId[string]int 存储单词和 id 之间的映射
	graph[][]int 存储顶点之间的关系
*/

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	var (
		wordId = make(map[string]int)	//所有顶点的集合，也是单词到id的映射表
		graph  = make([][]int, 0)	//而为数组存储图，采用的是邻接表存储，而非邻接矩阵
	)
	addVertice := func(word string) int { //添加顶点
		id, has := wordId[word]	//获取该节点对应的id
		if !has {	//如果该节点第一次出现，未包含在顶点集合里，将该顶点放入顶点集合
			id = len(wordId)
			wordId[word] = id	//单词映射成id
			graph = append(graph, []int{})	//为该顶点创建一个假链表，存储其可以到达的顶点
		}
		return id
	}

	addEdge := func(word string) int {
		id1 := addVertice(word)	//添加顶点，并返回该顶点的 id
		s := []byte(word)
		for index, char := range s { //添加虚拟节点
			s[index] = '*'
			id2 := addVertice(string(s))
			graph[id1] = append(graph[id1], id2) //添加顶点到其衍生的虚拟顶点的边
			graph[id2] = append(graph[id2], id1) //添加衍生顶点到原顶点的边
			s[index] = char
		}
		return id1
	}
	for _, word := range wordList {	//建模，将 wordList 中所有的单词存储到图中
		addEdge(word)
	}
	beginId := addEdge(beginWord) //将起始单词也存入图中，使其作为搜索的开始点
	endId, has := wordId[endWord] //获取结束单词对应的id
	if !has {
		return 0
	}
	const inf int = math.MaxInt16 //最大整数
	dist := make([]int, len(wordId)) //dist记录从开始顶点到目标顶点经过的距离
	for id := range dist {	//每个点到目标顶点的距离初始化为最大整数，同时也标记该顶点未被访问过
		dist[id] = inf
	}
	queue := []int{beginId}	//广度优先搜索的那个队列，起始只有开始的那个顶点
	dist[beginId] = 0	//不把起始顶点的那段距离算入，从它能到达的节点才开始计算距离，所以返回之前记得 + 1
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == endId {	//搜索到了目标顶点
			return dist[endId]/2 + 1
		}
		for _, next := range graph[curr] {	//遍历当前顶点能到达（未被访问过）的所有顶点，将其放入广搜队列
			if dist[next] == inf {
				dist[next] = dist[curr] + 1
				queue = append(queue, next)
			}
		}
	}
	return 0
}


//////////////////////////////////////////////下面是解答去看到的答案////////////////////////////////////////////////
//bfs,标准模板
//求无向图最短路径问题
func ladderLength3(beginWord string, endWord string, wordList []string) int {
	//这种解法如果wordlist很大就会超时
	wdict := make(map[string][]int)

	for id, wd := range wordList {
		for i := range wd {
			k := wd[:i] + "*" + wd[i + 1:]
			if _,ok := wdict[k];!ok {
				wdict[k] = []int{}
			}
			wdict[k] = append(wdict[k],id)
		}
	}
	wordList = append(wordList,beginWord)
	q := []int{len(wordList) - 1}
	used,l := make([]bool,len(wordList)),1
	for len(q) > 0 {
		nextQ := []int{}
		l ++
		for _,i := range q {
			w := wordList[i]
			for i := range w {
				k := w[:i] + "*" + w[i + 1:]
				for _,wid := range wdict[k] {
					if wordList[wid] == endWord {
						return l
					}
					if !used[wid] {
						used[wid] = true
						nextQ = append(nextQ,wid)
					}
				}
			}
		}
		q = nextQ
	}
	return 0
}

//双向BFS
func ladderLength4(beginWord string, endWord string, wordList []string) int {
	step := 0 

	wordMap := make(map[string]int)

	for i,w := range wordList {
		wordMap[w] = i
	}
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	startUsed := make([]bool,len(wordList))
	endUsed := make([]bool,len(wordList))
	if i,ok := wordMap[endWord];!ok {
		return 0 
	} else {
		endUsed[i] = true
	}
	for len(startQueue) > 0 {
		step ++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			chars := []byte(startQueue[i])
			for j := 0; j < len(chars); j++ {
				o := chars[j]

				for c := 'a'; c <= 'z'; c++ {
					chars[j] = byte(c)
					idw,ok := wordMap[string(chars)]

					if !ok || startUsed[idw] {
						continue
					}

					if endUsed[idw] {
						return step + 1
					} else {
						startQueue = append(startQueue,string(chars))
						startUsed[idw] = true
					}
				}

				chars[j] = o
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue,endQueue = endQueue,startQueue
			startUsed,endUsed = endUsed,startUsed
		}
	}
	return 0
}