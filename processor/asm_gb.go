// vim: noet:ts=3:sw=3:sts=3
package processor

// This file's structure was auto-generated
// Mnemonic: NOP
// Sets Flags: ----
func (p *GBProcessor) nop(opcode byte, params ...byte) {
	// Intentionally empty
}

// Mnemonic: LD rr,d16
//  rr := BC,DE,HL,SP
// Sets Flags: ----
func (p *GBProcessor) ld_rri(opcode byte, params ...byte) {
	val := uint16(params[0]) | (uint16(params[1]) << 8)
	switch opcode {
	case 0x01:
		p.regs.SetBC(val)
	case 0x11:
		p.regs.SetDE(val)
	case 0x21:
		p.regs.SetHL(val)
	case 0x31:
		p.regs.SP = val
	}
}

// Mnemonic: LD (rr),A
// rr := BC,DE,HL+,HL-
// Sets Flags: ----
func (p *GBProcessor) ld_xrra(opcode byte, params ...byte) {
	switch opcode {
	case 0x02:
		p.regs.A = p.readAddress(p.regs.BC())
	case 0x12:
		p.regs.A = p.readAddress(p.regs.DE())
	case 0x22:
		hl := p.regs.HL()
		p.regs.A = p.readAddress(hl)
		p.regs.SetHL(hl + 1)
	case 0x32:
		hl := p.regs.HL()
		p.regs.A = p.readAddress(hl)
		p.regs.SetHL(hl - 1)
	}
}

// Mnemonic: INC rr
//  rr := BC,DE,HL,SP
// Sets Flags: ----
func (p *GBProcessor) inc_rr(opcode byte, params ...byte) {
	switch opcode {
	case 0x03:
		p.regs.SetBC(p.regs.BC() + 1)
	case 0x13:
		p.regs.SetDE(p.regs.DE() + 1)
	case 0x23:
		p.regs.SetHL(p.regs.HL() + 1)
	case 0x33:
		p.regs.SP++
	}
}

// Mnemonic: INC r
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z0h-
func (p *GBProcessor) inc_r(opcode byte, params ...byte) {
	var val *uint8
	var f uint8

	switch opcode {
	case 0x3c:
		val = &p.regs.A
	case 0x04:
		val = &p.regs.B
	case 0x0c:
		val = &p.regs.C
	case 0x14:
		val = &p.regs.D
	case 0x1c:
		val = &p.regs.E
	case 0x24:
		val = &p.regs.H
	case 0x2c:
		val = &p.regs.L
	case 0x34:
		v := p.readAddress(p.regs.HL())
		val = &v
	}

	// Since +1, if bottom nibble is 1111, we will carry
	if (*val & 0xF) == 0xF {
		f |= GBFlagHalfCarry
	}

	*val++

	if *val == 0 {
		f |= GBFlagZero
	}

	f |= p.regs.F & GBFlagCarry
	p.regs.F = f

	if opcode == 0x34 {
		p.writeAddress(p.regs.HL(), *val)
	}
}

// Mnemonic: DEC r
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z1h-
func (p *GBProcessor) dec_r(opcode byte, params ...byte) {
	var val *uint8
	var f uint8 = GBFlagSubtract

	switch opcode {
	case 0x3c:
		val = &p.regs.A
	case 0x04:
		val = &p.regs.B
	case 0x0c:
		val = &p.regs.C
	case 0x14:
		val = &p.regs.D
	case 0x1c:
		val = &p.regs.E
	case 0x24:
		val = &p.regs.H
	case 0x2c:
		val = &p.regs.L
	case 0x35:
		v := p.readAddress(p.regs.HL())
		val = &v
	}

	// Since -1 if bottom nibble == 0000, we will carry
	if (*val & 0xF) == 0 {
		f |= GBFlagHalfCarry
	}

	*val--

	if *val == 0 {
		f |= GBFlagZero
	}

	f |= p.regs.F & GBFlagCarry
	p.regs.F = f

	if opcode == 0x35 {
		p.writeAddress(p.regs.HL(), *val)
	}
}

// Mnemonic: LD r,d8
//  r := A,B,C,D,E,H,L
// Sets Flags: ----
func (p *GBProcessor) ld_ri(opcode byte, params ...byte) {
	val := uint8(params[0])

	switch opcode {
	case 0x06:
		p.regs.B = val
	case 0x0e:
		p.regs.C = val
	case 0x16:
		p.regs.D = val
	case 0x1e:
		p.regs.E = val
	case 0x26:
		p.regs.H = val
	case 0x2e:
		p.regs.L = val
	case 0x3e:
		p.regs.A = val
	}
}

