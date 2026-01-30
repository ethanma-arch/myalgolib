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

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func solve(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

func main() {
	defer out.Flush()

	in.Split(bufio.ScanWords)

	for in.Scan() {
		token := in.Text()

		if len(token) == 0 {
			continue
		}

		num, err := strconv.Atoi(token)
		if err != nil {
			continue
		}

		res := solve(num)
		fmt.Fprintln(out, res)
	}
}
