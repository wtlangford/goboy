// vim: noet:ts=3:sw=3:sts=3
package audio

type Audio interface {
	ReadAddress(addr uint16) uint8
	WriteAddress(addr uint16, val byte)
}

type Dummy struct{}

func (d *Dummy) ReadAddress(addr uint16) uint8 {
	return 0
}

func (d *Dummy) WriteAddress(addr uint16, val byte) {
	return
}
