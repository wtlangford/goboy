// vim: noet:ts=3:sw=3:sts=3
package cartridge

import "log"

type Cartridge interface {
	ReadAddress(addr uint16) uint8
	WriteAddress(addr uint16, val uint8)
}

func NewCartridge(cartridgeROM []byte) Cartridge {
	// Cartridge ROM header length is 0x14F
	if len(cartridgeROM) < 0x14F {
		log.Panicln("Cartridge ROM is too short and does not contain a full header.")
	}
	switch cartridgeROM[0x147] {
	case 0x0:
		return NewRomOnlyCartridge(cartridgeROM)
	case 0x1, 0x2:
		return NewMbc1Cartridge(cartridgeROM)
	// refer to page 11 of manual for other cartridge types
	default:
		log.Panicf("Invalid cartridge type %x.\n", cartridgeROM[0x147])
		panic("unreachable")
	}
}
