package main

import (
	"flag"
	"fmt"
	"time"
)

type Location struct {
	Location *time.Location
}

func (l *Location) Set(s string) error {
	location, err := time.LoadLocation(s)
	if err != nil {
		return err
	}
	l.Location = location
	return nil
}

func (l *Location) String() string {
	return fmt.Sprintf("current time location is %v", l.Location)
}

func LocationFlag() *Location {
	f := Location{time.Local}
	flag.CommandLine.Var(&f, "location", "set time location")
	return &f
}
