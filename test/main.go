package main

import "fmt"

func main() {
	stack := []int{}
	testStack(stack, 0)

	s := []int{1, 2}
	change(&s)
	fmt.Println(s)
}

func testStack(stack []int, n int) {
	fmt.Printf("n=%d,statck:%v\n", n, stack)
	if n == 3 {
		return
	}
	stack = append(stack, n)
	for i := 0; i < 3; i++ {
		testStack(stack, n+1)
	}
}

func change(stack *[]int) {
	*stack = append(*stack, 3)
	(*stack)[0] = -1
}
