package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// STUB file so that it links in all the sub packages
// so we can test that things build.

import bus "github.com/wtlangford/goboy/bus/gb"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "No cartridge file specified!")
		return
	}

	fileName := os.Args[1]
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cartridge file does not exist!")
		return
	}

	bus.NewGBBus(fileBytes)
}
