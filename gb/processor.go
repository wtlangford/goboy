// vim: noet:ts=3:sw=3:sts=3
package gb

import (
	//"fmt"

	"github.com/wtlangford/goboy/bus"
)

const (
	GBFlagCarry     uint8 = 1 << 4
	GBFlagHalfCarry uint8 = 1 << 5
	GBFlagSubtract  uint8 = 1 << 6
	GBFlagZero      uint8 = 1 << 7
	GBFlagMask      uint8 = 0xF0
)

type Registers struct {
	A, F uint8
	B, C uint8
	D, E uint8
	H, L uint8
	SP   uint16
	PC   uint16
}

type State struct {
	SlowStep   bool
	IME        bool
	MClock     uint
	TClock     uint
	Interrupts byte
}

func (r *Registers) AF() uint16 {
	return (uint16(r.A) << 8) | uint16(r.F)
}

func (r *Registers) SetAF(af uint16) {
	r.A = uint8(af >> 8)
	r.F = uint8(af)
}

func (r *Registers) BC() uint16 {
	return (uint16(r.B) << 8) | uint16(r.C)
}

func (r *Registers) SetBC(bc uint16) {
	r.B = uint8(bc >> 8)
	r.C = uint8(bc)
}

func (r *Registers) DE() uint16 {
	return (uint16(r.D) << 8) | uint16(r.E)
}

func (r *Registers) SetDE(de uint16) {
	r.D = uint8(de >> 8)
	r.E = uint8(de)
}

func (r *Registers) HL() uint16 {
	return (uint16(r.H) << 8) | uint16(r.L)
}

func (r *Registers) SetHL(hl uint16) {
	r.H = uint8(hl >> 8)
	r.L = uint8(hl)
}

func (r *Registers) GetFlags() uint8 {
	return r.F & GBFlagMask
}

type Processor struct {
	opcodes   []Opcode
	cbOpcodes []Opcode
	regs      Registers
	state     State

	interrupts byte

	bus bus.Bus
}

func NewProcessor(bus bus.Bus) *Processor {
	proc := &Processor{bus: bus}
	proc.initOpcodes()
	proc.regs = Registers{
		A:  0x01,
		C:  0x13,
		E:  0xD8,
		F:  0xB0,
		H:  0x01,
		L:  0x4D,
		SP: 0xFFFE,
		PC: 0x0100,
	}
	proc.regs.PC = 0x0100
	return proc
}

func (p *Processor) Step() uint {
	var opcode Opcode
	op := p.readAddress(p.regs.PC)
	oplen := 1
	//var cb bool
	if op == 0xCB {
		op = p.readAddress(p.regs.PC + 1)
		opcode = p.cbOpcodes[op]
		oplen += 1
		//	cb = true
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
		p.state.SlowStep = false
		p.state.MClock += uint(opcode.LongCycles)
		p.state.TClock += 4 * uint(opcode.LongCycles)
		return uint(opcode.LongCycles)
	} else {
		p.state.MClock += uint(opcode.ShortCycles)
		p.state.TClock += 4 * uint(opcode.ShortCycles)
		return uint(opcode.ShortCycles)
	}

}

func (p *Processor) RunUntilVBlank() {
	totalSteps := p.bus.Gpu().MClocksToVBlank()
	for stepsTaken := uint(0); stepsTaken < totalSteps; {
		stepSize := p.Step()
		p.bus.Gpu().Step(stepSize)
		stepsTaken += stepSize
	}
}

func (p *Processor) readAddress(addr uint16) uint8 {
	return p.bus.ReadAddress(addr)
}

func (p *Processor) readAddress2(addr uint16) uint16 {
	// GB is little endian
	return uint16(p.readAddress(addr)) | uint16(p.readAddress(addr+1))<<8
}

func (p *Processor) writeAddress(addr uint16, val uint8) {
	p.bus.WriteAddress(addr, val)
}

func (p *Processor) writeAddress2(addr uint16, val uint16) {
	// GB is little endian.
	p.writeAddress(addr, uint8(val&0xFF))
	p.writeAddress(addr+1, uint8(val>>8))
}

func (p *Processor) pushStack(val uint16) {
	p.regs.SP -= 2
	p.writeAddress2(p.regs.SP, val)
}

func (p *Processor) popStack() uint16 {
	val := p.readAddress2(p.regs.SP)
	p.regs.SP += 2
	return val
}

func (p *Processor) GetInterrupts() byte {
	return p.state.Interrupts
}

func (p *Processor) SetInterrupts(val byte) {
	p.state.Interrupts = val
}
