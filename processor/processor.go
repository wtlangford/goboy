// vim: noet:ts=3:sw=3:sts=3
package processor

type Processor interface {
	initOpcodes()
	readAddress(addr uint16, bytes int) []byte
	writeAddress(addr uint16, val uint8)
	writeAddress2(addr uint16, val uint16)
	pushStack(val uint16)
	popStack() uint16
	Step()
}
