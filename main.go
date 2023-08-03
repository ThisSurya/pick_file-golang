package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var filePath = os.Args

	file, err := os.Open(filePath[1])
	if err != nil {
		fmt.Println("error expected ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}