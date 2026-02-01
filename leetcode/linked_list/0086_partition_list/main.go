package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readNextValidLine(r *bufio.Reader) (string, error) {
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			return line, nil
		}
		if err != nil {
			return "", err
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func solve(head *ListNode, x int) *ListNode {
	smallDummy := &ListNode{Val: 0}
	largeDummy := &ListNode{Val: 0}

	cur := head
	small := smallDummy
	large := largeDummy

	for cur != nil {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			large.Next = cur
			large = large.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}

	small.Next = largeDummy.Next
	return smallDummy.Next
}

func main() {
	defer out.Flush()

	for {
		// 读前两行
		line1, err1 := readNextValidLine(in)
		if err1 != nil && len(line1) == 0 {
			break
		}
		line2, err2 := readNextValidLine(in)
		if err2 != nil && len(line2) == 0 {
			break
		}

		// 构造链表
		parts := strings.Fields(line1)
		dummy := &ListNode{}
		tail := dummy
		for _, p := range parts {
			val, _ := strconv.Atoi(p)
			tail.Next = &ListNode{Val: val}
			tail = tail.Next
		}

		// 计算结果
		x, _ := strconv.Atoi(strings.TrimSpace(line2))
		head := solve(dummy.Next, x)
		for head != nil {
			fmt.Fprint(out, head.Val)
			if head.Next != nil {
				fmt.Fprint(out, " ")
			}
			head = head.Next
		}
		fmt.Fprintln(out)
		if err1 == io.EOF || err2 == io.EOF {
			break
		}
	}
}
