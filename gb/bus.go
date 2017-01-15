// vim: noet:ts=3:sw=3:sts=3
package gb

import (
	"github.com/wtlangford/goboy/cartridge"
	gpu "github.com/wtlangford/goboy/gpu"
	processor "github.com/wtlangford/goboy/processor"
)

type Bus struct {
	gpu       *Gpu
	cartridge cartridge.Cartridge
	processor *Processor

	internalRAM [0x2000]byte // 8kB
	highRAM     [127]byte
}

func NewBus(cartridgeData []byte) *Bus {
	bus := &Bus{}
	bus.gpu = NewGpu(bus)
	bus.cartridge = cartridge.NewCartridge(cartridgeData)
	bus.processor = NewProcessor(bus)
	return bus
}

func (b *Bus) ReadAddress(addr uint16) uint8 {
	switch {
	case addr < 0x8000: // ROM BANK 0 | Switchable ROM bank
		return b.cartridge.ReadAddress(addr)
	case addr < 0xA000: // Video RAM
		return b.gpu.ReadAddress(addr)
	case addr < 0xC000: // Switchable RAM
		return b.cartridge.ReadAddress(addr)
	case addr < 0xE000: // Internal RAM
		return uint8(b.internalRAM[0xE000-addr])
	case addr < 0xFE00: // Shadow RAM
		return uint8(b.internalRAM[0xFE00-addr])
	case addr < 0xFEA0: // Sprite Attribute Memory
		return b.gpu.ReadAddress(addr)
	case addr < 0xFF00: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFF10:
		// Timer, joypad, serial data, interrupt flag IO Registers
	case addr < 0xFF40:
		// Sound IO Registers
	case addr < 0xFF4C: // GPU Registers
		return b.gpu.ReadAddress(addr)
	case addr < 0xFF80: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFFFF: // High RAM
		return uint8(b.highRAM[0xFFFE-addr])
	case addr == 0xFFFF: // Interrupt Control Register
		return b.processor.GetInterrupts()
	}
	return 0
}

func (b *Bus) WriteAddress(addr uint16, val uint8) {
	switch {
	case addr < 0x8000: // Cartridge ROM - Control
		b.cartridge.WriteAddress(addr, val)
	case addr < 0xA000: // Video RAM
		b.gpu.WriteAddress(addr, val)
	case addr < 0xC000: // Cartridge RAM
		b.cartridge.WriteAddress(addr, val)
	case addr < 0xE000: // Internal Ram
		b.internalRAM[0xE000-addr] = byte(val)
	case addr < 0xFE00: // Shadow RAM
		b.internalRAM[0xFE00-addr] = byte(val)
	case addr < 0xFEA0: // Sprite Attribute Memory
		b.gpu.WriteAddress(addr, val)
	case addr < 0xFF00: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFF10:
		// Timer, joypad, serial data, interrupt flag IO Registers
	case addr < 0xFF40:
		// Sound IO Registers
	case addr < 0xFF4C: // GPU Registers
		b.gpu.WriteAddress(addr, val)
	case addr < 0xFF80: // Unused
		// LOG THIS.  Probably an error.
	case addr < 0xFFFF: // High RAM
		b.highRAM[0xFFFE-addr] = byte(val)
	case addr == 0xFFFF: // Interrupt Control Register
		b.processor.SetInterrupts(byte(val))
	}
}

func (b *Bus) Gpu() gpu.GPU {
	return b.gpu
}

func (b *Bus) Processor() processor.Processor {
	return b.processor
}

func (b *Bus) Cartridge() cartridge.Cartridge {
	return b.cartridge
}
