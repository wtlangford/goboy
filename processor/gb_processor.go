// vim: noet:ts=3:sw=3:sts=3
package processor

const (
	GBFlagCarry     uint8 = 1 << 4
	GBFlagHalfCarry       = 1 << 5
	GBFlagSubtract        = 1 << 6
	GBFlagZero            = 1 << 7
	GBFlagMask            = 0xF0
)

type GBRegisters struct {
	A, F uint8
	B, C uint8
	D, E uint8
	H, L uint8
	SP   uint16
	PC   uint16
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
}

func (p *GBProcessor) ProcessOpcode() {

}

func (p *GBProcessor) writeAddress(addr uint16, val uint8) {
}

func (p *GBProcessor) readAddress(addr uint16) uint8 {
	return 0
}
