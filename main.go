package main

// STUB file so that it links in all the sub packages
// so we can test that things build.

import bus "github.com/wtlangford/goboy/bus/gb"

func main() {
	bus.NewGBBus([]byte{})
}
