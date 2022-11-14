package main

import (
	"flag"
	"fmt"
	"strconv"
)

type Port string

func (p *Port) Set(s string) error {
	if s == "" {
		return fmt.Errorf("empty port number")
	}

	port, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if port < 0 || port > 65535 {
		return fmt.Errorf("illegal port number %d", port)
	}

	*p = Port(s)
	return nil
}

func (p *Port) String() string {
	return "port is " + string(*p)
}
func PortFlag(name string, value string, usage string) *Port {
	f := Port(value)
	flag.CommandLine.Var(&f, name, usage)
	return &f
}
