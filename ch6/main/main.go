package main

import (
	"fmt"
	"gopl.io/ch6/intset"
)

func main() {
	var a intset.IntSet
	a.Add(1)
	a.AddAll(1, 2, 3)
	fmt.Println(&a)
	a.AddAll(-1)
	fmt.Println(&a)
}
