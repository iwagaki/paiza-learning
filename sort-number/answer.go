package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	var n []int

	for i := 0; i < count; i++ {
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		n = append(n, m)
	}

	sort.Sort(sort.IntSlice(n))

	for i := 0; i < count; i++ {
		fmt.Println(n[i])
	}
}
