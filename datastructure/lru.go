package datastructure

// 最近使用的页面数据会在未来一段时期内仍然被使用,
// 已经很久没有使用的页面很有可能在未来较长的一段时间内仍然不会被使用.

type doubleNode struct {
	data string
	pre  *doubleNode
	next *doubleNode
}

type lru struct {
	len   int
	cap   int
	cache map[string]*doubleNode
	head  *doubleNode
	tail  *doubleNode
}

func newLRU(cap int) *lru {
	head, tail := &doubleNode{data: "dummy head"}, &doubleNode{data: "dummy tail"}
	head.next = tail
	tail.pre = head

	return &lru{
		len:   0,
		cap:   cap,
		cache: make(map[string]*doubleNode),
		head:  head,
		tail:  tail,
	}
}

func (r *lru) Dump() []string {
	res := make([]string, 0, r.cap)
	for n := r.head.next; n != r.tail; n = n.next {
		res = append(res, n.data)
	}
	return res
}

func (r *lru) Put(k, v string) {
	if r.len == r.cap {
		r.removeOne()
		delete(r.cache, k)
	}

	n, ok := r.cache[k]
	if !ok {
		node := r.addOne(v)
		r.cache[k] = node
		r.len += 1
		return
	}
	r.move2First(n, v)
}

func (r *lru) removeOne() {
	node := r.tail.pre
	if node == r.head {
		return
	}
	node.pre.next = r.tail
	r.tail.pre = node
}

func (r *lru) addOne(v string) *doubleNode {
	node := &doubleNode{
		data: v,
	}
	nextNode := r.head.next
	node.next = nextNode
	nextNode.pre = node
	r.head.next = node
	node.pre = r.head

	return node
}

func (r *lru) move2First(n *doubleNode, v string) {
	n.data = v
	if n.pre == r.head {
		return
	}

	n.pre.next = n.next
	n.next.pre = n.pre

	n.pre = r.head
	n.next = r.head.next
	r.head.next = n
}
