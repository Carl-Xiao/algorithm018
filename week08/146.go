package main

//实现LRU

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	//记录链表的长度
	size int
	//记录map的容量
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	l := &LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*Node{},
		head:     &Node{key: 0, value: 0},
		tail:     &Node{key: 0, value: 0},
	}
	l.head.next = l.tail
	l.tail.next = l.head
	return *l

}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.cache[key]; ok {
		//需要移动当前指针到队头
		this.moveToHead(value)
		return value.value
	}
	return -1
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) Put(key int, v int) {
	//存放节点
	//放在队列头
	//记录当前value
	if value, ok := this.cache[key]; ok {
		//已经存在就替换
		value.value = v
		this.moveToHead(value)
		this.size++
		if this.size > this.capacity {
			//淘汰一个不长使用的
			node := this.removeTail()
			//清理缓存
			delete(this.cache, node.key)
			this.size--
		}
	} else {
		//不存在就创建节点
		node := &Node{key: key, value: v}
		this.cache[key] = node
		this.addToHead(node)

	}
}
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
