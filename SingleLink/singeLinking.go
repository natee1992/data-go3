package SingleLink

import "fmt"

type Singele interface {
	String() string
}

// 涟表节点
type SingleLinkNode struct {
	value interface{}
	pNext *SingleLinkNode
}

// 构造一个节点
func NewSingeleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data, nil}
}

// 返回数据
func (node *SingleLinkNode) Value() interface{} {
	return node.value
}

// 返回节点
func (node *SingleLinkNode) PNext() *SingleLinkNode {
	return node.pNext
}

func (node *SingleLinkNode) String() string {
	return fmt.Sprintf("value:%v,pNext:%v", node.value, &node.pNext)
}
