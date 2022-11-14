package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List struct {
	next *List
	val  interface{}
}

func PrintList(head *List) {
	if head != nil {
		fmt.Println(head.val)
		PrintList(head.next)
	}
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	head := new(List)
	head.val = 0
	head.next = new(List)
	head.next.val = 1
	PrintList(head)
}
