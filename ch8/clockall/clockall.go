package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type RemoteTime struct {
	time    string
	isAlive bool
	city    string
}

var locationServerMap = make(map[string]*RemoteTime)

func main() {
	err := handleArg(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	go printClock()

	// try to keep all server connected
	for {
		for k, v := range locationServerMap {
			if v.isAlive == false {
				conn, err := net.Dial("tcp", k)
				if err != nil {
					fmt.Printf("connect to %s failed, err is:%v\n", k, err)
					continue
				}
				v.isAlive = true
				go handleConn(conn)
			}

		}
		time.Sleep(5 * time.Second)
	}

}

func handleConn(c net.Conn) {
	// close Conn and set server status
	defer func() {
		val, ok := locationServerMap[c.RemoteAddr().String()]
		if ok {
			val.isAlive = false
		}
		c.Close()
	}()
	input := bufio.NewScanner(c)
	for {
		if input.Scan() {
			val, ok := locationServerMap[c.RemoteAddr().String()]
			if ok {
				val.time = input.Text()
			}
		}
		time.Sleep(1 * time.Second)
	}

}

func handleArg(s []string) error {
	if len(s) == 0 {
		return fmt.Errorf("no arguments")
	}
	// clean map
	locationServerMap = make(map[string]*RemoteTime)

	for _, info := range s {
		// find = ,which will be used to spilt city and address
		equalPos := strings.Index(info, "=")
		if equalPos == -1 {
			return fmt.Errorf("failed to handle arguments")
		}
		locationServerMap[info[equalPos+1:]] = &RemoteTime{
			time:    "",
			city:    info[0:equalPos],
			isAlive: false,
		}
	}
	return nil
}

func printClock() {
	for locationServerMap != nil {
		for _, v := range locationServerMap {
			if v.isAlive {
				fmt.Printf("%s is now:%s\n", v.city, v.time)
			}
		}
		time.Sleep(1 * time.Second)
	}
	log.Fatal("print clock error")
}
