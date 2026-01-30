package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readString() string {
	in.Scan()
	return in.Text()
}

func readInt() int {
	in.Scan()
	num, _ := strconv.Atoi(in.Text())
	return num
}

type RandomizedSet struct {
	indexMap map[int]int
	elements []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indexMap: make(map[int]int),
		elements: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; ok {
		return false
	}

	// 2. 放入数组末尾
	this.elements = append(this.elements, val)
	// 3. 记录下标
	this.indexMap[val] = len(this.elements) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// 1. 检查是否存在
	id, ok := this.indexMap[val]
	if !ok {
		return false
	}

	// 2. 获取数组最后一个元素
	lastIndex := len(this.elements) - 1
	lastVal := this.elements[lastIndex]

	// 3. 将最后一个元素填入被删除元素的位置 (覆盖)
	this.elements[id] = lastVal

	// 4. 重要：更新 lastVal 在 map 中的下标
	this.indexMap[lastVal] = id

	// 5. 移除数组末尾元素
	this.elements = this.elements[:lastIndex]

	// 6. 从 map 中删除 val
	delete(this.indexMap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.elements))
	return this.elements[idx]
}

func main() {
	defer out.Flush()
	in.Split(bufio.ScanWords)

	obj := Constructor()

	for in.Scan() {
		op := in.Text()

		switch op {
		case "insert":
			val := readInt()
			res := obj.Insert(val)
			fmt.Fprintln(out, res)
		case "remove":
			val := readInt()
			res := obj.Remove(val)
			fmt.Fprintln(out, res)
		case "getRandom":
			res := obj.GetRandom()
			fmt.Fprintln(out, res)
		case "RandomizedSet":
			fmt.Fprintln(out, "null")
		}
	}
}
