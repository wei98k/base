package main

import (
	"io"
	"os"
)

func main() {
	var w io.Writer
	// w.Write([]byte("hello")) // nil pointer dereference
	w = os.Stdout
	w.Write([]byte("hello"))
	// w = new(bytes.Buffer)
	// w = nil
}
