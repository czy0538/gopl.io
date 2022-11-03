package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := []string{"", "foo", "foofoo", "fooaaafooaaafoo"}
	for i, t := range s {
		fmt.Printf("%d:%s\n", i, expand(t, Reverse))
	}

}

func expand(s string, f func(string) string) string {
	var buf bytes.Buffer

	for l, r := 0, 0; r < len(s); {
		r = strings.Index(s[r:], "foo")
		if r == -1 {
			buf.WriteString(s[l:])
			break
		}
		r += l //对子串进行的索引，因此应该加上偏移量
		buf.WriteString(s[l:r])
		buf.WriteString(f("foo"))
		r += 3
		l = r

	}
	return buf.String()
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
