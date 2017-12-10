package ui

import (
	comp "golang.org/x/image/draw"
	"image"
	"image/color"
	"image/draw"
	"log"
	"sync"
	"time"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var uintColorValues = []uint8{0xFF, 0xAA, 0x55, 0x00}

type frameBuffer struct {
	data  []byte
	image image.Image
	size  image.Point
}

type ShinyWindow struct {
	m         sync.Mutex
	bnd       image.Rectangle
	win       screen.Window
	tex       screen.Texture
	buf       screen.Buffer
	fps       int
	lt        time.Time
	frameskip int
}

func NewShinyWindow(width, height int) (*ShinyWindow, error) {
	return &ShinyWindow{
		bnd: image.Rectangle{
			Min: image.ZP,
			Max: image.Point{X: width, Y: height},
		},
	}, nil
}

func (w *ShinyWindow) Start() {
	driver.Main(func(s screen.Screen) {
		var err error
		w.win, err = s.NewWindow(&screen.NewWindowOptions{Width: w.bnd.Max.X, Height: w.bnd.Max.Y})

		if err != nil {
			log.Println(err)
			return
		}
		w.tex, err = s.NewTexture(w.bnd.Max)
		if err != nil {
			log.Println(err)
			return
		}
		w.buf, err = s.NewBuffer(w.bnd.Max)
		if err != nil {
			log.Println(err)
			return
		}
		w.lt = time.Now()
		var fpsFunc func()
		fpsFunc = func() {
			w.m.Lock()
			now := time.Now()
			log.Println(":::::RPS:::::", float64(w.fps)/(float64(now.Sub(w.lt))/1000/1000/1000))
			log.Println(":::::FPS:::::", float64(w.frameskip)/(float64(now.Sub(w.lt))/1000/1000/1000))
			log.Println("bnd", w.bnd)
			w.lt = now
			w.fps = 0
			w.frameskip = 0
			w.m.Unlock()
			time.AfterFunc(time.Second, fpsFunc)
		}
		time.AfterFunc(time.Second, fpsFunc)
		for {
			switch e := w.win.NextEvent().(type) {
			case lifecycle.Event:
				log.Println("lc", e)
				if e.To == lifecycle.StageAlive {
					w.win.Fill(w.bnd, color.White, draw.Src)
					w.win.Publish()
				}

				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				log.Println("paint", e)
				w.win.Publish()
			case size.Event:
				log.Println("size", e)
				w.bnd = e.Bounds()
				log.Println("bnd", w.bnd)
				w.win.Scale(w.bnd, w.tex, w.tex.Bounds(), draw.Src, nil)
				w.win.Publish()
			case frameBuffer:
				log.Println("fb")
				w.m.Lock()
				w.fps++
				w.m.Unlock()
				w.drawFramebuffer(s, e)
				w.win.Scale(w.bnd, w.tex, w.tex.Bounds(), draw.Src, nil)
				w.win.Publish()
			case mouse.Event:
			case key.Event:
				log.Println("key", e)
			}
		}
	})
}

func (w *ShinyWindow) Stop() {
	// TODO: How?
}

func (w *ShinyWindow) DrawGrayscale(data []byte, width, height int) {
	w.frameskip++
	fb := frameBuffer{size: image.Point{X: width, Y: height}}
	fbImage := image.NewGray(image.Rectangle{Max: fb.size})
	fb.image = fbImage
	data[width*33+33] = 3
	data[width*33+34] = 3
	data[width*34+33] = 2
	data[width*34+34] = 2
	var i int
	for y := 0; y < fb.size.Y; y++ {
		for x := 0; x < fb.size.X; x++ {
			b := uintColorValues[data[i]]
			i++
			fbImage.Pix[y*fbImage.Stride+x] = b
		}
	}
	if w.win != nil {
		w.win.Send(fb)
	}
}

func (w *ShinyWindow) drawFramebuffer(s screen.Screen, fb frameBuffer) error {
	m := w.buf.RGBA()
	comp.NearestNeighbor.Scale(m, m.Bounds(), fb.image, fb.image.Bounds(), comp.Src, nil)
	w.tex.Upload(image.ZP, w.buf, w.buf.Bounds())
	return nil
}
