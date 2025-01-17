// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

const MAXDEPTH = 3

type weblist struct {
	list  []string
	depth int
}

// !+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

// !+
func main() {
	worklist := make(chan weblist)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- weblist{os.Args[1:], 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > 3 {
			continue
		} else {
			for _, link := range list.list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string, depth int) {
						worklist <- weblist{crawl(link), depth + 1}
					}(link, list.depth)
				}
			}
		}

	}
}

//!-
