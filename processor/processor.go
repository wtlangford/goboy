// vim: noet:ts=3:sw=3:sts=3
package processor

type Processor interface {
	Step()

	GetInterrupts() byte
	SetInterrupts(val byte)
}
