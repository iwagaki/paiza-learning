package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	count := map[string]int{}
	word := []string{}

	for scanner.Scan() {
		key := scanner.Text()
		_, haskey := count[key]
		if haskey {
			count[key]++
		} else {
			count[key] = 1
			word = append(word, key)
		}
	}

	for _, key := range word {
		fmt.Printf("%s %d\n", key, count[key])
	}
}
