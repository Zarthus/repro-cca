package main

import (
	"fmt"
	"os"
)

func main() {
	hello(os.Stdout)
}

func hello(f *os.File) (int, error) {
	return fmt.Fprintf(f, "Hello!")
}
