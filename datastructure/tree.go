package datastructure

import (
	"errors"
)

const emptyMark string = "#"

type BinTreeNode struct {
	data       string
	leftChild  *BinTreeNode
	rightChild *BinTreeNode
}

func CreateNodeByHead(ds *dataSource) (createNode func() *BinTreeNode) {
	if ds == nil {
		return nil
	}

	return func() *BinTreeNode {
		v, err := ds.shift()
		if err != nil {
			return nil
		}
		if v == emptyMark {
			return nil
		}

		node := &BinTreeNode{data: v}
		node.leftChild = createNode()
		node.rightChild = createNode()
		return node
	}
}

func (b *BinTreeNode) DumpTail() []string {
	var buff []string
	dumpByTail := func() (dump func(b *BinTreeNode)) {
		return func(b *BinTreeNode) {
			if b == nil {
				return
			}
			dump(b.leftChild)
			dump(b.rightChild)
			buff = append(buff, b.data)
		}
	}
	dumpByTail()(b)
	return buff
}

func (b *BinTreeNode) DumpByMid() []string {
	var buff []string
	dumpByMid(&buff)(b)
	return buff
}

func dumpByMid(buff *[]string) (dump func(b *BinTreeNode)) {
	return func(b *BinTreeNode) {
		if b == nil {
			return
		}
		dump(b.leftChild)
		*buff = append(*buff, b.data)
		dump(b.rightChild)
	}
}

func (b *BinTreeNode) DumpByHead() []string {
	buff := []string{}
	dumpByHead(&buff)(b)
	return buff
}

func dumpByHead(buff *[]string) (dump func(b *BinTreeNode)) {
	return func(b *BinTreeNode) {
		if b == nil {
			return
		}
		*buff = append(*buff, b.data)
		dump(b.leftChild)
		dump(b.rightChild)
	}
}

type dataSource struct {
	data []string
}

func (ds *dataSource) shift() (string, error) {
	if len(ds.data) == 0 {
		return "", errors.New("empty")
	}
	v := ds.data[0]
	if len(ds.data) == 1 {
		ds.data = nil
		return v, nil
	}

	ds.data = ds.data[1:]
	return v, nil
}
