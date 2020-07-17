package SingleLink

import (
	"fmt"
	"strings"
)

// 单链表接口
type SingleLink interface {
	// 增删查改
	GetFirstNode() *SingleLinkNode            //抓取头部节点
	InsertNodeFront(node *SingleLinkNode)     // 头部插入
	InsertNodeBack(node *SingleLinkNode)      // 尾部插入
	GetNodeAtIndex(index int) *SingleLinkNode // 根据索引获取指定位置的节点
	DeleteNode(dest *SingleLinkNode) bool     // 删除一个节点
	DeleteAtIndex(index int)
	String() string
	InsertNodeValueFront(dest *SingleLinkNode, node *SingleLinkNode) bool
	InsertNodeValueBack(dest *SingleLinkNode, node *SingleLinkNode) bool
}

// 链表结构
type SingleLinkList struct {
	head   *SingleLinkNode //涟表头指针
	length int             //涟表长度
}

func NewSingleLinkList() *SingleLinkList {

	return &SingleLinkList{&SingleLinkNode{nil, nil}, 0}
}

func (nodeList *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return nodeList.head
}

func (nodeList *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if nodeList.length == 0 {
		nodeList.head = node
	} else {
		head := nodeList.head
		node.pNext = head
		nodeList.head = node
	}
	nodeList.length++
}

func (nodeList *SingleLinkList) InsertNodeBack(node *SingleLinkNode) {
	if nodeList.length == 0 {
		nodeList.head = node
	} else {
		current := nodeList.head
		for {
			if current.pNext == nil {
				current.pNext = node
				break
			} else {
				current = current.pNext
			}
		}
	}
	nodeList.length++
}

func (nodeList *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	current := nodeList.head
	for i := 0; i < nodeList.length; i++ {
		if i == index {
			return current
		} else {
			current = current.pNext
		}
	}
	return nil
}

func (nodeList *SingleLinkList) DeleteAtIndex(index int) {

	if nodeList.length == 1 {
		nodeList.head = nil
		nodeList.length--
	} else if nodeList.length > 1 {
		if index == 0 {
			head := nodeList.head
			nodeList.head = head.pNext
			nodeList.length--
		} else {
			current := nodeList.head
			for i := 0; i < nodeList.length; i++ {
				if i == index-1 {
					current.pNext = current.pNext.pNext
					nodeList.length--
				} else {
					current = current.pNext
				}
			}
		}
	}
}

func (nodeList *SingleLinkList) DeleteNode(node *SingleLinkNode) bool {
	if nodeList.length == 1 {
		if nodeList.head == node {
			nodeList.head = nil
			nodeList.length--
			return true
		} else {
			return false
		}
	} else if nodeList.length > 1 {
		if nodeList.head == node {
			nodeList.head = nodeList.head.pNext
			nodeList.length--
			return true
		} else {
			current := nodeList.head
			for current.pNext != nil {
				if current.pNext == node {
					current.pNext = node.pNext
					nodeList.length--
					return true
				}
			}
		}
	}
	return false
}

func (nodeList *SingleLinkList) String() string {
	str := ""
	current := nodeList.head
	for i := 0; i < nodeList.length+1; i++ {
		if current != nil {
			str += fmt.Sprintf("%v-->", current.value)
			current = current.pNext
		}
	}
	str += fmt.Sprintf("nil")
	return str
}

func (nodeList *SingleLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {
	destNode := NewSingeleLinkNode(dest)
	if nodeList.length >= 1 {
		if nodeList.length == 1 {
			nodeList.head = destNode
			nodeList.head.pNext = node
			nodeList.length++
			return true
		} else {
			current := nodeList.head
			if node == current {
				destNode.pNext = current
				nodeList.head = destNode
				nodeList.length++
				return true
			} else {
				for current.pNext != nil {
					if current.pNext == node {
						destNode.pNext = current.pNext
						current.pNext = destNode
						nodeList.length++
						return true
					} else {
						current = current.pNext
					}
				}
			}
		}
	}
	return false
}

func (nodeList *SingleLinkList) InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool {
	destNode := NewSingeleLinkNode(dest)
	current := nodeList.head
	for current != nil {
		if current == node {
			destNode.pNext = current.pNext
			current.pNext = destNode
			nodeList.length++
			return true
		} else {
			current = current.pNext
		}
	}
	return false
}

func (nodeList *SingleLinkList) FindString(data string) {
	phead := nodeList.head
	for phead.pNext != nil {
		if strings.Contains(phead.value.(string), data) {
			fmt.Println(phead.value.(string), "已找到")
		}
		phead = phead.pNext
	}
}

func (nodeList *SingleLinkList) GetMid() *SingleLinkNode {
	if nodeList.head.pNext == nil {
		return nil
	} else {
		phead1 := nodeList.head
		phead2 := nodeList.head
		for phead2 != nil && phead2.pNext != nil {
			phead1 = phead1.pNext
			phead2 = phead2.pNext.pNext
		}
		return phead1
	}
}

// 链表反转
func (nodeList *SingleLinkList) Reverse() {
	if nodeList.head == nil || nodeList.head.pNext == nil {
		return
	}
	var pre *SingleLinkNode  // 前面节点
	var cur = nodeList.head.pNext //当前节点
	for cur != nil{
		curNext := cur.pNext //后续节点
		cur.pNext = pre //反转第一步
		pre =cur //持续推进
		cur=curNext//持续推进
	}
	fmt.Println(pre)
	nodeList.head=pre
}
