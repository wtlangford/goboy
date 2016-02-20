// vim: noet:ts=3:sw=3:sts=3
package processor

const (
	GBFlagCarry     uint8 = 1 << 4
	GBFlagHalfCarry uint8 = 1 << 5
	GBFlagSubtract  uint8 = 1 << 6
	GBFlagZero      uint8 = 1 << 7
	GBFlagMask      uint8 = 0xF0
)

type GBRegisters struct {
	A, F uint8
	B, C uint8
	D, E uint8
	H, L uint8
	SP   uint16
	PC   uint16
}

type GBState struct {
	SlowStep bool
	IME      bool
	MClock   uint
	TClock   uint
}

func (r *GBRegisters) AF() uint16 {
	return (uint16(r.A) << 8) | uint16(r.F)
}

func (r *GBRegisters) SetAF(af uint16) {
	r.A = uint8(af >> 8)
	r.F = uint8(af)
}

func (r *GBRegisters) BC() uint16 {
	return (uint16(r.B) << 8) | uint16(r.C)
}

func (r *GBRegisters) SetBC(bc uint16) {
	r.B = uint8(bc >> 8)
	r.C = uint8(bc)
}

func (r *GBRegisters) DE() uint16 {
	return (uint16(r.D) << 8) | uint16(r.E)
}

func (r *GBRegisters) SetDE(de uint16) {
	r.D = uint8(de >> 8)
	r.E = uint8(de)
}

func (r *GBRegisters) HL() uint16 {
	return (uint16(r.H) << 8) | uint16(r.L)
}

func (r *GBRegisters) SetHL(hl uint16) {
	r.H = uint8(hl >> 8)
	r.L = uint8(hl)
}

func (r *GBRegisters) GetFlags() uint8 {
	return r.F & GBFlagMask
}

type GBProcessor struct {
	opcodes   []Opcode
	cbOpcodes []Opcode
	regs      GBRegisters
	state     GBState

	internalRAM [0x2000]byte // 8kB
	highRAM     [127]byte
	interrupts  byte
}

func (p *GBProcessor) Step() {
	var opcode Opcode
	op := p.readAddress(p.regs.PC)
	oplen := 1
	if op == 0xCB {
		op = p.readAddress(p.regs.PC + 1)
		opcode = p.cbOpcodes[op]
		oplen += 1
	} else {
		opcode = p.opcodes[op]
	}
	plen := int(opcode.ParamLen) - oplen

	params := make([]byte, plen)

	for i := 0; i < plen; i += 1 {
		params[i] = p.readAddress(p.regs.PC + uint16(oplen+i))
	}

	opcode.Func(p, opcode.Opcode, params...)

	if p.state.SlowStep {
		p.state.MClock += uint(opcode.LongCycles)
		p.state.TClock += 4 * uint(opcode.LongCycles)
	} else {
		p.state.MClock += uint(opcode.ShortCycles)
		p.state.TClock += 4 * uint(opcode.ShortCycles)
	}
	p.state.SlowStep = false

}

func (p *GBProcessor) readAddress(addr uint16) uint8 {
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

func (p *GBProcessor) readAddress2(addr uint16) uint16 {
	// GB is little endian
	return uint16(p.readAddress(addr)) | uint16(p.readAddress(addr+1)<<8)
}

func (p *GBProcessor) writeAddress(addr uint16, val uint8) {
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

func (p *GBProcessor) writeAddress2(addr uint16, val uint16) {
	// GB is little endian.
	p.writeAddress(addr, uint8(val&0xFF))
	p.writeAddress(addr+1, uint8(val>>8))
}

func (p *GBProcessor) pushStack(val uint16) {
	p.regs.SP -= 2
	p.writeAddress2(p.regs.SP, val)
}

func (p *GBProcessor) popStack() uint16 {
	val := p.readAddress2(p.regs.SP)
	p.regs.SP += 2
	return val
}
