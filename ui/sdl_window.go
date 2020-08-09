package ui

import (
	"log"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLWindow struct {
	win *sdl.Window
	rdr *sdl.Renderer

	fullscreen bool
}

func NewSDLWindow(width, height int) (*SDLWindow, error) {
	var (
		win SDLWindow
		err error
	)
	sdl.Do(func() {
		win.win, err = sdl.CreateWindow("Goboy", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(width), int32(height), sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
		if err != nil {
			return
		}
		win.rdr, err = sdl.CreateRenderer(win.win, -1, sdl.RENDERER_PRESENTVSYNC)
		if err != nil {
			return
		}
		win.rdr.SetIntegerScale(true)
	})
	if err != nil {
		return nil, err
	}
	return &win, nil
}

func (w *SDLWindow) Start() {
	// sdl.Do(w.win.Show)
}

func (w *SDLWindow) Stop() {
	w.win.Destroy()
}

func hexColorToColor(hex uint32) sdl.Color {
	return sdl.Color{
		R: uint8((hex >> 16) & 0xFF),
		G: uint8((hex >> 8) & 0xFF),
		B: uint8(hex & 0xFF),
	}
}

func (w *SDLWindow) DrawGrayscale(frameBuffer []byte, width, height int) {
	buf := make([]byte, len(frameBuffer))
	copy(buf, frameBuffer)
	sdl.Do(func() {
		sur, err := sdl.CreateRGBSurfaceWithFormatFrom(unsafe.Pointer(&buf[0]), int32(width), int32(height), 8, int32(width), sdl.PIXELFORMAT_INDEX8)
		if err != nil {
			log.Panicln(err)
		}
		defer sur.Free()
		// TODO: Configurable palette!!
		sur.Format.Palette.SetColors([]sdl.Color{
			hexColorToColor(0x9bbc0f),
			hexColorToColor(0x8bac0f),
			hexColorToColor(0x306230),
			hexColorToColor(0x0f380f),
			// {R: 0xFF, G: 0xFF, B: 0xFF},
			// {R: 0xAA, G: 0xAA, B: 0xAA},
			// {R: 0x55, G: 0x55, B: 0x55},
			// {R: 0x00, G: 0x00, B: 0x00},
		})
		tex, err := w.rdr.CreateTextureFromSurface(sur)
		if err != nil {
			log.Panicln(err)
		}
		w.rdr.Clear()
		err = w.rdr.Copy(tex, nil, w.targetRect(sur.W, sur.H))
		if err != nil {
			log.Panicln(err)
		}
		w.rdr.Present()
	})
}

func (w *SDLWindow) targetRect(imgW, imgH int32) *sdl.Rect {
	scrW, scrH := w.win.GetSize()

	wRatio := float64(scrW) / float64(imgW)
	hRatio := float64(scrH) / float64(imgH)

	var rect sdl.Rect
	if wRatio < hRatio {
		rect.W = scrW
		rect.H = imgH * scrW / imgW
		rect.Y = (scrH - rect.H) / 2
	} else {
		rect.W = imgW * scrH / imgH
		rect.H = scrH
		rect.X = (scrW - rect.W) / 2
	}

	return &rect
}

func (w *SDLWindow) Info() {
	log.Println(w.win.GetSize())
	log.Println(w.rdr.GetLogicalSize())
}

func (w *SDLWindow) ToggleFullscreen() {
	if w.fullscreen {
		w.win.SetFullscreen(0)
	} else {
		w.win.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	}
	w.fullscreen = !w.fullscreen
}
