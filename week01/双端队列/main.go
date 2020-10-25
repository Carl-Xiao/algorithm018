package main

// 设计实现双端队列。
// 你的实现需要支持以下操作：

// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。

//双端队列  用前索引跟后索引确认当前数据
type MyCircularDeque struct {
	data  []int
	front int
	end   int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:  make([]int, k),
		front: 0,
		end:   0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	this.front = (len(this.data) + this.front - 1) % len(this.data) // 先循环左移一位
	this.data[this.front] = value                                   // 填入数据
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// 插入元素
	this.data[this.end] = value                // 填入数据
	this.end = (this.end + 1) % len(this.data) // 循环右移一位
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	this.front = (this.front + 1) % len(this.data) // 循环右移一位
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	this.end = (len(this.data) + this.end - 1) % len(this.data) // 循环左移一位
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {

	// 获取头部元素
	return this.data[this.front]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	return this.data[this.end]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.end == this.front
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.front+1)%len(this.data) == this.front
}

func main() {

}
