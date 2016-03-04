// vim: noet:ts=3:sw=3:sts=3
package cartridge

type Cartridge interface {
	ReadAddress(addr uint16) uint8
	WriteAddress(addr uint16, val uint8)
}

func NewCartridge(cartridgeROM []byte) Cartridge {
	// Inspect ROM to identify Cartridge Type
	// Instantiate that type
	switch cartridgeROM[0x147] {
	default:
		return nil
	}
}
