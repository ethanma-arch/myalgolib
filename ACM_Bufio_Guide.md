# Go ACM 模式高效 I/O (bufio) 汇总指南

在 Go 语言的算法笔试（ACM 模式）中，标准库 `fmt.Scan` 和 `fmt.Println` 在处理大规模数据（$>10^5$）时会因 IO 瓶颈导致 **TLE (Time Limit Exceeded)**。

**`bufio` 是必修课。** 本文档汇总了 `Scanner` 和 `Reader` 的核心用法与模板。

---

## 1. 核心对象与初始化

无论是 Scanner 还是 Reader，都要配合 `bufio.NewWriter` 使用以加速输出。

**通用结构：**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var (
    // 选择 Scanner 或 Reader 之一
    // in = bufio.NewScanner(os.Stdin) 
    // in = bufio.NewReader(os.Stdin)
    
    out = bufio.NewWriter(os.Stdout)
)

func main() {
    // 【必须】程序结束前刷新缓冲区，否则可能丢失输出
    defer out.Flush()
    
    // 你的逻辑...
}
```

---

## 2. 方案 A：bufio.Scanner (推荐首选)

**适用场景：**
*   读取一连串由**空格**或**换行**分隔的数据（整数、单词）。
*   不需要区分“第一行”、“第二行”，只关心数据流。
*   单行数据长度不超过 64KB。

### 2.1 初始化模板

```go
var in = bufio.NewScanner(os.Stdin)

func init() {
    // 【关键】设置为按单词分割，这就无视了空格和换行的区别
    in.Split(bufio.ScanWords)
}
```

### 2.2 常用辅助函数

**读取一个整数 (`readInt`)**

```go
func readInt() int {
    in.Scan() // 移动到下一个 token
    x, _ := strconv.Atoi(in.Text())
    return x
}
```

**读取一个字符串 (`readString`)**

```go
func readString() string {
    in.Scan()
    return in.Text()
}
```

**读取数组 (`N` + `N` 个数)**

```go
n := readInt()
nums := make([]int, n)
for i := 0; i < n; i++ {
    nums[i] = readInt()
}
```

### 2.3 解决 Scanner 的 64KB 限制

如果题目输入包含**超长字符串**（如 10万长的 DNA 序列），Scanner 默认会报错。需要手动扩容：

```go
func main() {
    // 必须在 Scan 之前调用
    buf := make([]byte, 2048)
    // 设置最大缓冲区为 10MB
    in.Buffer(buf, 10*1024*1024) 
    
    // ...
}
```

---

## 3. 方案 B：bufio.Reader (处理复杂/多行)

**适用场景：**
*   题目严格要求**按行读取**（如：Line 1 是 A 数组，Line 2 是 B 数组）。
*   需要自定义分隔符（如逗号分隔）。
*   读取**超长**的单行数据（无 64KB 限制）。

### 3.1 初始化

```go
var in = bufio.NewReader(os.Stdin)
```

### 3.2 读取整行 (模板)

通常配合 `strings` 包进行处理。

```go
import "strings"

// 读取一行，去除首尾空格和换行
func readNextValidLine(r *bufio.Reader) (string, error) {
    for {
        line, err := r.ReadString('\n')
        // 必须 Trim，否则末尾会有 '\n' 甚至 '\r'
        line = strings.TrimSpace(line)
		
        // 如果读到内容了，返回
        if len(line) > 0 {
			return line, nil
        }
		
		// 如果读到 EOF 且 line 为空，说明真的结束了
		if err != nil {
		    return "", err	
        }
        // 否则说明读到的是个空行，继续下一轮循环读
    }
	
}
```

### 3.3 读取一行由空格分开的数组

```go
func readLineAsInts(r *bufio.Reader) []int {
    line, _ := r.ReadString('\n')
    line = strings.TrimSpace(line)
    if line == "" {
        return []int{}
    }
    
    // Fields 自动处理连续空格
    parts := strings.Fields(line) 
    nums := make([]int, 0, len(parts))
    for _, p := range parts {
        val, _ := strconv.Atoi(p)
        nums = append(nums, val)
    }
    return nums
}
```

---

## 4. 高效输出 (bufio.Writer)

永远不要在循环里用 `fmt.Println`。

```go
// 输出字符串
fmt.Fprintln(out, "Hello World")

// 输出整数
fmt.Fprintf(out, "%d\n", 123)

// 输出数组（空格分隔）
for i, v := range nums {
    fmt.Fprintf(out, "%d", v)
    if i < len(nums)-1 {
        fmt.Fprint(out, " ")
    }
}
fmt.Fprintln(out) // 别忘了最后换行
```

---

## 5. 决策树：用 Scanner 还是 Reader？

| 特性 | bufio.Scanner | bufio.Reader |
| :--- | :--- | :--- |
| **分割方式** | 自动 (`ScanWords` 忽略换行/空格) | 手动 (`ReadString` 按行读) |
| **代码量** | 少 (适合大多数题目) | 多 (需处理分割、清洗) |
| **内存限制** | 默认 64KB (可扩容) | **无限制** (自动扩容) |
| **使用场景** | 读一堆乱序/顺序的数字、单词 | 读特定行、超长串、特殊分隔符 |
| **推荐指数** | ⭐⭐⭐⭐⭐ (默认首选) | ⭐⭐⭐ (特殊情况备用) |

---

## 6. 万能拷贝模板

### 模板 1：Scanner (解决 90% 题目)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var (
    in  = bufio.NewScanner(os.Stdin)
    out = bufio.NewWriter(os.Stdout)
)

func init() {
    in.Split(bufio.ScanWords)
    // 若题目数据极大，解开下行注释
    // in.Buffer(make([]byte, 4096), 10*1024*1024)
}

func readInt() int {
    in.Scan()
    x, _ := strconv.Atoi(in.Text())
    return x
}

func main() {
    defer out.Flush()
    
    // 示例：读取 N 以及 N 个数字
    for in.Scan() { // 循环直到 EOF
        n, _ := strconv.Atoi(in.Text()) // 或直接用 readInt
        // logic...
        fmt.Fprintln(out, n)
    }
}
```

### 模板 2：Reader (解决按行/逗号分隔)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var (
    in  = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)
)

func readLineAsInts() []int {
    line, _ := in.ReadString('\n')
    // 逗号分隔处理： line = strings.ReplaceAll(line, ",", " ")
    parts := strings.Fields(line)
    res := make([]int, 0, len(parts))
    for _, s := range parts {
        v, _ := strconv.Atoi(s)
        res = append(res, v)
    }
    return res
}

func main() {
    defer out.Flush()
    
    for {
        line, err := in.ReadString('\n')
        line = strings.TrimSpace(line)
        if len(line) == 0 && err != nil {
            break // EOF
        }
        if len(line) == 0 {
            continue // 跳过空行
        }
        
        // logic...
        fmt.Fprintln(out, line)
    }
}
```
