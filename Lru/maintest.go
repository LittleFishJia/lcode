package main

import "fmt"

type Node struct {
	prev, next *Node
	key, value int
}

type LruCache struct {
	head, tail *Node
	cache      map[int]*Node
	capacity   int
}

func (l *LruCache) put(key int, value int) {
	targetNode, ok := l.cache[key]
	if ok {
		targetNode.value = value
		l.moveToHead(targetNode)
	} else {
		if l.capacity <= len(l.cache) {
			l.removeTail()
		}
		newNode := &Node{
			key:   key,
			value: value,
		}
		l.addToHead(newNode)
		l.cache[key] = newNode
	}
}

func (l *LruCache) get(key int) *Node {
	targetNode, ok := l.cache[key]
	if ok {
		l.moveToHead(targetNode)
		return targetNode
	}
	return nil
}

func (l *LruCache) addToHead(n *Node) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
}

func (l *LruCache) moveToHead(n *Node) {
	l.removeNode(n)
	l.addToHead(n)
}

func (l *LruCache) removeTail() {
	if l.tail.prev != l.head {
		tailNode := l.tail.prev
		l.removeNode(tailNode)
		delete(l.cache,tailNode.key)
	}
}

func (l *LruCache) removeNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func initLruCache(capacity int) *LruCache {
	l := &LruCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func main() {
	cache := initLruCache(2)
	cache.put(1,1)
	cache.put(2,2)
	cache.put(3,3)
	cache.put(4,4)
	fmt.Println(cache)
}
