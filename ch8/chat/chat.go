// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const timeout = 10

// !+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	out := make(chan string) // outgoing client messages
	reset := make(chan struct{})
	tick := time.NewTicker(1 * time.Second)

	defer func() {
		conn.Close()
		tick.Stop()
	}()

	input := bufio.NewScanner(conn)

	go clientWriter(conn, out)

	go func() {
		counter := 0
		for {
			select {
			case <-tick.C:
				counter++
				if counter > timeout {
					conn.Close()
					return
				}
			case <-reset:
				counter = 0
			}
		}
	}()

	who := conn.RemoteAddr().String()
	if input.Scan() {
		who = input.Text()
	}
	out <- "You are " + who
	messages <- who + " has arrived\n"
	entering <- out

	for input.Scan() {
		reset <- struct{}{}
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- out
	messages <- who + " has left\n"

}

func clientWriter(conn net.Conn, out <-chan string) {
	for msg := range out {
		fmt.Fprintf(conn, msg)
	}
}

//!-handleConn

// !+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
