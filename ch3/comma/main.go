// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", myComma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func myComma(s string) string {
	start := 0
	if s[0] == '+' || s[0] == '-' {
		start = 1
	}

	dot := strings.IndexByte(s, '.')
	if dot == -1 {
		return s[0:start] + myCommaInner(s[start:])
	} else {
		return s[0:start] + myCommaInner(s[start:dot]) + "." + myCommaInner(s[dot+1:])
	}

}

func myCommaInner(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	i := len(s) % 3
	buf.WriteString(s[0:i])
	for ; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

//!-
