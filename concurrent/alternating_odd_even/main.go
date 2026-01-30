// 题目：两个 goroutine 交替打印奇数和偶数，从 1 打印到 N。
// 奇数由 goroutine A 打印，偶数由 goroutine B 打印，顺序为 1,2,3,4,...,N。
// ACM 格式：从 stdin 读入 N，向 stdout 输出。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	// 读一行，解析为 N
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)
	if n <= 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	oddCh := make(chan struct{})
	evenCh := make(chan struct{})

	go func() {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-oddCh
			fmt.Fprintln(out, i)
			if i+1 <= n {
				evenCh <- struct{}{}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-evenCh
			fmt.Fprintln(out, i)
			if i+1 <= n {
				oddCh <- struct{}{}
			}
		}
	}()

	oddCh <- struct{}{}
	wg.Wait()
}
