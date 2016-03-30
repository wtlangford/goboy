// vim: noet:ts=3:sw=3:sts=3
package cartridge

type Cartridge interface {
	ReadAddress(addr uint16) uint8
	WriteAddress(addr uint16, val uint8)
}

func NewCartridge(cartridgeROM []byte) Cartridge {
	// Cartridge ROM header length is 0x14F
	if len(cartridgeROM) < 0x14F {
		panic("Cartridge ROM is too short. Cannot contain full header.")
	}
	switch cartridgeROM[0x147] {
	case 0x0:
		return NewRomOnlyCartridge(cartridgeROM)
	case 0x1, 0x2:
		return NewMbc1Cartridge(cartridgeROM)
	// refer to page 11 of manual for other cartridge types
	default:
		panic("Invalid cartridge type.")
	}
}
