// vim: noet:ts=3:sw=3:sts=3
package bus

import (
	"github.com/wtlangford/goboy/cartridge"
	"github.com/wtlangford/goboy/gpu"
	"github.com/wtlangford/goboy/processor"
)

type GBBus struct {
	Gpu       gpu.GPU
	Cartridge cartridge.Cartridge
	Processor processor.Processor

	internalRAM [0x2000]byte // 8kB
	highRAM     [127]byte
}

func (b *GBBus) ReadAddress(addr uint16) uint8 {
	switch {
	case addr < 0x8000: // ROM BANK 0 | Switchable ROM bank
		// FORWARD TO CARTRIDGE
		return b.Cartridge.ReadAddress(addr)
	case addr < 0xA000: // Video RAM
		// FORWARD TO GPU
		// Video RAM
		return b.Gpu.ReadAddress(addr)
	case addr < 0xC000: // Switchable RAM
		// FORWARD TO CARTRIDGE
		// Switchable RAM bank
		return b.Cartridge.ReadAddress(addr)
	case addr < 0xE000: // Internal RAM
		return uint8(b.internalRAM[0xE000-addr])
	case addr < 0xFE00: // Shadow RAM
		return uint8(b.internalRAM[0xFE00-addr])
	case addr < 0xFEA0: // Sprite Attribute Memory
		// FORWARD TO GPU
		// Sprite Attribute Memory
		return b.Gpu.ReadAddress(addr)
	case addr < 0xFF00: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFF4C: // IO Registers
		// FORWARD TO IO
	case addr < 0xFF80: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFFFF: // High RAM
		return uint8(b.highRAM[0xFFFE-addr])
	case addr == 0xFFFF: // Interrupt Control Register
		return b.Processor.GetInterrupts()
	}
	return 0
}

func (b *GBBus) WriteAddress(addr uint16, val uint8) {
	switch {
	case addr < 0x8000: // Cartridge ROM - Control
		// FORWARD TO CARTRIDGE
		// This is cartridge bank control
		b.Cartridge.WriteAddress(addr, val)
	case addr < 0xA000: // Video RAM
		// FORWARD TO GPU
		// Video RAM
		b.Gpu.WriteAddress(addr, val)
	case addr < 0xC000: // Cartridge RAM
		// FORWARD TO CARTRIDGE
		// Switchable RAM bank
		b.Cartridge.WriteAddress(addr, val)
	case addr < 0xE000: // Internal Ram
		b.internalRAM[0xE000-addr] = byte(val)
	case addr < 0xFE00: // Shadow RAM
		b.internalRAM[0xFE00-addr] = byte(val)
	case addr < 0xFEA0: // Sprite Attribute Memory
		// FORWARD TO GPU
		// Sprite Attribute Memory
		b.Gpu.WriteAddress(addr, val)
	case addr < 0xFF00: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFF4C: // IO Registers
		// FORWARD TO IO
	case addr < 0xFF80: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFFFF: // High RAM
		b.highRAM[0xFFFE-addr] = byte(val)
	case addr == 0xFFFF: // Interrupt Control Register
		b.Processor.SetInterrupts(byte(val))
	}
}
