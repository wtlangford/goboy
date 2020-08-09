package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	_ "github.com/veandco/go-sdl2/gfx"
	_ "github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/wtlangford/goboy/gb"
	"github.com/wtlangford/goboy/ui"
)

var (
	cartridgeFile = flag.String("cartridge", "", "cartridge to play")
)

func main() {
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

	sdlMain(fileBytes)
}

func realMain(fileBytes []byte, win ui.Window) {
	_ = gb.NewBus(fileBytes, win)
	go func() {
		w, h := 10, 10
		gscale := make([]byte, w*h, w*h)
		for {
			for i := 0; i < w*h; i++ {
				gscale[i] = byte(rand.Intn(4))
			}
			win.DrawGrayscale(gscale, w, h)
			time.Sleep(500 * time.Millisecond)
		}
		//win.Stop()
	}()
	win.Start()
}

func sdlMain(fileBytes []byte) {
	var win ui.Window
	var err error
	if err = sdl.Init(sdl.INIT_AUDIO | sdl.INIT_VIDEO | sdl.INIT_EVENTS); err != nil {
		log.Panicln(err)
	}
	defer sdl.Quit()
	defer log.Println("shutting down")

	sdl.Main(func() {
		log.Println("in SDL main")
		win, err = ui.NewSDLWindow(480, 432)
		if err != nil {
			return
		}
		w, h := 160, 144
		gscale := make([]byte, w*h, w*h)

		running := true
		go func() {
			for running {
				for i := 0; i < w*h; i++ {
					gscale[i] = byte(rand.Intn(4))
				}
				win.DrawGrayscale(gscale, w, h)
				time.Sleep(17 * time.Millisecond)
			}
		}()

		for running {
			sdl.Do(func() {
				switch e := sdl.PollEvent().(type) {
				case *sdl.QuitEvent:
					running = false
				case *sdl.KeyboardEvent:
					if e.Repeat != 0 {
						break
					}
					if e.Keysym.Sym == sdl.K_RETURN && e.State == sdl.PRESSED {
						win.(*ui.SDLWindow).ToggleFullscreen()
					} else if e.Keysym.Sym == sdl.K_ESCAPE && e.State == sdl.PRESSED {
						win.(*ui.SDLWindow).Info()
					}
				}
			})
		}
	})
	if err != nil {
		log.Panicln(err)
	}
}
