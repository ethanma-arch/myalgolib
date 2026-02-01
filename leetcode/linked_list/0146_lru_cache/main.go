// LeetCode 146 LRU 缓存（带 TTL）
// 输入：第一行 "capacity [ttl_seconds]"（ttl 可选，0 或省略表示不过期），之后每行 "get key" 或 "put key value"
// 输出：每次 get 输出一行结果（命中为值，未命中或已过期为 -1）
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"leetcode/common_lib"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Node struct {
	key, value int
	prev, next *Node
	ExpireTime int64
}

type LRUCache struct {
	cap        int
	ttlSec     int64 // 全局 TTL（秒），0 表示不过期
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int, ttlSec int64) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:   capacity,
		ttlSec: ttlSec,
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (c *LRUCache) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) addToHead(n *Node) {
	n.next = c.head.next
	n.prev = c.head
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) moveToHead(n *Node) {
	c.remove(n)
	c.addToHead(n)
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.cache[key]
	if !ok {
		return -1
	}
	if n.ExpireTime > 0 && time.Now().Unix() > n.ExpireTime {
		c.remove(n)
		delete(c.cache, key)
		return -1
	}
	c.moveToHead(n)
	return n.value
}

func (c *LRUCache) Put(key, value int) {
	if c.cap <= 0 {
		return
	}
	now := time.Now().Unix()
	expire := int64(0)
	if c.ttlSec > 0 {
		expire = now + c.ttlSec
	}
	if n, ok := c.cache[key]; ok {
		n.value = value
		n.ExpireTime = expire
		c.moveToHead(n)
		return
	}
	if len(c.cache) >= c.cap {
		lru := c.tail.prev
		c.remove(lru)
		delete(c.cache, lru.key)
	}
	node := &Node{key: key, value: value, ExpireTime: expire}
	c.cache[key] = node
	c.addToHead(node)
}

func main() {
	defer out.Flush()

	line, err := common_lib.ReadNextValidLine(in)
	if err != nil && len(line) == 0 {
		return
	}
	parts0 := strings.Fields(line)
	capacity, _ := strconv.Atoi(parts0[0])
	ttlSec := int64(0)
	if len(parts0) >= 2 {
		ttlSec, _ = strconv.ParseInt(parts0[1], 10, 64)
	}
	cache := Constructor(capacity, ttlSec)

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
		case "get":
			if len(parts) < 2 {
				continue
			}
			k, _ := strconv.Atoi(parts[1])
			fmt.Fprintln(out, cache.Get(k))
		case "put":
			if len(parts) < 3 {
				continue
			}
			k, _ := strconv.Atoi(parts[1])
			v, _ := strconv.Atoi(parts[2])
			cache.Put(k, v)
		}
		if err != nil {
			break
		}
	}
}
