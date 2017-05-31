package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	table := make([]bool, n)

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	count := 0

	for i := 0; i < m; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		is_all := true

		for j := 0; j < a; j++ {
			index := (b - 1 + j) % n
			if table[index] {
				is_all = false
				break
			}
		}

		if is_all {
			for j := 0; j < a; j++ {
				index := (b - 1 + j) % n
				table[index] = true
			}
			count += a
		}
	}

	fmt.Println(count)
}
