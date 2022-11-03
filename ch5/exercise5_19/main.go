package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(add(1, 2))
}

func add(a, b int) (result int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Fprintf(os.Stderr, "panic value is %v\n", p)
			result = p.(int)
		}
	}()
	panic(a + b)
}
