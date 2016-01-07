package processor

type Processor interface {
	initOpcodes()
	readAddress(addr uint16, bytes int) []byte
	writeAddress(addr uint16, val uint8)
}
