# Trace GC 示例

用于生成带 GC 事件的 trace 文件，方便用 `go tool trace` 观察程序运行过程中的 GC。

## 方式一：跑测试生成 trace

```bash
cd trace_gc_demo
go test -trace=trace.out -run TestGCTrigger .
```

## 方式二：直接跑 main

```bash
cd trace_gc_demo
go run .
```

会在当前目录生成 `trace.out`。

## 查看 trace（含 GC）

```bash
go tool trace trace.out
```

浏览器会打开本地页面，可：

- 点 **View trace**：时间轴里能看到 **GC** 的紫色条（表示 GC 在运行）。
- 在底部图例里勾选/查看 **GC** 相关事件。

程序里通过大量小对象分配 + 显式 `runtime.GC()` 触发多次 GC，便于在 trace 里看到明显的 GC 段。
