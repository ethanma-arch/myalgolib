package main

import (
	"os"
	"runtime"
	"runtime/trace"
	"testing"
)

// TestGCTrigger 会触发 GC。运行：
//
//	go test -trace=trace.out -run TestGCTrigger .
//
// 然后打开 trace：go tool trace trace.out
// 在网页里选 "View trace" 或 "GC" 相关视图即可看到 GC 事件。
func TestGCTrigger(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 20; i++ {
		var garbage [][]byte
		for j := 0; j < 1e5; j++ {
			garbage = append(garbage, make([]byte, 256))
		}
		_ = garbage
		runtime.GC()
	}
}
