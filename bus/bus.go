// vim: noet:ts=3:sw=3:sts=3
package bus

import (
	"github.com/wtlangford/goboy/cartridge"
	"github.com/wtlangford/goboy/gpu"
	"github.com/wtlangford/goboy/processor"
)

type Bus interface {
	ReadAddress(addr uint16) uint8
	WriteAddress(addr uint16, val uint8)
	Gpu() gpu.GPU
	Processor() processor.Processor
	Cartridge() cartridge.Cartridge
}
