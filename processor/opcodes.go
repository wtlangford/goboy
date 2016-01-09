// vim: noet:ts=3:sw=3:sts=3
package processor

type Opcode struct {
	Opcode      byte
	ParamLen    uint8
	ShortCycles uint8
	LongCycles  uint8
	Func        func(*GBProcessor, byte, ...byte)
}
