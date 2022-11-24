// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type Info struct {
	c     net.Conn
	shout string
	delay time.Duration
}

// !+
func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	done := sync.WaitGroup{}
	info := Info{c, "", 1 * time.Second}
	resetCount := make(chan struct{})
	ticks := time.NewTicker(1 * time.Second)
	defer ticks.Stop()

	echo := func(info Info) {
		fmt.Fprintln(info.c, "\t", strings.ToUpper(info.shout))
		time.Sleep(info.delay)
		fmt.Fprintln(info.c, "\t", info.shout)
		time.Sleep(info.delay)
		fmt.Fprintln(info.c, "\t", strings.ToLower(info.shout))
		done.Done()
	}

	go func() {
		for input.Scan() {
			resetCount <- struct{}{}
			done.Add(1)
			info.shout = input.Text()
			go echo(info)
		}
	}()
	for counter := 0; counter < 5; counter++ {
		select {
		case <-resetCount:
			counter = 0
		case <-ticks.C:
			println(counter)

		}
	}
	done.Wait()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
