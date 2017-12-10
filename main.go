package main

import (
	"io/ioutil"
	"log"
	//"math/rand"
	//"time"
	"flag"

	"github.com/wtlangford/goboy/ui"
)

// STUB file so that it links in all the sub packages
// so we can test that things build.

import "github.com/wtlangford/goboy/gb"

func main() {
	useShiny := flag.Bool("shiny", false, "use shiny")
	cartridgeFile := flag.String("cartridge", "", "cartridge to play")
	flag.Parse()
	//log.SetFlags(0)
	if *cartridgeFile == "" {
		log.Println("No cartridge file specified!")
		return
	}

	fileBytes, err := ioutil.ReadFile(*cartridgeFile)
	if err != nil {
		log.Println("Cartridge file does not exist!")
		return
	}

	var win ui.Window
	if *useShiny {
		win, err = ui.NewShinyWindow(960, 864)
	} else {
		win, err = ui.NewGTKWindow(960, 864)
	}
	if err != nil {
		log.Panicln(err)
	}

	bus := gb.NewBus(fileBytes, win)
	go bus.Run()
	/*
		go func() {
			gscale := make([]byte, 4, 4)
			for {
				gscale[0] = byte(rand.Intn(4))
				gscale[1] = byte(rand.Intn(4))
				gscale[2] = byte(rand.Intn(4))
				gscale[3] = byte(rand.Intn(4))
				win.DrawGrayscale(gscale, 2, 2)
				time.Sleep(1000 * time.Millisecond)
			}
			//win.Stop()
		}()*/
	win.Start()

}