// Mnemonic: RLCA
// Sets Flags: z00c
func (p *GBProcessor) rlca(opcode byte, params ...byte) {

	if p.regs.A&0x80 > 0 {
		p.regs.A = (p.regs.A << 1) + 1
		p.regs.F = GBFlagCarry
	} else {
		p.regs.A = p.regs.A << 1
		p.regs.F = 0
	}

	if p.regs.A == 0 {
		p.regs.F |= GBFlagZero
	}
}

// Mnemonic: LD (a16),SP
// Sets Flags: ----
func (p *GBProcessor) ld_xisp(opcode byte, params ...byte) {
	addr := uint16(params[1])<<8 + uint16(params[0])
	p.writeAddress2(addr, p.regs.SP)

}

// Mnemonic: ADD HL,BC
// Sets Flags: -0hc
func (p *GBProcessor) add_hlbc(opcode byte, params ...byte) {

}

// Mnemonic: LD A,(BC)
// Sets Flags: ----
func (p *GBProcessor) ld_axrr(opcode byte, params ...byte) {
	switch opcode {
	case 0x0a:
		p.regs.A = p.readAddress(p.regs.BC())
	case 0x1a:
		p.regs.A = p.readAddress(p.regs.DE())
	}
}

// Mnemonic: DEC rr
//  rr := BC,DE,HL,SP
// Sets Flags: ----
func (p *GBProcessor) dec_rr(opcode byte, params ...byte) {
	switch opcode {
	case 0x0B:
		p.regs.SetBC(p.regs.BC() - 1)
	case 0x1B:
		p.regs.SetDE(p.regs.DE() - 1)
	case 0x2B:
		p.regs.SetHL(p.regs.HL() - 1)
	case 0x3B:
		p.regs.SP--
	}
}

// Mnemonic: RRCA
// Sets Flags: 000c
func (p *GBProcessor) rrca(opcode byte, params ...byte) {
	if p.regs.A&0x01 > 0 {
		p.regs.A = (p.regs.A >> 1) | 0x80
		p.regs.F = GBFlagCarry
	} else {
		p.regs.A = p.regs.A >> 1
		p.regs.F = 0
	}

	if p.regs.A == 0 {
		p.regs.F |= GBFlagZero
	}
}

// Mnemonic: STOP 0
// Sets Flags: ----
func (p *GBProcessor) stop_0(opcode byte, params ...byte) {

}

// Mnemonic: RLA
// Sets Flags: 000c
func (p *GBProcessor) rla(opcode byte, params ...byte) {
	var low uint8

	if p.regs.F&GBFlagCarry > 0 {
		low = 0x01
	}

	if p.regs.A&0x80 > 0 {
		p.regs.F = GBFlagCarry
	} else {
		p.regs.F = 0
	}

	p.regs.A = (p.regs.A << 1) | low
}

// Mnemonic: JR r8
// Sets Flags: ----
func (p *GBProcessor) jr_i(opcode byte, params ...byte) {

}

// Mnemonic: ADD HL,DE
// Sets Flags: -0hc
func (p *GBProcessor) add_hlde(opcode byte, params ...byte) {

}

// Mnemonic: RRA
// Sets Flags: 000c
func (p *GBProcessor) rra(opcode byte, params ...byte) {
	var high uint8

	if p.regs.F&GBFlagCarry > 0 {
		high = 0x80
	}

	if p.regs.A&0x01 > 0 {
		p.regs.F = GBFlagCarry
	} else {
		p.regs.F = 0
	}

	p.regs.A = (p.regs.A >> 1) | high
}

// Mnemonic: JR NZ,r8
// Sets Flags: ----
func (p *GBProcessor) jr_nzi(opcode byte, params ...byte) {

}

// Mnemonic: DAA
// Sets Flags: z-0c
func (p *GBProcessor) daa(opcode byte, params ...byte) {
	val := uint16(p.regs.A)
	f := p.regs.F
	sub := (f&GBFlagSubtract == 0)
	carry := (f&GBFlagCarry > 0)
	halfc := (f&GBFlagHalfCarry > 0)
	f &= GBFlagSubtract

	if sub {
		if halfc || val&0xF > 9 {
			val += 0x06
		}
		if carry || val > 0x9F {
			val += 0x60
		}
	} else {
		if halfc {
			val -= 0x06
			if !carry {
				val &= 0xFF
			}
		}
		if carry {
			val -= 0x60
		}
	}
	if val > 0xFF {
		f |= GBFlagCarry
	}
	p.regs.A = uint8(val & 0xFF)
	if p.regs.A == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f
}

// Mnemonic: JR Z,r8
// Sets Flags: ----
func (p *GBProcessor) jr_zi(opcode byte, params ...byte) {

}

// Mnemonic: ADD HL,HL
// Sets Flags: -0hc
func (p *GBProcessor) add_hlhl(opcode byte, params ...byte) {

}

