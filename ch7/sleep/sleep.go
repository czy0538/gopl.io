// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 179.

// The sleep program sleeps for a specified period of time.
package main

import (
	"flag"
	"fmt"
	"time"
)

type Weather struct {
	Weather string
}

func (w *Weather) Set(s string) error {
	if s == "" {
		return fmt.Errorf("empty weather")
	}
	w.Weather = s
	return nil
}

func (w *Weather) String() string {
	return "weather is " + w.Weather
}
func WeatherFlag(name string, value string, usage string) *Weather {
	f := Weather{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

// !+sleep
var period = flag.Duration("period", 1*time.Second, "sleep period")
var weather = WeatherFlag("weather", "", "set weather")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println(weather)
}

//!-sleep
