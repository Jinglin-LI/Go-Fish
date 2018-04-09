package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logWriter struct{}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// or just use io.Copy(os.Stdout, file)
	lw := logWriter{}
	io.Copy(lw, file)

	// fmt.Println(os.Args[1])
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
