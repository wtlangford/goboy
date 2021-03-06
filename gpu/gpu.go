// vim: noet:ts=3:sw=3:sts=3
package gpu

type GPU interface {
	ReadAddress(addr uint16) uint8
	WriteAddress(addr uint16, val byte)
	DMALoad(data []byte)
	Step(stepLength uint)

	MClocksToVBlank() uint
}
