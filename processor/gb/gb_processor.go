// vim: noet:ts=3:sw=3:sts=3
package gb

import "github.com/wtlangford/goboy/bus"

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
	SlowStep   bool
	IME        bool
	MClock     uint
	TClock     uint
	Interrupts byte
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

	interrupts byte

	bus bus.Bus
}

func NewGBProcessor(bus bus.Bus) *GBProcessor {
	proc := &GBProcessor{bus: bus}
	proc.initOpcodes()
	return proc
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

	p.regs.PC += uint16(opcode.ParamLen)

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

func (p *GBProcessor) RunUntilVBlank() {
	endClock := p.state.MClock + p.bus.Gpu().MClocksToVBlank()
	for p.state.MClock < endClock {
		p.Step()
	}
}

func (p *GBProcessor) readAddress(addr uint16) uint8 {
	return p.bus.ReadAddress(addr)
}

func (p *GBProcessor) readAddress2(addr uint16) uint16 {
	// GB is little endian
	return uint16(p.readAddress(addr)) | uint16(p.readAddress(addr+1)<<8)
}

func (p *GBProcessor) writeAddress(addr uint16, val uint8) {
	p.bus.WriteAddress(addr, val)
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

func (p *GBProcessor) GetInterrupts() byte {
	return p.state.Interrupts
}

func (p *GBProcessor) SetInterrupts(val byte) {
	p.state.Interrupts = val
}
