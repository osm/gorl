package main

import (
	"fmt"
	"os"
)

func exitf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}
