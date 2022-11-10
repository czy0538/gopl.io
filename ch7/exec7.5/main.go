package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type LimitedReader struct {
	r io.Reader
	n int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[0:l.n]
	}
	n, err = l.r.Read(p)
	// 如果不减会被反复的调用
	l.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	scanner := bufio.NewScanner(LimitReader(os.Stdin, 3))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//scanner := bufio.NewScanner(io.LimitReader(os.Stdin, 3))
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
}