// Mnemonic: LD A,(HL+)
// Sets Flags: ----
func (p *GBProcessor) ld_axhli(opcode byte, params ...byte) {
	p.regs.A = p.readAddress(p.regs.HL())
	p.regs.SetHL(p.regs.HL() + 1)
}

// Mnemonic: CPL
// Sets Flags: -11-
func (p *GBProcessor) cpl(opcode byte, params ...byte) {
	p.regs.A = ^p.regs.A
	p.regs.F |= GBFlagSubtract | GBFlagHalfCarry

}

// Mnemonic: JR NC,r8
// Sets Flags: ----
func (p *GBProcessor) jr_nci(opcode byte, params ...byte) {

}

// Mnemonic: LD (HL),d8
// Sets Flags: ----
func (p *GBProcessor) ld_xhli(opcode byte, params ...byte) {
	p.writeAddress(p.regs.HL(), uint8(params[0]))
}

// Mnemonic: SCF
// Sets Flags: -001
func (p *GBProcessor) scf(opcode byte, params ...byte) {
	p.regs.F |= GBFlagCarry
	p.regs.F &= ^(GBFlagSubtract | GBFlagHalfCarry)
}

// Mnemonic: JR C,r8
// Sets Flags: ----
func (p *GBProcessor) jr_ci(opcode byte, params ...byte) {

}

// Mnemonic: ADD HL,SP
// Sets Flags: -0hc
func (p *GBProcessor) add_hlsp(opcode byte, params ...byte) {

}

// Mnemonic: LD A,(HL-)
// Sets Flags: ----
func (p *GBProcessor) ld_axhld(opcode byte, params ...byte) {
	p.regs.A = p.readAddress(p.regs.HL())
	p.regs.SetHL(p.regs.HL() - 1)
}

// Mnemonic: CCF
// Sets Flags: -00c
func (p *GBProcessor) ccf(opcode byte, params ...byte) {
	p.regs.F ^= GBFlagCarry
	p.regs.F &= ^(GBFlagSubtract | GBFlagHalfCarry)
}

// Mnemonic: HALT
// Sets Flags: ----
func (p *GBProcessor) halt(opcode byte, params ...byte) {

}

// Mnemonic: LD r1,r2
//   r1,r2 := A,B,C,D,E,H,L,(HL)
// Sets Flags: ----
func (p *GBProcessor) ld_rr(opcode byte, params ...byte) {
	var dest *uint8

	switch {
	case opcode >= 0x78:
		dest = &p.regs.A
	case opcode >= 0x70:
		v := p.readAddress(p.regs.HL())
		dest = &v
	case opcode >= 0x68:
		dest = &p.regs.L
	case opcode >= 0x60:
		dest = &p.regs.H
	case opcode >= 0x58:
		dest = &p.regs.E
	case opcode >= 0x50:
		dest = &p.regs.D
	case opcode >= 0x48:
		dest = &p.regs.C
	case opcode >= 0x40:
		dest = &p.regs.B
	}

	switch opcode & 0x7 {
	case 0x0, 0x8:
		*dest = p.regs.B
	case 0x1, 0x9:
		*dest = p.regs.C
	case 0x2, 0xA:
		*dest = p.regs.D
	case 0x3, 0xB:
		*dest = p.regs.E
	case 0x4, 0xC:
		*dest = p.regs.H
	case 0x5, 0xD:
		*dest = p.regs.L
	case 0x6, 0xE:
		*dest = p.readAddress(p.regs.HL())
	case 0x7, 0xF:
		*dest = p.regs.A
	}

}

// Mnemonic: ADD A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z0hc
func (p *GBProcessor) add_an(opcode byte, params ...byte) {
	var val uint8
	var f uint8

	switch opcode {
	case 0x87:
		val = p.regs.A
	case 0x80:
		val = p.regs.B
	case 0x81:
		val = p.regs.C
	case 0x82:
		val = p.regs.D
	case 0x83:
		val = p.regs.E
	case 0x84:
		val = p.regs.H
	case 0x85:
		val = p.regs.L
	case 0x86:
		val = p.readAddress(p.regs.HL())
	case 0xC6:
		val = params[0]
	}

	res16Bit := uint16(val) + uint16(p.regs.A)
	if res16Bit > 0xFF {
		f |= GBFlagCarry
	}

	if (val&0xF)+(p.regs.A&0xF) > 0xF {
		f |= GBFlagHalfCarry
	}

	p.regs.A += val
	if p.regs.A == 0 {
		f |= GBFlagZero
	}

	p.regs.F = f
}

