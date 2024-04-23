//package main
//
//import "fmt"
//
//type Node struct {
//	key, value int
//	prev, next *Node
//}
//
//type LRUCache struct {
//	cache map[int]*Node
//	capacity int
//	head, tail *Node
//}
//
//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		cache: make(map[int]*Node),
//		capacity: capacity,
//		head: &Node{},
//		tail: &Node{},
//	}
//	l.head.next = l.tail
//	l.tail.prev = l.head
//	return l
//}
//
//func (this *LRUCache) addToHead(node *Node) {
//	node.prev = this.head
//	node.next = this.head.next
//	this.head.next = node
//	this.head.next.prev = node
//}
//
//func (this *LRUCache) removeNode(node *Node) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//}
//
//func (this *LRUCache) moveToHead(node *Node) {
//	this.removeNode(node)
//	this.addToHead(node)
//}
//
//func (this *LRUCache) removeTail() *Node {
//	node := this.tail.prev
//	this.removeNode(node)
//	return node
//}
//
//func (this *LRUCache) Get(key int) int {
//	if node, ok := this.cache[key]; ok {
//		this.moveToHead(node)
//		return node.value
//	}
//	return -1
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if node, ok := this.cache[key]; ok {
//		node.value = value
//		this.moveToHead(node)
//	} else {
//		node := &Node{key: key, value: value}
//		this.cache[key] = node
//		this.addToHead(node)
//		if len(this.cache) > this.capacity {
//			tail := this.removeTail()
//			delete(this.cache, tail.key)
//		}
//	}
//}
//
//func main() {
//	cache := Constructor(2)
//	cache.Put(1, 1)
//	cache.Put(2, 2)
//	fmt.Println(cache.Get(1)) // 输出 1
//	cache.Put(3, 3)
//	fmt.Println(cache.Get(2)) // 输出 -1
//	cache.Put(4, 4)
//	fmt.Println(cache.Get(1)) // 输出 -1
//	fmt.Println(cache.Get(3)) // 输出
//}