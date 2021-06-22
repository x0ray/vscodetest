/*
   This is a demo file extracted from the demo by Suzy Mueller
   at: https://youtu.be/6r08zGi38Tk
*/

// Package stringutil string utility functions
package stringutil

import "os"

// Reverse returns the argument string reversed
func Reverse(s string) string {
	// The following code for the reverse function was borrowed from the Go standard library
	r := []rune(s)
	// fixed a bug in the following: len(r); --> len(r)-1
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		// fixed bug in the following: r[i], r[i] = r[j], r[j] --> r[i], r[j] = r[j], r[i]
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseFile returns the contents of the file specified in the first argument in reverse order
func ReverseFile(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	// I used intellisense: ctrl+" " ife <enter>
	// for the following...
	if err != nil {
		return "", err
	}
	// I used intellisense: "re"<tab>, "R"<tab>, "con"<tab>
	// , ctrl+" "
	// hover over detected error in return, select quick fix: ctrl+. <enter>
	return Reverse(string(contents)), err
}

// used: ctrl+shft+p go:gen  to generate go test file for all functions
// can also right click and select Go: ...
