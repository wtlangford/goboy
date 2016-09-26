// vim: noet:ts=3:sw=3:sts=3
package gb

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

	scanline        byte
	scanlineCompare byte
	mode            gpuMode
	modeClock       uint

	registerMap map[uint16]*byte

	bus bus.Bus
}

type gpuMode uint

const (
	gpuModeHorizontalBlank gpuMode = iota
	gpuModeVerticalBlank
	gpuModeScanlineOAM
	gpuModeScanlineVRAM
)

const (
	lcdcBackgroundEnabled byte = 1 << iota
	lcdcColor0Tranparency
	lcdcSpriteSize
	lcdcBackgroundTileTableAddress
	lcdcTilePatternTableAddress
	lcdcWindowEnabled
	lcdcWindowTileTableAddress
	lcdcLCDEnabled
)

const (
	_ byte = 1 << iota
	_
	scanlineCoincidenceFlag
	interruptOnControllerMode00
	interruptOnControllerMode01
	interruptOnControllerMode10
	interruptOnCoincidence
)

type Screen struct {
	ScrollY, ScrollX byte
}
type Window struct {
	WindowY, WindowX byte
}

func NewGBGpu(bus bus.Bus) *GBGpu {
	gpu := GBGpu{bus: bus}

	gpu.registerMap = map[uint16]*byte{
		0xFF40: &gpu.LCDC,
		0xFF41: &gpu.Stat,
		0xFF42: &gpu.ScrollY,
		0xFF43: &gpu.ScrollX,
		0xFF44: &gpu.scanline,
		0xFF45: &gpu.scanlineCompare,
		//FF46: DMA is write only
		0xFF47: &gpu.BackgroundPalette,
		0xFF48: &gpu.ObjectPalette0,
		0xFF49: &gpu.ObjectPalette1,
		0xFF4A: &gpu.WindowY,
		0xFF4B: &gpu.WindowX,
	}

	return &gpu
}

func (g *GBGpu) ReadAddress(addr uint16) byte {
	switch {
	case addr >= 0x8000 && addr < 0xA000:
		return g.vram[addr-0x8000]
	case addr > 0xFE00 && addr < 0xFEA0:
		return g.oam[addr-0xFE00]
	case addr >= 0xFF40:
		reg := g.registerMap[addr]
		if reg != nil {
			return *reg
		} else {
			panic(fmt.Sprintf("Invalid read of address %d from GPU", addr))
		}
	}
	panic(fmt.Sprintf("Invalid read of address %d from GPU", addr))
}

func (g *GBGpu) readVRAM(addr uint16, length uint16) []byte {
	if addr < 0x8000 || (addr+length) > 0xA000 {
		panic(fmt.Sprintf("Invalid read of length %d at address %d from GPU", length, addr))
	}
	addr -= 0x8000
	return g.vram[addr : addr+length]
}

func (g *GBGpu) WriteAddress(addr uint16, val byte) {
	switch {
	case addr >= 0x8000 && addr < 0xA000:
		g.vram[addr-0x800] = val
	case addr > 0xFE00 && addr < 0xFEA0:
		g.oam[addr-0xFE00] = val
	case addr >= 0xFF40:
		if addr == 0xFF41 {
			val = (val &^ 0x3) | (g.Stat & 0x3)
		} else if addr == 0xFF44 {
			val = 0
		} else if addr == 0xFF46 {
			// uint16(val) << 8
			// DMA stuff
			return
		}
		reg := g.registerMap[addr]
		if reg != nil {
			*reg = val
		}
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

func (g *GBGpu) renderScan() {
	/*
	 *	Relevant addresses:
	 *		tile map 0 => 0x9800 - 0x9BFF
	 *		tile map 1 => 0x9C00 - 0x9FFF
	 *		tile set 0 => 0x8000 - 0x8FFF
	 *		tile set 1 => 0x8800 - 0x97FF
	 *		background palette => 0xFF47
	 */
	if g.LCDC&lcdcLCDEnabled > 0 {

		if g.LCDC&lcdcBackgroundEnabled > 0 {
			renderWindow := g.LCDC&lcdcWindowEnabled > 0
			g.renderBackground(renderWindow)
		}

		g.renderSprites()
	}
}

// Draws one line of the Backgound to the frame buffer
// Optionally draws one line of the window.
func (g *GBGpu) renderBackground(renderWindow bool) {
	// Figure out where in background to draw
	// Get pixel data
	// Map to proper color
	// Draw to position on buffer
	// Optionally repeat for window
	xPos := int(g.ScrollX)
	yPos := int(g.scanline + g.ScrollY)

	var tileMapAddress uint16
	if g.LCDC&lcdcBackgroundTileTableAddress == 0 {
		tileMapAddress = 0x9800
	} else {
		tileMapAddress = 0x9C00
	}

	var tileDataAddress int
	if g.LCDC&lcdcTilePatternTableAddress == 0 {
		tileDataAddress = 0x8000
	} else {
		tileDataAddress = 0x8800
	}

	// Tiles are 8 lines high, but we have to skip whole tiles.
	// integer division yields number of whole tiles to skip (vertically)
	// 32 tiles per row
	tileMapAddress += uint16(yPos) / 8 * 32

	y := yPos % 8
	for tileIndex := 0; tileIndex < 32; tileIndex++ {
		x := xPos % 8
		tile := g.ReadAddress(tileMapAddress + uint16(xPos)/8)
		tileData := g.readVRAM(uint16(tileDataAddress+int(int8(tile))), 16)
		tileColor := g.readPixel(tileData, x, y)
		// TODO:
		// Grab tile shade from tilePatternPalette
		// Put shade into frame buffer
		// Finish writing loop to actually be right.
		// (Increment x properly, handle wrapping of background.)
		// Also handle wrapping vertically (should be done outside of loop.  If y > 255 ...)

	}

}

// Read one pixel value from the tile data
// (0, 0) is the top-left pixel
// Tile data maps like this:
// 11110000 11001100 -> 3 3 2 2 1 1 0 0
func (g *GBGpu) readPixel(tileData []byte, x, y int) byte {
	if x < 0 || x > 7 || y < 0 || y > 7 {
		panic(fmt.Sprintf("Invalid pixel address in tile: (%i, %i)", x, y))
	}
	tileData = tileData[y*2 : (y*2)+2]
	return ((tileData[1] >> uint(6-x)) & 2) | ((tileData[0] >> uint(7-x)) & 1)
}

func (g *GBGpu) renderSprites() {}

func (g *GBGpu) Step(stepLength uint) {

	g.modeClock += stepLength

	switch g.mode {
	case gpuModeScanlineOAM:
		if g.modeClock >= 20 {
			g.modeClock -= 20
			g.mode = gpuModeScanlineVRAM
		}

	case gpuModeScanlineVRAM:
		if g.modeClock >= 43 {
			g.modeClock -= 43
			g.mode = gpuModeHorizontalBlank
			g.renderScan()
		}

	case gpuModeHorizontalBlank:
		if g.modeClock >= 51 {
			g.modeClock -= 51
			g.mode = gpuModeScanlineOAM
			g.scanline++
			if g.scanline == 143 {
				g.mode = gpuModeVerticalBlank
				// TODO: render frame?
			}
		}

	case gpuModeVerticalBlank:
		if g.modeClock >= 114 {
			g.modeClock -= 114
			g.scanline++
			if g.scanline == 154 {
				g.scanline = 0
				g.mode = gpuModeScanlineOAM
			}
		}
	}
}

func (g *GBGpu) MClocksToVBlank() uint {
	return 17556
}
