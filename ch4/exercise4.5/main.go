package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []string

func main() {
	input := bufio.NewScanner(os.Stdin) //读取一行，并删除行尾部的换行符
	input.Scan()                        // 读取，读到一行时返回true，否则返回false
	s := input.Text()                   // 获取读取的结果
	strings := make([]string, 0, len(s))
	for _, i := range s {
		strings = append(strings, string(i))
	}
	s2 := make([]string, len(strings))
	copy(s2, strings)
	fmt.Println(unique(s2))
	fmt.Println(leetcode_unqiue(strings))
	fmt.Println(string_unqiue(s))
}

func unique(strings []string) []string {
	i := 0
	for _, s := range strings {
		if strings[i] == s {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}

func leetcode_unqiue(strings []string) []string {
	temp := make(stack, 0, len(strings))
	temp = append(temp, strings[0])
	for _, s := range strings[1:] {
		if len(temp) != 0 && s == temp[len(temp)-1] {
			temp = temp[:len(temp)-1]
		} else {
			temp = append(temp, s)
		}
	}
	return temp
}
func string_unqiue(s string) string {
	bs := []byte(s)
	temp := make([]byte, 0, len(bs))
	temp = append(temp, bs[0])
	for _, s := range bs[1:] {
		if len(temp) != 0 && s == temp[len(temp)-1] {
			temp = temp[:len(temp)-1]
		} else {
			temp = append(temp, s)
		}
	}
	return string(temp)
}
