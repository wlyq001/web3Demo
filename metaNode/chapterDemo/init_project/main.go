//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(os.Args)
}
