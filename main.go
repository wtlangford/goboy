package main

import (
	"io/ioutil"
	"log"
	"os"
)

// STUB file so that it links in all the sub packages
// so we can test that things build.

import "github.com/wtlangford/goboy/gb"

func main() {
	if len(os.Args) != 2 {
		log.Println("No cartridge file specified!")
		return
	}

	fileName := os.Args[1]
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("Cartridge file does not exist!")
		return
	}

	gb.NewBus(fileBytes)
}
