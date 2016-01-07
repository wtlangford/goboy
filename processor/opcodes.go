package processor

type Opcode struct {
	Opcode      byte
	ParamLen    uint8
	ShortCycles uint8
	LongCycles  uint8
	Func        func(*GBProcessor, byte, ...byte)
}
