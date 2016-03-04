// vim: noet:ts=3:sw=3:sts=3
package cartridge

import "fmt"

const cartridgeSize int = 0x8000

type RomOnlyCartridge struct {
	rom [cartridgeSize]byte
}

func NewRomOnlyCartridge(cartridgeROM []byte) *RomOnlyCartridge {
	if len(cartridgeROM) > cartridgeSize {
		panic(fmt.Sprintf("Trying to load too big ROM of size %d into ROM Only Cartridge", len(cartridgeROM)))
	}
	c := RomOnlyCartridge{}
	copy(c.rom[:], cartridgeROM)
	return &c
}

func (c *RomOnlyCartridge) ReadAddress(addr uint8) uint8 {
	if int(addr) > len(c.rom) {
		panic(fmt.Sprintf("Invalid read at address %d on ROM Only Cartridge", addr))
	}
	return c.rom[addr]
}

func (c *RomOnlyCartridge) WriteAddress(addr uint8, val uint8) {
	panic(fmt.Sprintf("Invalid write to address %d on ROM Only Cartridge", addr))
}