// Mnemonic: ADC A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z0hc
func (p *GBProcessor) adc_n(opcode byte, params ...byte) {
	var val, carry uint8
	var f uint8

	switch opcode {
	case 0x85:
		val = p.regs.A
	case 0x88:
		val = p.regs.B
	case 0x89:
		val = p.regs.C
	case 0x8A:
		val = p.regs.D
	case 0x8B:
		val = p.regs.E
	case 0x8C:
		val = p.regs.H
	case 0x8D:
		val = p.regs.L
	case 0x8E:
		val = p.readAddress(p.regs.HL())
	case 0xCE:
		val = params[0]
	}

	carry = (p.regs.F & GBFlagCarry) / GBFlagCarry
	val += carry

	res16Bit := uint16(val) + uint16(p.regs.A)
	if res16Bit > 0xFF {
		f |= GBFlagCarry
	}

	if (val&0xF)+(p.regs.A&0xF) > 0xF {
		f |= GBFlagHalfCarry
	}

	p.regs.A += val
	if p.regs.A == 0 {
		f |= GBFlagZero
	}

	p.regs.F = f
}

// Mnemonic: SUB A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z1hc
func (p *GBProcessor) sub_n(opcode byte, params ...byte) {
	var val uint8
	var f uint8 = GBFlagSubtract

	switch opcode {
	case 0x97:
		val = p.regs.A
	case 0x90:
		val = p.regs.B
	case 0x91:
		val = p.regs.C
	case 0x92:
		val = p.regs.D
	case 0x93:
		val = p.regs.E
	case 0x94:
		val = p.regs.H
	case 0x95:
		val = p.regs.L
	case 0x96:
		val = p.readAddress(p.regs.HL())
	case 0xD6:
		val = params[0]
	}

	if p.regs.A&0xF >= val&0xF {
		f |= GBFlagHalfCarry
	}
	if p.regs.A >= val {
		f |= GBFlagCarry
	}

	p.regs.A -= val
	if p.regs.A == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f
}

// Mnemonic: SBC A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z1hc
func (p *GBProcessor) sbc_n(opcode byte, params ...byte) {
	var val, carry uint8
	var f uint8 = GBFlagSubtract

	switch opcode {
	case 0x9F:
		val = p.regs.A
	case 0x98:
		val = p.regs.B
	case 0x99:
		val = p.regs.C
	case 0x9A:
		val = p.regs.D
	case 0x9B:
		val = p.regs.E
	case 0x9C:
		val = p.regs.H
	case 0x9D:
		val = p.regs.L
	case 0x9E:
		val = p.readAddress(p.regs.HL())
	case 0xDE:
		val = params[0]
	}

	carry = (p.regs.F & GBFlagCarry) / GBFlagCarry
	val += carry

	if p.regs.A&0xF >= val&0xF {
		f |= GBFlagHalfCarry
	}
	if p.regs.A >= val {
		f |= GBFlagCarry
	}

	p.regs.A -= val
	if p.regs.A == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f
}

// Mnemonic: AND n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z010
func (p *GBProcessor) and_n(opcode byte, params ...byte) {
	var val uint8
	var f uint8

	switch opcode {
	case 0xA7:
		val = p.regs.A
	case 0xA0:
		val = p.regs.B
	case 0xA1:
		val = p.regs.C
	case 0xA2:
		val = p.regs.D
	case 0xA3:
		val = p.regs.E
	case 0xA4:
		val = p.regs.H
	case 0xA5:
		val = p.regs.L
	case 0xA6:
		val = p.readAddress(p.regs.HL())
	case 0xE6:
		val = params[0]
	}

	p.regs.A = p.regs.A & val

	if p.regs.A == 0 {
		f |= GBFlagZero
	}

	f |= GBFlagHalfCarry
	p.regs.F = f
}

// Mnemonic: XOR n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z000
func (p *GBProcessor) xor_n(opcode byte, params ...byte) {
	var val uint8
	var f uint8

	switch opcode {
	case 0xAF:
		val = p.regs.A
	case 0xA8:
		val = p.regs.B
	case 0xA9:
		val = p.regs.C
	case 0xAA:
		val = p.regs.D
	case 0xAB:
		val = p.regs.E
	case 0xAC:
		val = p.regs.H
	case 0xAD:
		val = p.regs.L
	case 0xAE:
		val = p.readAddress(p.regs.HL())
	case 0xEE:
		val = params[0]
	}

	p.regs.A = p.regs.A ^ val

	if p.regs.A == 0 {
		f |= GBFlagZero
	}

	p.regs.F = f
}

// Mnemonic: OR n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z000
func (p *GBProcessor) or_n(opcode byte, params ...byte) {
	var val uint8
	var f uint8

	switch opcode {
	case 0xB7:
		val = p.regs.A
	case 0xB0:
		val = p.regs.B
	case 0xB1:
		val = p.regs.C
	case 0xB2:
		val = p.regs.D
	case 0xB3:
		val = p.regs.E
	case 0xB4:
		val = p.regs.H
	case 0xB5:
		val = p.regs.L
	case 0xB6:
		val = p.readAddress(p.regs.HL())
	case 0xF6:
		val = params[0]
	}

	p.regs.A = p.regs.A | val

	if p.regs.A == 0 {
		f |= GBFlagZero
	}

	p.regs.F = f
}

