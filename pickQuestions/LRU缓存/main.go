package main

/**

 请你设计并实现一个满足
 LRU (最近最少使用) 缓存 约束的数据结构。



 实现
 LRUCache 类：





 LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-
value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。




 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



 示例：


输入
["LRUCache", "put",   "put", "get", "put",  "get", "put", "get", "get", "get"]
[[2],        [1, 1], [2, 2], [1],   [3, 3], [2],   [4, 4], [1],   [3],   [4]]
输出
[null,        null,   null,  1,     null,   -1,    null,   -1,     3,     4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4




 提示：


 1 <= capacity <= 3000
 0 <= key <= 10000
 0 <= value <= 10⁵
 最多调用 2 * 10⁵ 次 get 和 put


 Related Topics 设计 哈希表 链表 双向链表 👍 2351 👎 0

*/

func main() {
	cache := Constructor(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//cache.Get(1)
	//cache.Put(3, 3)
	//cache.Get(2)
	//cache.Put(4, 4)
	//cache.Get(1)
	//cache.Get(3)
	//cache.Get(4)
	cache.Put(2, 1)
	cache.Put(1,1)
	cache.Put(2,3)
	cache.Put(4,1)
	cache.Get(1)
	cache.Get(2)
	cache.Get(2)
}

//	使用双向链表实现
/*
	我们采用双向链表的底层实现，有头尾节点，
	然后每次将操作过的节点放到头的位置，重新设置尾节点
*/

type LRUCache struct {
	capacity int
	m map[int]*Node
	head, tail *Node
}

type Node struct {
	Key int
	Value int
	Pre, Next *Node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.moveToHead(v)
		return v.Value
	}
	return -1
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) deleteNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) removeTail() int {
	node := this.tail.Pre
	this.deleteNode(node)
	return node.Key
}

func (this *LRUCache) addToHead(node *Node) {
	this.head.Next.Pre = node
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next = node
}

func (this *LRUCache) Put(key int, value int)  {
	if v, ok := this.m[key]; ok {
		v.Value = value
		this.moveToHead(v)
		return
	}

	if this.capacity == len(this.m) {
		rmKey := this.removeTail()
		delete(this.m, rmKey)
	}

	newNode := &Node{Key: key, Value: value}
	this.addToHead(newNode)
	this.m[key] = newNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		capacity: capacity,
		m: map[int]*Node{},
		head: head,
		tail: tail,
	}
}
