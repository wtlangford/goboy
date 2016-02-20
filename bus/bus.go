// vim: noet:ts=3:sw=3:sts=3
package bus

import (
	"github.com/wtlangford/goboy/cartridge"
	"github.com/wtlangford/goboy/gpu"
	"github.com/wtlangford/goboy/processor"
)

type Bus struct {
	gpu       GPU
	cartridge Cartridge
	processor Processor

	internalRAM [0x2000]byte // 8kB
	highRAM     [127]byte
}

func (b *Bus) readAddress(addr uint16) uint8 {
	switch {
	case addr < 0x8000:
		// FORWARD TO CARTRIDGE
		// ROM BANK 0 | Switchable ROM bank
	case addr < 0xA000: // Video RAM
		// FORWARD TO GPU
		// Video RAM
	case addr < 0xC000: // Switchable RAM
		// FORWARD TO CARTRIDGE
		// Switchable RAM bank
	case addr < 0xE000: // Internal RAM
		return uint8(p.internalRAM[0xE000-addr])
	case addr < 0xFE00: // Shadow RAM
		return uint8(p.internalRAM[0xFE00-addr])
	case addr < 0xFEA0: // Sprite Attribute Memory
		// FORWARD TO GPU
		// Sprite Attribute Memory
	case addr < 0xFF00: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFF4C: // IO Registers
		// FORWARD TO IO
	case addr < 0xFF80: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFFFF: // High RAM
		return uint8(p.highRAM[0xFFFE-addr])
	case addr == 0xFFFF: // Interrupt Control Register
		return uint8(p.interrupts)
	}
	return 0
}

func (b *Bus) writeAddress(addr uint16) uint8 {
	switch {
	case addr < 0x8000: // Cartridge ROM - Control
		// FORWARD TO CARTRIDGE
		// This is cartridge bank control
	case addr < 0xA000: // Video RAM
		// FORWARD TO GPU
		// Video RAM
	case addr < 0xC000: // Cartridge RAM
		// FORWARD TO CARTRIDGE
		// Switchable RAM bank
	case addr < 0xE000: // Internal Ram
		p.internalRAM[0xE000-addr] = byte(val)
	case addr < 0xFE00: // Shadow RAM
		p.internalRAM[0xFE00-addr] = byte(val)
	case addr < 0xFEA0: // Sprite Attribute Memory
		// FORWARD TO GPU
		// Sprite Attribute Memory
	case addr < 0xFF00: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFF4C: // IO Registers
		// FORWARD TO IO
	case addr < 0xFF80: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFFFF: // High RAM
		p.highRAM[0xFFFE-addr] = byte(val)
	case addr == 0xFFFF: // Interrupt Control Register
		p.interrupts = byte(val)
	}
}