// Mnemonic: CP n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z1hc
func (p *GBProcessor) cp_n(opcode byte, params ...byte) {
	var val uint8
	var f uint8 = GBFlagSubtract

	switch opcode {
	case 0xBF:
		val = p.regs.A
	case 0xB8:
		val = p.regs.B
	case 0xB9:
		val = p.regs.C
	case 0xBA:
		val = p.regs.D
	case 0xBB:
		val = p.regs.E
	case 0xBC:
		val = p.regs.H
	case 0xBD:
		val = p.regs.L
	case 0xBE:
		val = p.readAddress(p.regs.HL())
	case 0xFE:
		val = params[0]
	}

	if p.regs.A&0xF >= val&0xF {
		f |= GBFlagHalfCarry
	}
	if p.regs.A >= val {
		f |= GBFlagCarry
	}

	if p.regs.A-val == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f
}

// Mnemonic: RET NZ
// Sets Flags: ----
func (p *GBProcessor) ret_nz(opcode byte, params ...byte) {

}

// Mnemonic: POP rr
//  rr := AF,BC,DE,HL
// Sets Flags: ----
func (p *GBProcessor) pop(opcode byte, params ...byte) {
	switch opcode {
	case 0xf1:
		p.regs.SetAF(p.popStack())
	case 0xc1:
		p.regs.SetBC(p.popStack())
	case 0xd1:
		p.regs.SetDE(p.popStack())
	case 0xe1:
		p.regs.SetHL(p.popStack())
	}
}

// Mnemonic: JP NZ,a16
// Sets Flags: ----
func (p *GBProcessor) jp_nzi(opcode byte, params ...byte) {

}

// Mnemonic: JP a16
// Sets Flags: ----
func (p *GBProcessor) jp_i(opcode byte, params ...byte) {

}

// Mnemonic: CALL NZ,a16
// Sets Flags: ----
func (p *GBProcessor) call_nzi(opcode byte, params ...byte) {

}

// Mnemonic: RET Z
// Sets Flags: ----
func (p *GBProcessor) ret_z(opcode byte, params ...byte) {

}

// Mnemonic: RET
// Sets Flags: ----
func (p *GBProcessor) ret(opcode byte, params ...byte) {

}

// Mnemonic: JP Z,a16
// Sets Flags: ----
func (p *GBProcessor) jp_zi(opcode byte, params ...byte) {

}

// Mnemonic: PREFIX CB
// Sets Flags: ----
func (p *GBProcessor) prefix_cb(opcode byte, params ...byte) {

}

// Mnemonic: CALL Z,a16
// Sets Flags: ----
func (p *GBProcessor) call_zi(opcode byte, params ...byte) {

}

// Mnemonic: CALL a16
// Sets Flags: ----
func (p *GBProcessor) call_i(opcode byte, params ...byte) {

}

// Mnemonic: RET NC
// Sets Flags: ----
func (p *GBProcessor) ret_nc(opcode byte, params ...byte) {

}

// Mnemonic: JP NC,a16
// Sets Flags: ----
func (p *GBProcessor) jp_nci(opcode byte, params ...byte) {

}

// Mnemonic: CALL NC,a16
// Sets Flags: ----
func (p *GBProcessor) call_nci(opcode byte, params ...byte) {

}

// Mnemonic: RET C
// Sets Flags: ----
func (p *GBProcessor) ret_c(opcode byte, params ...byte) {

}

// Mnemonic: RETI
// Sets Flags: ----
func (p *GBProcessor) reti(opcode byte, params ...byte) {

}

// Mnemonic: JP C,a16
// Sets Flags: ----
func (p *GBProcessor) jp_ci(opcode byte, params ...byte) {

}

// Mnemonic: CALL C,a16
// Sets Flags: ----
func (p *GBProcessor) call_ci(opcode byte, params ...byte) {

}

// Mnemonic: LDH (a8),A
// Sets Flags: ----
func (p *GBProcessor) ldh_xia(opcode byte, params ...byte) {
	p.writeAddress(uint16(params[0])+0xFF00, p.regs.A)
}

// Mnemonic: LD (C),A
// Sets Flags: ----
func (p *GBProcessor) ld_xca(opcode byte, params ...byte) {
	p.writeAddress(uint16(p.regs.C)+0xFF00, p.regs.A)
}

// Mnemonic: ADD SP,r8
// Sets Flags: 00hc
func (p *GBProcessor) add_spi(opcode byte, params ...byte) {

}

