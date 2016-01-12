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

// Mnemonic: LD (BC),A
// Sets Flags: ----
func (p *GBProcessor) ld_xbca(opcode byte, params ...byte) {

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

// Mnemonic: LD (DE),A
// Sets Flags: ----
func (p *GBProcessor) ld_xdea(opcode byte, params ...byte) {

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

// Mnemonic: LD (HL+),A
// Sets Flags: ----
func (p *GBProcessor) ld_xhlia(opcode byte, params ...byte) {

}

// Mnemonic: DAA
// Sets Flags: z-0c
func (p *GBProcessor) daa(opcode byte, params ...byte) {

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

}

// Mnemonic: CPL
// Sets Flags: -11-
func (p *GBProcessor) cpl(opcode byte, params ...byte) {
	p.regs.A = ^p.regs.A
	p.regs.F = GBFlagSubtract | GBFlagHalfCarry

}

// Mnemonic: JR NC,r8
// Sets Flags: ----
func (p *GBProcessor) jr_nci(opcode byte, params ...byte) {

}

// Mnemonic: LD (HL-),A
// Sets Flags: ----
func (p *GBProcessor) ld_xhlda(opcode byte, params ...byte) {

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

}

// Mnemonic: LD A,d8
// Sets Flags: ----
func (p *GBProcessor) ld_ai(opcode byte, params ...byte) {

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
//   r1,r2 := A,B,C,D,E,H,L,(HL),n
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
	case 0, 8:
		*dest = p.regs.B
	case 1, 9:
		*dest = p.regs.C
	case 2, 0xA:
		*dest = p.regs.D
	case 3, 0xB:
		*dest = p.regs.E
	case 4, 0xC:
		*dest = p.regs.H
	case 5, 0xD:
		*dest = p.regs.L
	case 6, 0xE:
		*dest = p.readAddress(p.regs.HL())
	case 7, 0xF:
		*dest = p.regs.A
	}

}

// Mnemonic: ADD A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z0hc
func (p *GBProcessor) add_an(opcode byte, params ...byte) {
	var val *uint8
	var f uint8

	switch opcode {
	case 0x87:
		val = &p.regs.A
	case 0x80:
		val = &p.regs.B
	case 0x81:
		val = &p.regs.C
	case 0x82:
		val = &p.regs.D
	case 0x83:
		val = &p.regs.E
	case 0x84:
		val = &p.regs.H
	case 0x85:
		val = &p.regs.L
	case 0x86:
		v := p.readAddress(p.regs.HL())
		val = &v
	case 0xC6:
		val = &params[0]
	}

	res16Bit := uint16(*val) + uint16(p.regs.A)
	if res16Bit&0x100 > 0 {
		f |= GBFlagCarry
	}

	if ((*val&0xF)+(p.regs.A&0xF))&0x10 > 0 {
		f |= GBFlagHalfCarry
	}

	p.regs.A = *val + p.regs.A
	if p.regs.A == 0 {
		f |= GBFlagZero
	}
}

// Mnemonic: ADC A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z0hc
func (p *GBProcessor) adc_n(opcode byte, params ...byte) {

}

// Mnemonic: SUB A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z1hc
func (p *GBProcessor) sub_n(opcode byte, params ...byte) {

}

// Mnemonic: SBC A,n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z1hc
func (p *GBProcessor) sbc_n(opcode byte, params ...byte) {

}

// Mnemonic: AND n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z010
func (p *GBProcessor) and_n(opcode byte, params ...byte) {
	var val *uint8
	var f uint8

	switch opcode {
	case 0xA7:
		val = &p.regs.A
	case 0xA0:
		val = &p.regs.B
	case 0xA1:
		val = &p.regs.C
	case 0xA2:
		val = &p.regs.D
	case 0xA3:
		val = &p.regs.E
	case 0xA4:
		val = &p.regs.H
	case 0xA5:
		val = &p.regs.L
	case 0xA6:
		v := p.readAddress(p.regs.HL())
		val = &v
	case 0xE6:
		val = &params[0]
	}

	p.regs.A = p.regs.A & *val

	if p.regs.A == 0 {
		f |= GBFlagZero
	}

	f |= GBFlagHalfCarry
}

// Mnemonic: XOR n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z000
func (p *GBProcessor) xor_n(opcode byte, params ...byte) {
	var val *uint8
	var f uint8

	switch opcode {
	case 0xAF:
		val = &p.regs.A
	case 0xA8:
		val = &p.regs.B
	case 0xA9:
		val = &p.regs.C
	case 0xAA:
		val = &p.regs.D
	case 0xAB:
		val = &p.regs.E
	case 0xAC:
		val = &p.regs.H
	case 0xAD:
		val = &p.regs.L
	case 0xAE:
		v := p.readAddress(p.regs.HL())
		val = &v
	case 0xEE:
		val = &params[0]
	}

	p.regs.A = p.regs.A ^ *val

	if p.regs.A == 0 {
		f |= GBFlagZero
	}
}

// Mnemonic: OR n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z000
func (p *GBProcessor) or_n(opcode byte, params ...byte) {
	var val *uint8
	var f uint8

	switch opcode {
	case 0xB7:
		val = &p.regs.A
	case 0xB0:
		val = &p.regs.B
	case 0xB1:
		val = &p.regs.C
	case 0xB2:
		val = &p.regs.D
	case 0xB3:
		val = &p.regs.E
	case 0xB4:
		val = &p.regs.H
	case 0xB5:
		val = &p.regs.L
	case 0xB6:
		v := p.readAddress(p.regs.HL())
		val = &v
	case 0xF6:
		val = &params[0]
	}

	p.regs.A = p.regs.A | *val

	if p.regs.A == 0 {
		f |= GBFlagZero
	}
}

// Mnemonic: CP n
//  n := A,B,C,D,E,H,L,(HL),#
// Sets Flags: z1hc
func (p *GBProcessor) cp_n(opcode byte, params ...byte) {

}

// Mnemonic: RET NZ
// Sets Flags: ----
func (p *GBProcessor) ret_nz(opcode byte, params ...byte) {

}

// Mnemonic: POP BC
// Sets Flags: ----
func (p *GBProcessor) pop_bc(opcode byte, params ...byte) {

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

// Mnemonic: PUSH BC
// Sets Flags: ----
func (p *GBProcessor) push_bc(opcode byte, params ...byte) {

}

// Mnemonic: RST 00H
// Sets Flags: ----
func (p *GBProcessor) rst_00h(opcode byte, params ...byte) {

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

// Mnemonic: RST 08H
// Sets Flags: ----
func (p *GBProcessor) rst_08h(opcode byte, params ...byte) {

}

// Mnemonic: RET NC
// Sets Flags: ----
func (p *GBProcessor) ret_nc(opcode byte, params ...byte) {

}

// Mnemonic: POP DE
// Sets Flags: ----
func (p *GBProcessor) pop_de(opcode byte, params ...byte) {

}

// Mnemonic: JP NC,a16
// Sets Flags: ----
func (p *GBProcessor) jp_nci(opcode byte, params ...byte) {

}

// Mnemonic: CALL NC,a16
// Sets Flags: ----
func (p *GBProcessor) call_nci(opcode byte, params ...byte) {

}

// Mnemonic: PUSH DE
// Sets Flags: ----
func (p *GBProcessor) push_de(opcode byte, params ...byte) {

}

// Mnemonic: RST 10H
// Sets Flags: ----
func (p *GBProcessor) rst_10h(opcode byte, params ...byte) {

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

// Mnemonic: RST 18H
// Sets Flags: ----
func (p *GBProcessor) rst_18h(opcode byte, params ...byte) {

}

// Mnemonic: LDH (a8),A
// Sets Flags: ----
func (p *GBProcessor) ldh_xia(opcode byte, params ...byte) {

}

// Mnemonic: POP HL
// Sets Flags: ----
func (p *GBProcessor) pop_hl(opcode byte, params ...byte) {

}

// Mnemonic: LD (C),A
// Sets Flags: ----
func (p *GBProcessor) ld_xca(opcode byte, params ...byte) {

}

// Mnemonic: PUSH HL
// Sets Flags: ----
func (p *GBProcessor) push_hl(opcode byte, params ...byte) {

}

// Mnemonic: RST 20H
// Sets Flags: ----
func (p *GBProcessor) rst_20h(opcode byte, params ...byte) {

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

// Mnemonic: RST 28H
// Sets Flags: ----
func (p *GBProcessor) rst_28h(opcode byte, params ...byte) {

}

// Mnemonic: LDH A,(a8)
// Sets Flags: ----
func (p *GBProcessor) ldh_axi(opcode byte, params ...byte) {

}

// Mnemonic: POP AF
// Sets Flags: znhc
func (p *GBProcessor) pop_af(opcode byte, params ...byte) {

}

// Mnemonic: LD A,(C)
// Sets Flags: ----
func (p *GBProcessor) ld_axc(opcode byte, params ...byte) {

}

// Mnemonic: DI
// Sets Flags: ----
func (p *GBProcessor) di(opcode byte, params ...byte) {

}

// Mnemonic: PUSH AF
// Sets Flags: ----
func (p *GBProcessor) push_af(opcode byte, params ...byte) {

}

// Mnemonic: RST 30H
// Sets Flags: ----
func (p *GBProcessor) rst_30h(opcode byte, params ...byte) {

}

// Mnemonic: LD HL,SP+r8
// Sets Flags: 00hc
func (p *GBProcessor) ld_hlspi(opcode byte, params ...byte) {

}

// Mnemonic: LD SP,HL
// Sets Flags: ----
func (p *GBProcessor) ld_sphl(opcode byte, params ...byte) {
	p.regs.SP = p.regs.HL()
}

// Mnemonic: LD A,(a16)
// Sets Flags: ----
func (p *GBProcessor) ld_axi(opcode byte, params ...byte) {

}

// Mnemonic: EI
// Sets Flags: ----
func (p *GBProcessor) ei(opcode byte, params ...byte) {

}

// Mnemonic: RST 38H
// Sets Flags: ----
func (p *GBProcessor) rst_38h(opcode byte, params ...byte) {

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

// Mnemonic: RL B
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_b(opcode byte, params ...byte) {

}

// Mnemonic: RL C
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_c(opcode byte, params ...byte) {

}

// Mnemonic: RL D
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_d(opcode byte, params ...byte) {

}

// Mnemonic: RL E
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_e(opcode byte, params ...byte) {

}

// Mnemonic: RL H
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_h(opcode byte, params ...byte) {

}

// Mnemonic: RL L
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_l(opcode byte, params ...byte) {

}

// Mnemonic: RL (HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_xhl(opcode byte, params ...byte) {

}

// Mnemonic: RL A
// Sets Flags: z00c
func (p *GBProcessor) cb_rl_a(opcode byte, params ...byte) {

}

// Mnemonic: RR B
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_b(opcode byte, params ...byte) {

}

// Mnemonic: RR C
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_c(opcode byte, params ...byte) {

}

// Mnemonic: RR D
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_d(opcode byte, params ...byte) {

}

// Mnemonic: RR E
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_e(opcode byte, params ...byte) {

}

// Mnemonic: RR H
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_h(opcode byte, params ...byte) {

}

// Mnemonic: RR L
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_l(opcode byte, params ...byte) {

}

// Mnemonic: RR (HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_xhl(opcode byte, params ...byte) {

}

// Mnemonic: RR A
// Sets Flags: z00c
func (p *GBProcessor) cb_rr_a(opcode byte, params ...byte) {

}

// Mnemonic: SLA B
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_b(opcode byte, params ...byte) {

}

// Mnemonic: SLA C
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_c(opcode byte, params ...byte) {

}

// Mnemonic: SLA D
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_d(opcode byte, params ...byte) {

}

// Mnemonic: SLA E
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_e(opcode byte, params ...byte) {

}

// Mnemonic: SLA H
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_h(opcode byte, params ...byte) {

}

// Mnemonic: SLA L
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_l(opcode byte, params ...byte) {

}

// Mnemonic: SLA (HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_xhl(opcode byte, params ...byte) {

}

// Mnemonic: SLA A
// Sets Flags: z00c
func (p *GBProcessor) cb_sla_a(opcode byte, params ...byte) {

}

// Mnemonic: SRA B
// Sets Flags: z000
func (p *GBProcessor) cb_sra_b(opcode byte, params ...byte) {

}

// Mnemonic: SRA C
// Sets Flags: z000
func (p *GBProcessor) cb_sra_c(opcode byte, params ...byte) {

}

// Mnemonic: SRA D
// Sets Flags: z000
func (p *GBProcessor) cb_sra_d(opcode byte, params ...byte) {

}

// Mnemonic: SRA E
// Sets Flags: z000
func (p *GBProcessor) cb_sra_e(opcode byte, params ...byte) {

}

// Mnemonic: SRA H
// Sets Flags: z000
func (p *GBProcessor) cb_sra_h(opcode byte, params ...byte) {

}

// Mnemonic: SRA L
// Sets Flags: z000
func (p *GBProcessor) cb_sra_l(opcode byte, params ...byte) {

}

// Mnemonic: SRA (HL)
// Sets Flags: z000
func (p *GBProcessor) cb_sra_xhl(opcode byte, params ...byte) {

}

// Mnemonic: SRA A
// Sets Flags: z000
func (p *GBProcessor) cb_sra_a(opcode byte, params ...byte) {

}

// Mnemonic: SWAP B
// Sets Flags: z000
func (p *GBProcessor) cb_swap_b(opcode byte, params ...byte) {

}

// Mnemonic: SWAP C
// Sets Flags: z000
func (p *GBProcessor) cb_swap_c(opcode byte, params ...byte) {

}

// Mnemonic: SWAP D
// Sets Flags: z000
func (p *GBProcessor) cb_swap_d(opcode byte, params ...byte) {

}

// Mnemonic: SWAP E
// Sets Flags: z000
func (p *GBProcessor) cb_swap_e(opcode byte, params ...byte) {

}

// Mnemonic: SWAP H
// Sets Flags: z000
func (p *GBProcessor) cb_swap_h(opcode byte, params ...byte) {

}

// Mnemonic: SWAP L
// Sets Flags: z000
func (p *GBProcessor) cb_swap_l(opcode byte, params ...byte) {

}

// Mnemonic: SWAP (HL)
// Sets Flags: z000
func (p *GBProcessor) cb_swap_xhl(opcode byte, params ...byte) {

}

// Mnemonic: SWAP A
// Sets Flags: z000
func (p *GBProcessor) cb_swap_a(opcode byte, params ...byte) {

}

// Mnemonic: SRL B
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_b(opcode byte, params ...byte) {

}

// Mnemonic: SRL C
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_c(opcode byte, params ...byte) {

}

// Mnemonic: SRL D
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_d(opcode byte, params ...byte) {

}

// Mnemonic: SRL E
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_e(opcode byte, params ...byte) {

}

// Mnemonic: SRL H
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_h(opcode byte, params ...byte) {

}

// Mnemonic: SRL L
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_l(opcode byte, params ...byte) {

}

// Mnemonic: SRL (HL)
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_xhl(opcode byte, params ...byte) {

}

// Mnemonic: SRL A
// Sets Flags: z00c
func (p *GBProcessor) cb_srl_a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 0,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_0a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 1,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_1a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 2,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_2a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 3,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_3a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 4,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_4a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 5,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_5a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 6,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_6a(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,B
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7b(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,C
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7c(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,D
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7d(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,E
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7e(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,H
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7h(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,L
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7l(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,(HL)
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7xhl(opcode byte, params ...byte) {

}

// Mnemonic: BIT 7,A
// Sets Flags: z01-
func (p *GBProcessor) cb_bit_7a(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_0b(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_0c(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_0d(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_0e(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_0h(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_0l(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_0xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 0,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_0a(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_1b(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_1c(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_1d(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_1e(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_1h(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_1l(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_1xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 1,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_1a(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_2b(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_2c(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_2d(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_2e(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_2h(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_2l(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_2xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 2,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_2a(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_3b(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_3c(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_3d(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_3e(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_3h(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_3l(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_3xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 3,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_3a(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_4b(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_4c(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_4d(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_4e(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_4h(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_4l(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_4xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 4,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_4a(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_5b(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_5c(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_5d(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_5e(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_5h(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_5l(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_5xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 5,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_5a(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_6b(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_6c(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_6d(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_6e(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_6h(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_6l(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_6xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 6,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_6a(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,B
// Sets Flags: ----
func (p *GBProcessor) cb_res_7b(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,C
// Sets Flags: ----
func (p *GBProcessor) cb_res_7c(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,D
// Sets Flags: ----
func (p *GBProcessor) cb_res_7d(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,E
// Sets Flags: ----
func (p *GBProcessor) cb_res_7e(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,H
// Sets Flags: ----
func (p *GBProcessor) cb_res_7h(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,L
// Sets Flags: ----
func (p *GBProcessor) cb_res_7l(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,(HL)
// Sets Flags: ----
func (p *GBProcessor) cb_res_7xhl(opcode byte, params ...byte) {

}

// Mnemonic: RES 7,A
// Sets Flags: ----
func (p *GBProcessor) cb_res_7a(opcode byte, params ...byte) {

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

	switch opcode & 0x7 {
	case 0:
		p.regs.B |= bit
	case 1:
		p.regs.C |= bit
	case 2:
		p.regs.D |= bit
	case 3:
		p.regs.E |= bit
	case 4:
		p.regs.H |= bit
	case 5:
		p.regs.L |= bit
	case 6:
		v := p.readAddress(p.regs.HL())
		p.writeAddress(p.regs.HL(), v|bit)
	case 7:
		p.regs.A |= bit
	}
}
