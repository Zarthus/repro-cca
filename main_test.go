package main

import (
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	_, w, err := os.Pipe()
	_, err = hello(w)

	if err != nil {
		t.Fatalf("%v", err)
	}
}
