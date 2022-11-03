package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	var words []string
	for input.Scan() {
		words = append(words, input.Text())
	}
	wordfreq(words)
	fmt.Println("here")
}

func wordfreq(words []string) {
	fmt.Println(words)
	wordMap := make(map[string]uint)
	for _, s := range words {
		wordMap[s]++
	}

	for k, v := range wordMap {
		fmt.Printf("word %s : %d time(s)\n", k, v)
	}
}
