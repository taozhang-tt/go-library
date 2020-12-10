package main

import "fmt"

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

//双向链表的节点
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

/*
size: 当前缓存的数据量
capacity: 缓存的容量
cache: 缓存的内容
head, tail: 双向链表的头尾节点，维护缓存的数据被访问的顺序
*/
type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkedNode
	head, tail     *DLinkedNode
}
//从双向链表中删除一个节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
//在链表的头部添加一个节点
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.prev = this.head
	node.next.prev = node
	this.head.next = node
}
//删除链表的最后一个节点
func (this *LRUCache) removeTail() (node *DLinkedNode) {
	node = this.tail.prev
	this.removeNode(node)
	return node
}
//创建缓存并初始化，为了方便操作，双向链表添加链空的头尾节点
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		size: 0,
		capacity: capacity,
		cache: map[int]*DLinkedNode{},
		head: &DLinkedNode{},
		tail: &DLinkedNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}
/*
从缓存中获取key对应的value
	如果key不存在，直接返回 -1
	如果key存在，将key对应的节点移动到双向列表的首部（其实就是先删除，再添加到首部）
*/
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.removeNode(node)
		this.addToHead(node)
		return node.value
	}
}

/*
向缓存中存入新key
	如果 key 存在
		更新 key 对应的 value，将 key 对应的节点移动到双向链表的首部
	如果 key 不存在
		添加一个新节点到 cache 和双向链表中
		添加完后要判断缓存的数据是不是超过了容量，如果超过容量，需要删除最后一个节点
*/
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.removeNode(node)
		this.addToHead(node)
	} else {
		node = &DLinkedNode{
			key: key,
			value: value,
		}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			node = this.removeTail()
			delete(this.cache, node.key)
			this.size--
		}
	}
}