// Mnemonic: JP (HL)
// Sets Flags: ----
func (p *GBProcessor) jp_xhl(opcode byte, params ...byte) {

}

// Mnemonic: LD (a16),A
// Sets Flags: ----
func (p *GBProcessor) ld_xia(opcode byte, params ...byte) {

}

// Mnemonic: LDH A,(a8)
// Sets Flags: ----
func (p *GBProcessor) ldh_axi(opcode byte, params ...byte) {
	p.regs.A = p.readAddress(uint16(params[0]) + 0xFF00)
}

// Mnemonic: LD A,(C)
// Sets Flags: ----
func (p *GBProcessor) ld_axc(opcode byte, params ...byte) {
	p.regs.A = p.readAddress(uint16(p.regs.C) + 0xFF00)
}

// Mnemonic: DI
// Sets Flags: ----
func (p *GBProcessor) di(opcode byte, params ...byte) {

}

// Mnemonic: PUSH rr
//  rr := AF,BC, DE, HL
// Sets Flags: ----
func (p *GBProcessor) push_rr(opcode byte, params ...byte) {
	switch opcode {
	case 0xf5:
		p.pushStack(p.regs.AF())
	case 0xc5:
		p.pushStack(p.regs.BC())
	case 0xd5:
		p.pushStack(p.regs.DE())
	case 0xe5:
		p.pushStack(p.regs.HL())
	}
}

// Mnemonic: RST a
//  a := 0H,8H,10H,18H,20H,28H,30H,38H
// Sets Flags: ----
func (p *GBProcessor) rst(opcode byte, params ...byte) {
	var addr uint16
	switch opcode {
	case 0xc7:
		addr = 0x0
	case 0xcf:
		addr = 0x8
	case 0xd7:
		addr = 0x10
	case 0xdf:
		addr = 0x18
	case 0xe7:
		addr = 0x20
	case 0xef:
		addr = 0x28
	case 0xf7:
		addr = 0x30
	case 0xff:
		addr = 0x38
	}
	p.pushStack(p.regs.PC)
	p.regs.PC = addr
}

// Mnemonic: LD HL,SP+r8
// Sets Flags: 00hc
func (p *GBProcessor) ld_hlspi(opcode byte, params ...byte) {
	var f uint8
	val := uint16(params[0])

	if (val&0xF)+(p.regs.SP&0xF) > 0xF {
		f |= GBFlagHalfCarry
	}
	if (val&0xFF)+(p.regs.SP&0xFF) > 0xFF {
		f |= GBFlagCarry
	}

	p.regs.SetHL(p.regs.SP + val)
	p.regs.F = f
}

// Mnemonic: LD SP,HL
// Sets Flags: ----
func (p *GBProcessor) ld_sphl(opcode byte, params ...byte) {
	p.regs.SP = p.regs.HL()
}

// Mnemonic: LD A,(a16)
// Sets Flags: ----
func (p *GBProcessor) ld_axi(opcode byte, params ...byte) {
	addr := uint16(params[1])<<8 | uint16(params[0])
	p.regs.A = p.readAddress(addr)
}

// Mnemonic: EI
// Sets Flags: ----
func (p *GBProcessor) ei(opcode byte, params ...byte) {

}

// Mnemonic: RLC r
//   r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_rlc_r(opcode byte, params ...byte) {
	var dest *uint8

	switch opcode {
	case 0x07:
		dest = &p.regs.A
	case 0x00:
		dest = &p.regs.B
	case 0x01:
		dest = &p.regs.C
	case 0x02:
		dest = &p.regs.D
	case 0x03:
		dest = &p.regs.E
	case 0x04:
		dest = &p.regs.H
	case 0x05:
		dest = &p.regs.L
	case 0x06:
		v := p.readAddress(p.regs.HL())
		dest = &v
	}

	if *dest&0x80 > 0 {
		*dest = (*dest << 1) + 1
		p.regs.F = GBFlagCarry
	} else {
		p.regs.A = p.regs.A << 1
		p.regs.F = 0
	}

	if *dest == 0 {
		p.regs.F |= GBFlagZero
	}

	if opcode == 0x06 {
		p.writeAddress(p.regs.HL(), *dest)
	}
}

// Mnemonic: RRC r
//   r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_rrc_r(opcode byte, params ...byte) {
	var dest *uint8

	switch opcode {
	case 0x0F:
		dest = &p.regs.A
	case 0x08:
		dest = &p.regs.B
	case 0x09:
		dest = &p.regs.C
	case 0x0A:
		dest = &p.regs.D
	case 0x0B:
		dest = &p.regs.E
	case 0x0C:
		dest = &p.regs.H
	case 0x0D:
		dest = &p.regs.L
	case 0x0E:
		v := p.readAddress(p.regs.HL())
		dest = &v
	}

	if p.regs.A&0x01 > 0 {
		p.regs.A = (p.regs.A >> 1) | 0x80
		p.regs.F = GBFlagCarry
	} else {
		p.regs.A = p.regs.A >> 1
		p.regs.F = 0
	}

	if p.regs.A == 0 {
		p.regs.F |= GBFlagZero
	}

	if opcode == 0x0E {
		p.writeAddress(p.regs.HL(), *dest)
	}
}

