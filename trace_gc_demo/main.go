// 生成带 GC 事件的 trace，便于用 go tool trace 观察。
// 运行：go run . 然后 go tool trace trace.out
package main

import (
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()

	// 大量小对象分配，触发多次 GC
	for i := 0; i < 20; i++ {
		var garbage [][]byte
		for j := 0; j < 1e5; j++ {
			garbage = append(garbage, make([]byte, 256))
		}
		_ = garbage
		runtime.GC() // 也可依赖自动 GC，这里显式触发便于在 trace 里看到
	}
}
