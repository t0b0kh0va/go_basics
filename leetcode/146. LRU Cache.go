//https://leetcode.com/problems/lru-cache/
//#hashtable #design #linkedlist #doublylinkedlist

/*
   Solution is based on hashmap and double linked list.
   listNode represent each node of such list.
*/
type listNode struct {
	prev *listNode
	next *listNode
	val  int
	key  int
}

/*
   LRUCache contains:
   > m - map, which stores key to nodes of list;
   > head - fake head of linked list;
   > tail - fake tail of linked list;
   > cap - capacity;
   > len - actual amount of stored elements.
*/
type LRUCache struct {
	m    map[int]*listNode
	head *listNode
	tail *listNode
	cap  int
	len  int
}

/*
   In this case I use "fake head/tail" conteption.
   Fake nodes allow us to unify interaction with linked list,
   ant types of corner cases are covered by fake head/tail.
   It means my list always contains at least two nodes, in this
   way I should insert/remove/reorder only middle nodes of list,
   I shouldn't handle "hmmmm now length of list is 1 node, if i want
   insert element I should create a new one, mark it as a tail..."
*/
func Constructor(capacity int) LRUCache {
	fakeLast := &listNode{val: -1, prev: nil, next: nil}
	fakeHead := &listNode{val: -2, prev: nil, next: fakeLast}
	fakeLast.prev = fakeHead
	return LRUCache{
		cap:  capacity,
		len:  0,
		m:    make(map[int]*listNode, capacity),
		head: fakeHead,
		tail: fakeLast,
	}
}

/*
   Get function check if element represent in cache.
   First of all - look up at hashmap. If there is no neccessary element
   -> return -1. Otherwise we should find node of list (simple task,
   because map stores specific pointer) and then put it right after fake head.
   In case if we want to access first element -> we can do nothig, just return val.
*/
func (c *LRUCache) Get(key int) int {
	v, ok := c.m[key]
	if !ok {
		return -1
	}

	if c.head.next == v {
		return v.val
	}

	c.swapFirst(v)

	return v.val
}

/*
   swapFirst - put passed element right after fake head,
   some "magic" with pointers.
*/
func (c *LRUCache) swapFirst(v *listNode) {
	frst := c.head.next
	next := v.next
	prev := v.prev

	c.head.next = v
	v.prev = c.head
	prev.next = next
	next.prev = prev
	frst.prev = v
	v.next = frst
}

func (c *LRUCache) Put(key int, value int) {
	if c.Get(key) != -1 {
		c.head.next.val = value
		return
	}

	var cur *listNode
	if c.len >= c.cap {
		c.len--
		cur = c.tail.prev
		cur.prev.next = c.tail
		c.tail.prev = cur.prev
		delete(c.m, cur.key)
	} else {
		cur = &listNode{}
	}

	c.len++
	next := c.head.next
	cur.val = value
	cur.key = key
	cur.next = next
	cur.prev = c.head
	c.head.next = cur
	next.prev = cur
	c.m[key] = cur
}