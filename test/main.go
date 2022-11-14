package main

import "fmt"

var m = make(map[string]int)

func main() {
	m["fuck"] = 1
	m["fuc"] = 2
	m["fu"] = 3

	for k, v := range m {
		fmt.Printf("k:%v v%v\n", k, v)
		v = 2
	}

	v := m["fuck"]
	v = 2
	fmt.Println(v)
	fmt.Println(m)
}
