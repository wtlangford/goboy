package ui

type Window interface {
	Start()
	Stop()
	DrawGrayscale(frameBuffer []byte, width, height int)
}
