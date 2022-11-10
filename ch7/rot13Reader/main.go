package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b = 'A' + (b-'A'+13)%26
	} else if b >= 'a' && b <= 'z' {
		b = 'a' + (b-'a'+13)%26
	}
	return b
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	l, err := r.r.Read(p)
	for i := 0; i < l; i++ {
		p[i] = rot13(p[i])
	}
	return l, err // 一定要返回EOF，不然会无限的刷新
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
