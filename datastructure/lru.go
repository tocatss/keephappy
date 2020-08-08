package datastructure

import (
	"errors"
	"fmt"
	"log"
)

// 最近使用的页面数据会在未来一段时期内仍然被使用,
// 已经很久没有使用的页面很有可能在未来较长的一段时间内仍然不会被使用.

// LRU: 有长度和容量，并可以添加元素和取得元素
// 添加元素：
// 1. 添加新元素： 在头部添加，在size不够的时候需要从尾部删除。=> 需要记录头和尾。
// 2. 添加已存在元素： 将该元素移动到头 （为了快速找到元素选择map结构，为了快速移动选择双向链表。
// 取得元素：选择map结构

type dNode struct {
	data string
	pre  *dNode
	next *dNode
}

type lruModel struct {
	len   int
	cap   int
	cache map[string]*dNode
	head  *dNode
	tail  *dNode
}

func NewLRUModel(cap int) *lruModel {
	head := &dNode{}
	tail := &dNode{}

	head.next = tail
	tail.pre = head
	return &lruModel{
		len:   0,
		cap:   cap,
		cache: make(map[string]*dNode),
		head:  head,
		tail:  tail,
	}
}

func (m *lruModel) Get(k string) (string, error) {
	if n, ok := m.cache[k]; ok {
		return n.data, nil
	}
	return "", fmt.Errorf("key: %s not found ", k)
}

func (m *lruModel) Put(v string) {
	if n, ok := m.cache[v]; ok {
		// move to head.
		m.move2Head(n)
		return
	}

	n := &dNode{
		data: v,
	}
	if m.len == m.cap {
		m.deleteFromTail()
	}
	if err := m.add2Head(n); err != nil {
		log.Print(err)
	}
	return
}

func (m *lruModel) Dump() []string {
	res := make([]string, 0, m.len)
	for n := m.head.next; n != m.tail; n = n.next {
		res = append(res, n.data)
	}
	return res
}

func (m *lruModel) move2Head(n *dNode) {
	if m.head.next == n {
		return
	}

	n.pre.next = n.next
	n.next.pre = n.pre
	n.next = m.head.next
	n.pre = m.head

	m.head.next.pre = n
	m.head.next = n
}

func (m *lruModel) deleteFromTail() {
	if m.tail.pre == m.head {
		return
	}

	deleteNode := m.tail.pre
	deleteNode.pre.next = m.tail
	m.tail.pre = deleteNode.pre

	delete(m.cache, deleteNode.data)
	m.len--
}

func (m *lruModel) add2Head(n *dNode) error {
	if n == nil {
		return nil
	}
	if m.len == m.cap {
		return errors.New("")
	}

	n.next = m.head.next
	n.pre = m.head
	m.head.next.pre = n
	m.head.next = n

	m.len++
	m.cache[n.data] = n
	return nil
}
