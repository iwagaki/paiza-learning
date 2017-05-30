package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	n := make([]int, count)

	for i := 0; i < count; i++ {
		scanner.Scan()
		n[i], _ = strconv.Atoi(scanner.Text())
	}

	sort.Sort(sort.IntSlice(n))

	for i := 0; i < count; i++ {
		fmt.Println(n[i])
	}
}
