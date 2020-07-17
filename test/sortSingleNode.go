package test

import (
	"../SingleLink"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// ssaucykitty,pastori,libo55@dreamwiz.com


func FindData(file string) string {

	f, err := os.Open(file)
	if err != nil {
		return err.Error()
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	nodeList := SingleLink.NewSingleLinkList()

	i := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		lineStr := string(line)
		nodeStr := SingleLink.NewSingeleLinkNode(lineStr)
		nodeList.InsertNodeFront(nodeStr)
		i++
	}

	fmt.Println("内存载入完成")

	for ;;{
		fmt.Println("请输入要查询的用户名")
		var inputStr string
		fmt.Scanln(&inputStr)
		start:= time.Now()
		nodeList.FindString(inputStr)

		fmt.Println("本次查找需要",time.Since(start))
	}

}
