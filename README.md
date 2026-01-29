MyAlgoLib
=========

这是一个用 Go 实现的算法与数据结构练习仓库，主要用于刷 LeetCode / ACM 及模拟大厂面试中的算法题。

目录结构（示例）：

- `leetcode/`：按题号分类的题解代码，例如：
  - `leetcode/0026_remove_duplicates_from_sorted_array/`
  - `leetcode/0027_remove_element/`
  - `leetcode/0169_majority-element/`

每个题目目录下一般包含：

- `main.go`：题目的可运行解法（ACM 风格，从标准输入读、向标准输出写）。
- `input.txt`（可选）：本地调试用的输入样例，通过重定向到标准输入运行。

运行示例
--------

以 `0026_remove_duplicates_from_sorted_array` 为例，在项目根目录执行：

```bash
cd leetcode/0026_remove_duplicates_from_sorted_array
go run main.go < input.txt
```

或直接在评测环境中，让评测系统把用例通过标准输入喂给程序。

环境要求
--------

- Go 1.20 或以上版本
- 已配置好 GOPATH / Go 环境变量

后续规划
--------

- 持续补充更多 LeetCode / ACM 题解
- 根据真实面试需求，整理一套通用的 Go ACM 模版