// Mnemonic: RL r
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_r(opcode byte, params ...byte) {
	var f uint8
	var val *uint8
	switch opcode {
	case 0x17:
		val = &p.regs.A
	case 0x10:
		val = &p.regs.B
	case 0x11:
		val = &p.regs.C
	case 0x12:
		val = &p.regs.D
	case 0x13:
		val = &p.regs.E
	case 0x14:
		val = &p.regs.H
	case 0x15:
		val = &p.regs.L
	case 0x16:
		v := p.readAddress(p.regs.HL())
		val = &v
	}
	high := (*val & 0xF0)
	low := (p.regs.F & GBFlagCarry) / GBFlagCarry
	*val = (*val << 1) | low
	if *val == 0 {
		f |= GBFlagZero
	}
	if high > 0 {
		f |= GBFlagCarry
	}
	p.regs.F = f
	if opcode == 0x16 {
		p.writeAddress(p.regs.HL(), *val)
	}
}

// Mnemonic: RR r
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_r(opcode byte, params ...byte) {
	var f uint8
	var val *uint8
	switch opcode {
	case 0x1F:
		val = &p.regs.A
	case 0x18:
		val = &p.regs.B
	case 0x19:
		val = &p.regs.C
	case 0x1A:
		val = &p.regs.D
	case 0x1B:
		val = &p.regs.E
	case 0x1C:
		val = &p.regs.H
	case 0x1D:
		val = &p.regs.L
	case 0x1E:
		v := p.readAddress(p.regs.HL())
		val = &v
	}
	low := (*val & 0x1)
	high := (p.regs.F & GBFlagCarry) / GBFlagCarry
	*val = (*val >> 1) | (high << 7)
	if *val == 0 {
		f |= GBFlagZero
	}
	if low > 0 {
		f |= GBFlagCarry
	}
	p.regs.F = f
	if opcode == 0x16 {
		p.writeAddress(p.regs.HL(), *val)
	}
}

// Mnemonic: SLA r
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_r(opcode byte, params ...byte) {
	var val *uint8
	var f uint8
	switch opcode {
	case 0x27:
		val = &p.regs.A
	case 0x20:
		val = &p.regs.B
	case 0x21:
		val = &p.regs.C
	case 0x22:
		val = &p.regs.D
	case 0x23:
		val = &p.regs.E
	case 0x24:
		val = &p.regs.H
	case 0x25:
		val = &p.regs.L
	case 0x26:
		v := p.readAddress(p.regs.HL())
		val = &v
	}
	if *val&0x80 > 0 {
		f = GBFlagCarry
	}
	*val <<= 1
	if *val == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f

	if opcode == 0x26 {
		p.writeAddress(p.regs.HL(), *val)
	}

}

// Mnemonic: SRA r
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_sra_r(opcode byte, params ...byte) {
	var val *uint8
	var f uint8
	switch opcode {
	case 0x2F:
		val = &p.regs.A
	case 0x28:
		val = &p.regs.B
	case 0x29:
		val = &p.regs.C
	case 0x2A:
		val = &p.regs.D
	case 0x2B:
		val = &p.regs.E
	case 0x2C:
		val = &p.regs.H
	case 0x2D:
		val = &p.regs.L
	case 0x2E:
		v := p.readAddress(p.regs.HL())
		val = &v
	}
	if *val&0x1 > 0 {
		f |= GBFlagCarry
	}
	*val = (*val >> 1) | (*val & 0x80)
	if *val == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f

	if opcode == 0x2E {
		p.writeAddress(p.regs.HL(), *val)
	}

}

// Mnemonic: SWAP r
//   r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z000
func (p *GBProcessor) cb_swap_r(opcode byte, params ...byte) {
	var val *uint8
	var f uint8

	switch opcode {
	case 0x37:
		val = &p.regs.A
	case 0x30:
		val = &p.regs.B
	case 0x31:
		val = &p.regs.C
	case 0x32:
		val = &p.regs.D
	case 0x33:
		val = &p.regs.E
	case 0x34:
		val = &p.regs.H
	case 0x35:
		val = &p.regs.L
	case 0x36:
		v := p.readAddress(p.regs.HL())
		val = &v
	}

	*val = (*val&0x0F)<<4 | (*val&0xF0)>>4

	if *val == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f

	if opcode == 0x36 {
		p.writeAddress(p.regs.HL(), *val)
	}
}

