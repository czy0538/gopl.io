package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type LineCounter int
type WordCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	*l = 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}
	return int(*l), scanner.Err()
}

func (w *WordCounter) Write(p []byte) (int, error) {
	*w = 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return int(*w), scanner.Err()
}

const text = `hello world
inject young girls`

func main() {
	var words WordCounter
	var lines LineCounter
	fmt.Println(fmt.Fprintf(&words, text))
	fmt.Println(fmt.Fprintf(&lines, text))
}
