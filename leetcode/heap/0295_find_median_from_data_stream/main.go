// LeetCode 295 数据流的中位数
// 输入：每行 "add num" 或 "median"；median 时输出当前中位数（偶数个取两中间数平均值）
// 输出：每次 median 输出一行，保留一位小数或整数
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"leetcode/common_lib"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

// MinHeap 小顶堆（右半部分，较大的一半）
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MaxHeap 大顶堆：嵌入 MinHeap，只重写 Less
type MaxHeap struct {
	MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

type MedianFinder struct {
	left  *MaxHeap // 左半，大顶堆
	right *MinHeap // 右半，小顶堆
}

func Constructor() MedianFinder {
	left := &MaxHeap{}
	right := &MinHeap{}
	heap.Init(left)
	heap.Init(right)
	return MedianFinder{left: left, right: right}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.left, num)
	heap.Push(m.right, heap.Pop(m.left).(int))
	if m.left.Len() < m.right.Len() {
		heap.Push(m.left, heap.Pop(m.right).(int))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.left.Len() > m.right.Len() {
		return float64(m.left.MinHeap[0])
	}
	return float64(m.left.MinHeap[0]+(*m.right)[0]) / 2
}

func main() {
	defer out.Flush()

	finder := Constructor()
	for {
		line, err := common_lib.ReadNextValidLine(in)
		if err != nil && len(line) == 0 {
			break
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				continue
			}
			num, _ := strconv.Atoi(parts[1])
			finder.AddNum(num)
		case "median":
			med := finder.FindMedian()
			if med == float64(int(med)) {
				fmt.Fprintf(out, "%d\n", int(med))
			} else {
				fmt.Fprintf(out, "%.1f\n", med)
			}
		}
		if err != nil {
			break
		}
	}
}
