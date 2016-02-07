// vim: noet:ts=3:sw=3:sts=3
package cartridge

import "fmt"

const cartridgeSize int = 0x8000

type RomOnlyCartridge struct {
	rom [cartridgeSize]byte
}

func NewRomOnlyCartridge(cartridgeROM []byte) *RomOnlyCartridge {
	if len(cartridgeROM) > cartridgeSize {
		panic(fmt.Sprintf("CartridgeROM is too big: is %d bytes", len(cartridgeROM)))
	}
	c := RomOnlyCartridge{}
	copy(c.rom[:], cartridgeROM)
	return &c
}

func (c *RomOnlyCartridge) ReadAddress(addr uint8) byte {
	if addr < 0 || int(addr) > len(c.rom) {
		panic(fmt.Sprintf("Address %d is out of bounds", addr))
	}
	return c.rom[addr]
}

func (c *RomOnlyCartridge) WriteAddress(addr uint8, val uint8) {
	panic(fmt.Sprintf("Illegal write in ROM Only Cartridge to address %d", addr))
}
