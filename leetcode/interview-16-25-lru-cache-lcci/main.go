package main

import "fmt"

/*
面试题 16.25. LRU 缓存
	https://leetcode-cn.com/problems/lru-cache-lcci/
题目描述
	设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。
	它应该支持以下操作： 获取数据 get 和 写入数据 put 。
	获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
	写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

示例:
	LRUCache cache = new LRUCache(2//缓存容量);
	cache.put(1, 1);
	cache.put(2, 2);
	cache.get(1);       // 返回  1
	cache.put(3, 3);    // 该操作会使得密钥 2 作废
	cache.get(2);       // 返回 -1 (未找到)
	cache.put(4, 4);    // 该操作会使得密钥 1 作废
	cache.get(1);       // 返回 -1 (未找到)
	cache.get(3);       // 返回  3
	cache.get(4);       // 返回  4
*/
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

//添加一个节点到首部
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
	node.next.prev = node
}

//删除某个节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

//移动某个节点到首部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

//删除最后一个节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode, capacity),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key] //从cache中获取key对应到节点
	if !ok {                    //节点不存在直接返回 -1
		return -1
	} else { //节点存在，将其移动到首部，并返回节点的 value
		this.moveToHead(node)
		return node.value
	}
}

/*
put 操作
先判断缓存中是否已存在
	如果存在，修改节点的value，并移动到首部
	如果不存在
		判断当前缓存的容量
		容量已满，删除最后一个节点再添加
*/
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok { //缓存中已存在该key，更新，并移动到缓存列表首部
		node.value = value
		this.moveToHead(node)
	} else { //不存在，新添加
		node = &DLinkedNode{
			key:   key,
			value: value,
		}
		this.cache[key] = node
		this.size++
		this.addToHead(node)
		if this.size > this.capacity {
			node = this.removeTail()
			delete(this.cache, node.key)
			this.size--
		}
	}
}
