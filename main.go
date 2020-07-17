package main

import (
	"./SingleLink"
	"./double_link"
	"./test"
	"fmt"
	"time"
)

func main1() {
	list := SingleLink.NewSingleLinkList()
	node1 := SingleLink.NewSingeleLinkNode(1)
	node2 := SingleLink.NewSingeleLinkNode(2)
	node3 := SingleLink.NewSingeleLinkNode(3)
	list.InsertNodeFront(node1)
	fmt.Println(list)
	list.InsertNodeFront(node2)
	fmt.Println(list)
	list.InsertNodeFront(node3)
	fmt.Println(list)

}

func main3() {
	list := SingleLink.NewSingleLinkList()
	node1 := SingleLink.NewSingeleLinkNode(1)
	node2 := SingleLink.NewSingeleLinkNode(2)
	node3 := SingleLink.NewSingeleLinkNode(3)
	node4 := SingleLink.NewSingeleLinkNode(4)
	list.InsertNodeBack(node1)
	fmt.Println(list)
	list.InsertNodeBack(node2)
	fmt.Println(list)
	list.InsertNodeBack(node3)
	fmt.Println(list)
	list.InsertNodeBack(node4)
	fmt.Println(list)
	//node := list.GetNodeAtIndex(1)
	//list.InsertNodeValueFront(9, node)
	//list.InsertNodeValueFront(8, node)
	//list.InsertNodeValueFront(7, node)
	//list.InsertNodeValueFront(6, node)
	//fmt.Println(list.InsertNodeValueBack(9,node))
	//fmt.Println(list)
	//fmt.Println(list.GetNodeAtIndex(0))
	//fmt.Println(list.GetMid())
	list.Reverse()
	fmt.Println(list)
	//fmt.Println(list.GetFirstNode())
	//list.DeleteAtIndex(2)
	//fmt.Println(list)
	//fmt.Println(list.GetNodeAtIndex(0))
	//fmt.Println(list.DeleteNode(node1))
	//fmt.Println(list)
}

func main2() {
	test.FindData("/Users/natee/go-data/yincheng/day4/140W某信用卡购物网数据/140W某信用卡购物网数据/140W某信用卡购物网数据.csv")
}

func main4() {
	list := double_link.NewDoubleLinkList()
	node1 := double_link.NewDoubleLinkNode(1)
	node2 := double_link.NewDoubleLinkNode(2)
	node3 := double_link.NewDoubleLinkNode(3)
	node4 := double_link.NewDoubleLinkNode(4)
	list.InsertBack(node1)
	fmt.Println(list)
	list.InsertBack(node2)
	fmt.Println(list)
	list.InsertBack(node3)
	fmt.Println(list)
	list.InsertBack(node4)
	fmt.Println(list)
}

func fbnq(n int,m map[int]int) int {
	if n <= 2{
		return n
	}else{
		var num1 int
		var num2 int
		if _,ok := m[n-1];ok{
			num1 = m[n-1]
		}else {
			num1 = fbnq(n-1,m)
		}
		if _,ok := m[n-2];ok{
			num1 = m[n-2]
		}else {
			num1 = fbnq(n-2,m)
		}
		return num1 + num2
	}
}

func main()  {
	var fbnqMap map[int]int
	fbnqMap = make(map[int]int)
	startTime := time.Now()
	fmt.Println(fbnq(40,fbnqMap))
	fmt.Println(time.Since(startTime))
}
