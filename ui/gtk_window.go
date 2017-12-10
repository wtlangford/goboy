package ui

import (
	"log"
	"runtime"
	"sync"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var doInit sync.Once
var colorValues = []byte{0xFF, 0xAA, 0x55, 0x00}

type GTKWindow struct {
	m    sync.Mutex
	app  *gtk.Application
	win  *gtk.ApplicationWindow
	pbuf *gdk.Pixbuf
}

func NewGTKWindow(width, height int) (*GTKWindow, error) {
	var err error
	w := &GTKWindow{}
	doInit.Do(initialize)
	w.app, err = gtk.ApplicationNew("com.github.wtlangford.goboy", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return nil, err
	}
	w.app.Connect("activate", func(app *gtk.Application) {
		log.Println("ACTIVATE")
		var err error
		//w.win, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
		w.win, err = gtk.ApplicationWindowNew(w.app)
		if err != nil {
			return
		}
		da, err := gtk.DrawingAreaNew()
		if err != nil {
			return
		}
		w.win.Add(da)
		w.win.SetTitle("Game")
		w.win.Connect("destroy", w.app.Quit)
		w.win.ShowAll()

		da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
			w.m.Lock()
			if w.pbuf != nil {
				width, height := w.win.GetSize()
				pb2, _ := w.pbuf.ScaleSimple(width, height, gdk.INTERP_NEAREST)
				gtk.GdkCairoSetSourcePixBuf(cr, pb2, 0, 0)
				cr.Rectangle(0, 0, float64(width), float64(height))
				cr.Fill()
			}
			w.m.Unlock()
		})

	})
	return w, nil
}

func (w *GTKWindow) DrawGrayscale(frameBuffer []byte, width, height int) {
	pbuf, err := gdk.PixbufNew(gdk.COLORSPACE_RGB, false, 8, width, height)
	if err != nil {
		log.Println("Pixbuf err:", err)
		return
	}
	frameBuffer[width*33+33] = 3
	frameBuffer[width*33+34] = 3
	frameBuffer[width*34+33] = 2
	frameBuffer[width*34+34] = 2
	numChans := pbuf.GetNChannels()
	rowStride := pbuf.GetRowstride()
	pixData := pbuf.GetPixels()
	var i int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b := frameBuffer[i]
			i++
			offset := y*rowStride + x*numChans
			pixData[offset] = colorValues[b]
			pixData[offset+1] = colorValues[b]
			pixData[offset+2] = colorValues[b]
		}
	}
	w.m.Lock()
	w.pbuf = pbuf
	w.m.Unlock()
	if w.win != nil {
		w.win.QueueDraw()
	} else {
		log.Println("skip")
	}
}

func (w *GTKWindow) Start() {
	w.app.Run(nil)
}

func (w *GTKWindow) Stop() {
	w.app.Quit()
}

func initialize() {
	runtime.LockOSThread()
	//gtk.Init(nil)
}