// Mnemonic: SRL r
//   r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_r(opcode byte, params ...byte) {
	var val *uint8
	var f uint8
	switch opcode {
	case 0x3F:
		val = &p.regs.A
	case 0x38:
		val = &p.regs.B
	case 0x39:
		val = &p.regs.C
	case 0x3A:
		val = &p.regs.D
	case 0x3B:
		val = &p.regs.E
	case 0x3C:
		val = &p.regs.H
	case 0x3D:
		val = &p.regs.L
	case 0x3E:
		v := p.readAddress(p.regs.HL())
		val = &v
	}
	if *val&0x1 > 0 {
		f |= GBFlagCarry
	}
	*val >>= 1
	if *val == 0 {
		f |= GBFlagZero
	}
	p.regs.F = f

	if opcode == 0x3E {
		p.writeAddress(p.regs.HL(), *val)
	}
}

// Mnemonic: BIT b,r
//  b := 0,1,2,3,4,5,6,7
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_br(opcode byte, params ...byte) {
	var bit uint8
	var val uint8
	var f uint8

	switch {
	case opcode >= 0x78:
		bit = 1 << 7
	case opcode >= 0x70:
		bit = 1 << 6
	case opcode >= 0x68:
		bit = 1 << 5
	case opcode >= 0x60:
		bit = 1 << 4
	case opcode >= 0x58:
		bit = 1 << 3
	case opcode >= 0x50:
		bit = 1 << 2
	case opcode >= 0x48:
		bit = 1 << 1
	case opcode >= 0x40:
		bit = 1
	}

	switch opcode & 0x0F {
	case 0, 8:
		val = p.regs.B
	case 1, 9:
		val = p.regs.C
	case 2, 0xA:
		val = p.regs.D
	case 3, 0xB:
		val = p.regs.E
	case 4, 0xC:
		val = p.regs.H
	case 5, 0xD:
		val = p.regs.L
	case 6, 0xE:
		val = p.readAddress(p.regs.HL())
	case 7, 0xF:
		val = p.regs.A
	}

	if val&bit == 0 {
		f |= GBFlagZero
	}
	f |= (p.regs.F & GBFlagCarry) | GBFlagHalfCarry
}

// Mnemonic: RES b,r
//  b := 0,1,2,3,4,5,6,7
//  r := A,B,C,D,E,H,L,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_br(opcode byte, params ...byte) {
	var bit uint8

	switch true {
	case opcode >= 0x78:
		bit = 1 << 7
	case opcode >= 0x70:
		bit = 1 << 6
	case opcode >= 0x68:
		bit = 1 << 5
	case opcode >= 0x60:
		bit = 1 << 4
	case opcode >= 0x58:
		bit = 1 << 3
	case opcode >= 0x50:
		bit = 1 << 2
	case opcode >= 0x48:
		bit = 1 << 1
	case opcode >= 0x40:
		bit = 1
	}

	switch opcode & 0x0F {
	case 0, 8:
		p.regs.B &= ^bit
	case 1, 9:
		p.regs.C &= ^bit
	case 2, 0xA:
		p.regs.D &= ^bit
	case 3, 0xB:
		p.regs.E &= ^bit
	case 4, 0xC:
		p.regs.H &= ^bit
	case 5, 0xD:
		p.regs.L &= ^bit
	case 6, 0xE:
		v := p.readAddress(p.regs.HL())
		p.writeAddress(p.regs.HL(), v&(^bit))
	case 7, 0xF:
		p.regs.A &= ^bit
	}

}

// Mnemonic: SET b,r
//   b := 0,1,2,3,4,5,6,7
//   r := A,B,C,D,E,H,L,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_set_br(opcode byte, params ...byte) {
	var bit uint8

	switch true {
	case opcode >= 0xF8:
		bit = 1 << 7
	case opcode >= 0xF0:
		bit = 1 << 6
	case opcode >= 0xE8:
		bit = 1 << 5
	case opcode >= 0xE0:
		bit = 1 << 4
	case opcode >= 0xD8:
		bit = 1 << 3
	case opcode >= 0xD0:
		bit = 1 << 2
	case opcode >= 0xC8:
		bit = 1 << 1
	case opcode >= 0xC0:
		bit = 1
	}

	switch opcode & 0x0F {
	case 0, 8:
		p.regs.B |= bit
	case 1, 9:
		p.regs.C |= bit
	case 2, 0xA:
		p.regs.D |= bit
	case 3, 0xB:
		p.regs.E |= bit
	case 4, 0xC:
		p.regs.H |= bit
	case 5, 0xD:
		p.regs.L |= bit
	case 6, 0xE:
		v := p.readAddress(p.regs.HL())
		p.writeAddress(p.regs.HL(), v|bit)
	case 7, 0xF:
		p.regs.A |= bit
	}
}
