// vim: noet:ts=3:sw=3:sts=3
package cartridge

import (
	"fmt"
	"os"
)

const (
	maxRomSize  = 2 * 1024 * 1024 // 16Mbit
	maxRamSize  = 32 * 1024       // 32KByte
	romBankSize = 16 * 1024       // 16KByte
	ramBankSize = 8 * 1024        // 8KByte
)

type Mbc1Cartridge struct {
	rom [maxRomSize]byte
	ram [maxRamSize]byte

	romBank []byte
	ramBank []byte

	addressMode bool
	ramEnabled  bool

	romSwitch byte
	ramSwitch byte
}

func NewMbc1Cartridge(cartridgeROM []byte) *Mbc1Cartridge {
	c := Mbc1Cartridge{}
	c.remapBanks()
	if len(cartridgeROM) > maxRomSize {
		panic(fmt.Sprintf("Trying to load too big ROM of size %d into MBC1 cartridge.", len(cartridgeROM)))
	}
	copy(c.rom[:], cartridgeROM)
	return &c
}

func (c *Mbc1Cartridge) ReadAddress(addr uint16) byte {
	switch {
	case addr < 0x4000:
		return c.rom[addr]
	case addr < 0x8000:
		return c.romBank[addr-0x4000]
	case addr >= 0xA000 && addr < 0xC000:
		if c.ramEnabled {
			return c.ramBank[addr-0xA000]
		} else {
			fmt.Fprintf(os.Stderr, "Reading from RAM address %d on MBC1 cartridge while RAM is disabled.", addr)
			return 0xff
		}
	}
	panic(fmt.Sprintf("Invalid read of address %d on MBC1 cartridge.", addr))
}

func (c *Mbc1Cartridge) WriteAddress(addr uint16, val uint8) {
	switch {
	case addr < 0x8000:
		c.writeControl(addr, val)
	case addr >= 0xA000 && addr < 0xC000:
		if c.ramEnabled {
			c.ramBank[addr-0xA000] = byte(val)
		} else {
			fmt.Fprintf(os.Stderr, "Writing to RAM address %d on MBC1 cartridge while RAM is disabled.", addr)
		}
	}
}

func (c *Mbc1Cartridge) writeControl(addr uint16, val uint8) {
	switch {
	case addr >= 0x6000:
		c.addressMode = (val & 1) == 1
		c.remapBanks()
	case addr >= 0x4000:
		c.ramSwitch = val
		c.remapBanks()
	case addr >= 0x2000:
		c.romSwitch = val
		c.remapBanks()
	default:
		c.ramEnabled = (val & 0xf) == 0xA
	}
}

func (c *Mbc1Cartridge) remapBanks() {
	romSelect := uint16(c.romSwitch & 0x1f) // Low 5 bits
	if romSelect == 0 {
		romSelect = 1
	}
	ramSelect := uint16(c.ramSwitch & 0x3) // Low 2 bits

	if c.addressMode == false { // 16/8
		romSelect |= ramSelect << 5
		ramSelect = 0
	}

	c.romBank = c.rom[romBankSize*romSelect : romBankSize*(romSelect+1)]
	c.ramBank = c.ram[ramBankSize*ramSelect : ramBankSize*(ramSelect+1)]
}