package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, j := range tropsort(prereqs) {
		fmt.Printf("%d:%s\n", i, j)
	}
}

func tropsort(m map[string][]string) []string {
	var visited = make(map[string]bool)
	var courses []string
	var order []string
	for key := range prereqs {
		courses = append(courses, key)
	}
	sort.Strings(courses)

	var visit func([]string)
	visit = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				visit(m[item])
				order = append(order, item)
			}

		}
	}

	visit(courses)

	return order
}
