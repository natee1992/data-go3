package double_link

import "fmt"

type DoubleLinkNode struct {
	prev  *DoubleLinkNode
	value interface{}
	next *DoubleLinkNode
}

func (node *DoubleLinkNode) String() string {
	return fmt.Sprintf("prev: %v, value: %v, next: %v", node.prev, node.value, node.next)
}

func NewDoubleLinkNode(data interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{nil, data, nil}
}

func (node *DoubleLinkNode) Prev() *DoubleLinkNode {
	return node.prev
}

func (node *DoubleLinkNode) Next() *DoubleLinkNode {
	return node.next
}

func (node *DoubleLinkNode) Value() interface{} {
	return node.value
}


