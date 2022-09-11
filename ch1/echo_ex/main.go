package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//ex1()
	ex2()
}

func ex1() {
	var s string
	for _, t := range os.Args {
		s += t + " "
	}
	fmt.Println(s)
}

func ex2() {
	fmt.Println(strings.Join(os.Args, "\n"))
}
