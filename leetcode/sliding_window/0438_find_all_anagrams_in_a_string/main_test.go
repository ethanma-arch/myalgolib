package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	// 定义测试用例结构体
	tests := []struct {
		name     string // 用例名称，用于报错时区分
		s        string // 输入 s
		p        string // 输入 p
		expected []int  // 期望输出
	}{
		{
			name:     "Example 1",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "Example 2",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			name:     "No Match",
			s:        "hello",
			p:        "world",
			expected: []int{}, // 期望空切片
		},
		{
			name:     "s is shorter than p",
			s:        "ab",
			p:        "abc",
			expected: []int{},
		},
		{
			name:     "Same characters repeated",
			s:        "aaaaa",
			p:        "aa",
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "Empty s",
			s:        "",
			p:        "a",
			expected: []int{},
		},
	}

	// 遍历所有测试用例
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 调用你的算法函数
			got := main(tc.s, tc.p)

			// 【关键处理】处理 nil 切片和空切片 []int{} 不相等的问题
			// 如果算法返回 nil，而期望是 []int{}，reflect.DeepEqual 会认为不相等
			// 所以这里做一个归一化处理
			if len(got) == 0 && len(tc.expected) == 0 {
				return // 都是空的，算通过
			}

			// 使用反射比较切片内容是否一致
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("\nCase: %s\nInput: s=\"%s\", p=\"%s\"\nExpected: %v\nGot:      %v",
					tc.name, tc.s, tc.p, tc.expected, got)
			}
		})
	}
}
