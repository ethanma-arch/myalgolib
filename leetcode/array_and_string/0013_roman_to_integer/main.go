package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

var symbolValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func solve(s string) int {
	n := len(s)
	res := 0

	for i := 0; i < n; i++ {
		val := symbolValues[s[i]]
		if i < n-1 && val < symbolValues[s[i+1]] {
			res -= val
		} else {
			res += val
		}
	}

	return res
}

func main() {
	defer out.Flush()

	in.Split(bufio.ScanWords)

	for in.Scan() {
		s := in.Text()
		if len(s) == 0 {
			continue
		}

		res := solve(s)
		fmt.Fprintln(out, res)
	}
}
