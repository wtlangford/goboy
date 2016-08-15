// vim: noet:ts=3:sw=3:sts=3
package processor

type Processor interface {
	Step() uint

	GetInterrupts() byte
	SetInterrupts(val byte)
}
