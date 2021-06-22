// Package main
package main

import (
	"fmt"

	ut "github.com/x0ray/vscodetest/stringutil"
)

const (
	pgm     = "vscodetest"
	ver     = "0.0.1"
	verdate = "21June2021"
)

var FileName string = "hello.txt"

// main entry to program
func main() {
	fmt.Printf("Command: %s version: %s of %s, started.\n", pgm, ver, verdate)

	revstr, err := ut.ReverseFile(FileName)
	if err != nil {
		fmt.Printf("Command: %s error: %v\n", pgm, err)
	} else {
		fmt.Printf("Command: %s reversed file: %s, content follows...\n%s\n", pgm, FileName, revstr)
	}

	fmt.Printf("Command: %s ended.\n", pgm)
}
