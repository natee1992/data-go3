package double_link

import "fmt"

type DoubleLinkInterface interface {
	// 增删查改
	GetFirstNode() *DoubleLinkNode            //抓取头部节点
	InsertNodeFront(node *DoubleLinkNode)     // 头部插入
	InsertNodeBack(node *DoubleLinkNode)      // 尾部插入
	GetNodeAtIndex(index int) *DoubleLinkNode // 根据索引获取指定位置的节点
	DeleteNode(dest *DoubleLinkNode) bool     // 删除一个节点
	DeleteAtIndex(index int)
	String() string
	InsertNodeValueFront(dest *DoubleLinkNode, node *DoubleLinkNode) bool
	InsertNodeValueBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool
}

type DoubleLinkList struct {
	head   *DoubleLinkNode
	length int
}

func (nodeList *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return nodeList.head.next
}

// 新建一个双链表
func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head: head, length: 0}
}

func (nodeList *DoubleLinkList) GetLength() int {
	return nodeList.length
}

func (nodeList *DoubleLinkList) InsertHead(node *DoubleLinkNode) {
	phead := nodeList.head
	if phead.next == nil {
		node.next = nil
		phead.next = node
		node.prev = phead
	} else {
		phead.next.prev = node
		node.next = phead.next
		phead.next = node
		node.prev = phead
	}
	nodeList.length++
}

func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode) {
	phead := dlist.head //头节点
	if phead.next == nil {
		node.next = nil   //nil
		phead.next = node //只有一个节点直接链接上
		node.prev = phead //上一个节点
		dlist.length++
	} else {
		for phead.next != nil {
			phead = phead.next //循环下去
		}
		phead.next = node //后缀
		node.prev = phead //前缀

		dlist.length++
	}
}

func (nodeList *DoubleLinkList) String() string {
	var listStr1 string
	var listStr2 string
	phead := nodeList.head
	for phead.next != nil {
		listStr1 += fmt.Sprintf("%v-->", phead.next.value)
		phead = phead.next
	}
	listStr1 += fmt.Sprintf("nil")
	listStr1 += "\n"
	listStr2 = "nil"
	for phead != nodeList.head {
		listStr2 += fmt.Sprintf("<--%v", phead.value)
		phead = phead.prev
	}
	return listStr1 + listStr2 + "\n"
}
