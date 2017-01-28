// vim: noet:ts=3:sw=3:sts=3
package gb

import (
	"log"
)

type VolumeEnvelope struct {
	initialVolume uint8
	amplify       bool
	sweep         uint8
}

type LengthCounter struct {
	length        uint8
	lengthEnabled bool
}

type WaveChannel struct {
	LengthCounter

	enabled     bool
	outputLevel uint8
	frequency   uint16
	waveData    [16]byte
}

func (wc WaveChannel) readRegister(reg uint16) (val uint8) {
	switch reg {
	case 0:
		if wc.enabled {
			val = 1 << 7
		}
	case 1:
		val = wc.length
	case 2:
		val = (wc.outputLevel & 3) << 5
	case 4:
		if wc.lengthEnabled {
			val = 1 << 6
		}
	case 3:
		fallthrough // not readable
	default:
		log.Panicf("Invalid read of register %d from square channel", reg)
	}
	return
}

type NoiseChannel struct {
	VolumeEnvelope
	LengthCounter

	divisorCode  uint8
	divisorShift uint8
	widthMode    bool
}

func (nc NoiseChannel) readRegister(reg uint16) (val uint8) {
	switch reg {
	case 1:
		val = nc.length & 63
	case 2:
		val |= (nc.initialVolume & 15) << 4
		if nc.amplify {
			val |= 1 << 3
		}
		val |= nc.sweep & 7
	case 3:
		val |= (nc.divisorShift & 15) << 4
		if nc.widthMode {
			val |= 1 << 3
		}
		val |= nc.divisorCode & 7
	case 4:
		if nc.lengthEnabled {
			val = 1 << 6
		}
	default:
		log.Panicf("Invalid read of register %d from square channel", reg)
	}
	return
}

type SquareChannel struct {
	VolumeEnvelope
	LengthCounter

	duty      uint8
	frequency uint16
}

func (sc SquareChannel) readRegister(reg uint16) (val uint8) {
	switch reg {
	case 1:
		val = (sc.duty & 2) << 6
	case 2:
		val |= (sc.initialVolume & 15) << 4
		if sc.amplify {
			val |= 1 << 3
		}
		val |= sc.sweep & 7
	case 4:
		if sc.lengthEnabled {
			val = 1 << 6
		}
	case 3:
		fallthrough // not readable
	default:
		log.Panicf("Invalid read of register %d from square channel", reg)
	}
	return
}

type SquareSweepChannel struct {
	SquareChannel

	shadowFrq uint8

	sweepTime  uint8
	sweepMode  bool
	sweepShift uint8
}

func (ssc SquareSweepChannel) readRegister(reg uint16) (val uint8) {
	if reg == 0 {
		val |= (ssc.sweepTime & 7) << 4
		if ssc.sweepMode {
			val |= 1 << 3
		}
		val |= ssc.sweepShift & 7
		return
	} else {
		return ssc.SquareChannel.readRegister(reg)
	}
}

type Audio struct {
	wave        WaveChannel
	noise       NoiseChannel
	square      SquareChannel
	squareSweep SquareSweepChannel
}

func (apu *Audio) ReadAddress(addr uint16) uint8 {
	switch {
	case addr < 0xFF10:
		log.Panicf("Invalid read of address %d from audio unit", addr)
	case addr <= 0xFF14:
		// square sweep
		return apu.squareSweep.readRegister(addr - 0xFF10)
	case addr <= 0xFF19:
		// square
		return apu.square.readRegister(addr - 0xFF15)
	case addr <= 0xFF1E:
		// wave
		return apu.wave.readRegister(addr - 0xFF1A)
	case addr <= 0xFF23:
		// noise
		return apu.noise.readRegister(addr - 0xFF1F)
	case addr <= 0xFF26:
		// general stuff
		return 0
	default:
		log.Panicf("Invalid read of address %d from audio unit", addr)
	}
	panic("unreachable")
}

func (apu *Audio) WriteAddress(addr uint16, val byte) {
	log.Panicf("f")
}
