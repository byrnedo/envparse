package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	VERSION = "1.0"
	B_EMPTY = []byte{}
)

func main() {

	var file = ""

	if len(os.Args) < 2 {
		os.Exit(1)
	} else if (os.Args[1]) == "--help" {
		fmt.Println("Usage: \n\t./envParse FILE_TO_PARSE")
		os.Exit(1)
	}

	file = os.Args[1]

	handle, err := os.Open(file)
	if err != nil {
		os.Exit(2)
	}

	defer handle.Close()

	buf := bufio.NewReader(handle)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(os.ExpandEnv(string(line)))
	}
}
