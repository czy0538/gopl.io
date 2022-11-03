// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"strings"
)

// !+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//!-

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	fmt.Println(max(1, 2, 3, 4))
	//!-slice

	fmt.Println(StringsJoin("/", "a"))
}

func max_5_15(vals ...int) int {
	if len(vals) == 0 {
		panic("Requires at least one parameter")
	}
	m := vals[0]

	for i := 1; i < len(vals); i++ {
		if m < vals[i] {
			m = vals[i]
		}
	}
	return m
}
func max(val int, vals ...int) int {
	for _, i := range vals {
		if i > val {
			val = i
		}
	}
	return val
}

func StringsJoin(seq, a string, s ...string) string {
	ss := make([]string, 0, len(s)+1)
	ss = append(ss, a)
	ss = append(ss, s...)
	return strings.Join(ss, seq)
}
