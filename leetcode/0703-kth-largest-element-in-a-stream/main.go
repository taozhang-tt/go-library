package main

import (
	"container/heap"
)

/*
703. 数据流中的第 K 大元素
	https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
题目描述
	设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
	请实现 KthLargest 类：
	KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
	int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
示例：

输入：
	["KthLargest", "add", "add", "add", "add", "add"]
	[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
	输出：
	[null, 4, 5, 5, 8, 8]
解释：
	KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
	kthLargest.add(3);   // return 4
	kthLargest.add(5);   // return 5
	kthLargest.add(10);  // return 5
	kthLargest.add(9);   // return 8
	kthLargest.add(4);   // return 8
*/
func main() {

}

type KthLargest struct {
	data []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Len() int {
	return len(kl.data)
}

func (kl *KthLargest) Swap(i, j int) {
	kl.data[i], kl.data[j] = kl.data[j], kl.data[i]
}

func (kl *KthLargest) Less(i, j int) bool {
	return kl.data[i] < kl.data[j]
}

func (kl *KthLargest) Push(v interface{}) {
	kl.data = append(kl.data, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.data
	v := a[len(a)-1]
	kl.data = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.data[0]
}
