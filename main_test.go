// Package main
package main

import "testing"

func Test_main(t *testing.T) {
	t.Run("Test main()", func(t *testing.T) {
		main()
	})
}

func Test_mainBadFile(t *testing.T) {
	FileName = "FileName.bad"
	t.Run("Test main()", func(t *testing.T) {
		main()
	})
}
