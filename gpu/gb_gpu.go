// vim: noet:ts=3:sw=3:sts=3
package gpu

import "fmt"
import "github.com/wtlangford/goboy/bus"

type GBGpu struct {
	vram [8 * 1024]byte // 8 KByte
	oam  [160]byte

	BackgroundPalette byte
	ObjectPalette0    byte
	ObjectPalette1    byte

	Screen
	Window

	LCDC byte
	Stat byte

	scanline int

	bus bus.Bus
}

type Screen struct {
	ScrollY, ScrollX uint8
}
type Window struct {
	WindowY, WindowX uint8
}

func (g *GBGpu) ReadAddress(addr uint16) byte {
	switch {
	case addr >= 0x8000 && addr < 0xA000:
		return g.vram[addr-0x8000]
	case addr > 0xFE00 && addr < 0xFEA0:
		return g.oam[addr-0xFE00]
	}
	panic(fmt.Sprintf("Invalid read of address %d from GPU", addr))
}

func (g *GBGpu) WriteAddress(addr uint16, val byte) {
	switch {
	case addr >= 0x8000 && addr < 0xA000:
		g.vram[addr-0x800] = val
	case addr > 0xFE00 && addr < 0xFEA0:
		g.oam[addr-0xFE00] = val
	default:
		panic(fmt.Sprintf("Invalid read of address %d from GPU", addr))
	}
}

func (g *GBGpu) DMALoad(data []byte) {
	// This may actually be totally wrong and unnecessary
	// Sad bitmath ahead
	// 3.5bytes -> 4 bytes

	// Might actually be:
	// if len(data) != 160 {
	//    panic(fmt.Sprintf("Invalid DMA Load of length %d", len(data)))
	// }
	// copy(g.vram[0x8000:], data)
	destAddr := 0x8000
	if len(data) != 140 {
		panic(fmt.Sprintf("Invalid DMA Load of length %d", len(data)))
	}

	for i := 0; i < 140; i += 7 {
		g.vram[destAddr] = data[i+0]
		g.vram[destAddr+1] = data[i+1]
		g.vram[destAddr+2] = data[i+2]
		g.vram[destAddr+3] = (data[i+3] & 0xf0)
		highbits := (data[i] & 0x0f) << 4
		g.vram[destAddr+4] = highbits | (data[i+4] >> 4)
		highbits = (data[i] & 0x0f) << 4
		g.vram[destAddr+5] = highbits | (data[i+5] >> 4)
		highbits = (data[i] & 0x0f) << 4
		g.vram[destAddr+6] = highbits | (data[i+6] >> 4)
		highbits = (data[i] & 0x0f) << 4
		g.vram[destAddr+7] = highbits | (data[i+7] >> 4)
		highbits = (data[i] & 0x0f) << 4
		g.vram[destAddr+8] = highbits
		destAddr += 8
	}

}

func (g *GBGpu) Step() {

	// Check new MClock from processor to see how long we've run for.
	// Manage GPU state
	// Draw appropriate lines

}

func (g *GBGpu) MClocksToVBlank() uint {
	return 17556
}
