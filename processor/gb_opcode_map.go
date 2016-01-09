// vim: noet:ts=3:sw=3:sts=3
package processor

// This file was auto-generated

func (p *GBProcessor) initOpcodes() {
	p.opcodes = []Opcode{
		{0x00, 1, 4, 4, (*GBProcessor).nop},
		{0x01, 3, 12, 12, (*GBProcessor).ld_rri},
		{0x02, 1, 8, 8, (*GBProcessor).ld_xbca},
		{0x03, 1, 8, 8, (*GBProcessor).inc_rr},
		{0x04, 1, 4, 4, (*GBProcessor).inc_r},
		{0x05, 1, 4, 4, (*GBProcessor).dec_r},
		{0x06, 2, 8, 8, (*GBProcessor).ld_ri},
		{0x07, 1, 4, 4, (*GBProcessor).rlca},
		{0x08, 3, 20, 20, (*GBProcessor).ld_xisp},
		{0x09, 1, 8, 8, (*GBProcessor).add_hlbc},
		{0x0a, 1, 8, 8, (*GBProcessor).ld_axrr},
		{0x0b, 1, 8, 8, (*GBProcessor).dec_rr},
		{0x0c, 1, 4, 4, (*GBProcessor).inc_r},
		{0x0d, 1, 4, 4, (*GBProcessor).dec_r},
		{0x0e, 2, 8, 8, (*GBProcessor).ld_ri},
		{0x0f, 1, 4, 4, (*GBProcessor).rrca},
		{0x10, 2, 4, 4, (*GBProcessor).stop_0},
		{0x11, 3, 12, 12, (*GBProcessor).ld_rri},
		{0x12, 1, 8, 8, (*GBProcessor).ld_xdea},
		{0x13, 1, 8, 8, (*GBProcessor).inc_rr},
		{0x14, 1, 4, 4, (*GBProcessor).inc_r},
		{0x15, 1, 4, 4, (*GBProcessor).dec_r},
		{0x16, 2, 8, 8, (*GBProcessor).ld_ri},
		{0x17, 1, 4, 4, (*GBProcessor).rla},
		{0x18, 2, 12, 12, (*GBProcessor).jr_i},
		{0x19, 1, 8, 8, (*GBProcessor).add_hlde},
		{0x1a, 1, 8, 8, (*GBProcessor).ld_axrr},
		{0x1b, 1, 8, 8, (*GBProcessor).dec_rr},
		{0x1c, 1, 4, 4, (*GBProcessor).inc_r},
		{0x1d, 1, 4, 4, (*GBProcessor).dec_r},
		{0x1e, 2, 8, 8, (*GBProcessor).ld_ri},
		{0x1f, 1, 4, 4, (*GBProcessor).rra},
		{0x20, 2, 8, 12, (*GBProcessor).jr_nzi},
		{0x21, 3, 12, 12, (*GBProcessor).ld_rri},
		{0x22, 1, 8, 8, (*GBProcessor).ld_xhlia},
		{0x23, 1, 8, 8, (*GBProcessor).inc_rr},
		{0x24, 1, 4, 4, (*GBProcessor).inc_r},
		{0x25, 1, 4, 4, (*GBProcessor).dec_r},
		{0x26, 2, 8, 8, (*GBProcessor).ld_ri},
		{0x27, 1, 4, 4, (*GBProcessor).daa},
		{0x28, 2, 8, 12, (*GBProcessor).jr_zi},
		{0x29, 1, 8, 8, (*GBProcessor).add_hlhl},
		{0x2a, 1, 8, 8, (*GBProcessor).ld_axhli},
		{0x2b, 1, 8, 8, (*GBProcessor).dec_rr},
		{0x2c, 1, 4, 4, (*GBProcessor).inc_r},
		{0x2d, 1, 4, 4, (*GBProcessor).dec_r},
		{0x2e, 2, 8, 8, (*GBProcessor).ld_ri},
		{0x2f, 1, 4, 4, (*GBProcessor).cpl},
		{0x30, 2, 8, 12, (*GBProcessor).jr_nci},
		{0x31, 3, 12, 12, (*GBProcessor).ld_rri},
		{0x32, 1, 8, 8, (*GBProcessor).ld_xhlda},
		{0x33, 1, 8, 8, (*GBProcessor).inc_rr},
		{0x34, 1, 12, 12, (*GBProcessor).inc_r},
		{0x35, 1, 12, 12, (*GBProcessor).dec_r},
		{0x36, 2, 12, 12, (*GBProcessor).ld_xhli},
		{0x37, 1, 4, 4, (*GBProcessor).scf},
		{0x38, 2, 8, 12, (*GBProcessor).jr_ci},
		{0x39, 1, 8, 8, (*GBProcessor).add_hlsp},
		{0x3a, 1, 8, 8, (*GBProcessor).ld_axhld},
		{0x3b, 1, 8, 8, (*GBProcessor).dec_rr},
		{0x3c, 1, 4, 4, (*GBProcessor).inc_r},
		{0x3d, 1, 4, 4, (*GBProcessor).dec_r},
		{0x3e, 2, 8, 8, (*GBProcessor).ld_ai},
		{0x3f, 1, 4, 4, (*GBProcessor).ccf},
		{0x40, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x41, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x42, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x43, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x44, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x45, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x46, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x47, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x48, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x49, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x4a, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x4b, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x4c, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x4d, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x4e, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x4f, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x50, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x51, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x52, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x53, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x54, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x55, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x56, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x57, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x58, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x59, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x5a, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x5b, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x5c, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x5d, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x5e, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x5f, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x60, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x61, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x62, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x63, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x64, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x65, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x66, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x67, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x68, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x69, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x6a, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x6b, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x6c, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x6d, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x6e, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x6f, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x70, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x71, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x72, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x73, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x74, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x75, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x76, 1, 4, 4, (*GBProcessor).halt},
		{0x77, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x78, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x79, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x7a, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x7b, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x7c, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x7d, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x7e, 1, 8, 8, (*GBProcessor).ld_rr},
		{0x7f, 1, 4, 4, (*GBProcessor).ld_rr},
		{0x80, 1, 4, 4, (*GBProcessor).add_ab},
		{0x81, 1, 4, 4, (*GBProcessor).add_ac},
		{0x82, 1, 4, 4, (*GBProcessor).add_ad},
		{0x83, 1, 4, 4, (*GBProcessor).add_ae},
		{0x84, 1, 4, 4, (*GBProcessor).add_ah},
		{0x85, 1, 4, 4, (*GBProcessor).add_al},
		{0x86, 1, 8, 8, (*GBProcessor).add_axhl},
		{0x87, 1, 4, 4, (*GBProcessor).add_aa},
		{0x88, 1, 4, 4, (*GBProcessor).adc_ab},
		{0x89, 1, 4, 4, (*GBProcessor).adc_ac},
		{0x8a, 1, 4, 4, (*GBProcessor).adc_ad},
		{0x8b, 1, 4, 4, (*GBProcessor).adc_ae},
		{0x8c, 1, 4, 4, (*GBProcessor).adc_ah},
		{0x8d, 1, 4, 4, (*GBProcessor).adc_al},
		{0x8e, 1, 8, 8, (*GBProcessor).adc_axhl},
		{0x8f, 1, 4, 4, (*GBProcessor).adc_aa},
		{0x90, 1, 4, 4, (*GBProcessor).sub_b},
		{0x91, 1, 4, 4, (*GBProcessor).sub_c},
		{0x92, 1, 4, 4, (*GBProcessor).sub_d},
		{0x93, 1, 4, 4, (*GBProcessor).sub_e},
		{0x94, 1, 4, 4, (*GBProcessor).sub_h},
		{0x95, 1, 4, 4, (*GBProcessor).sub_l},
		{0x96, 1, 8, 8, (*GBProcessor).sub_xhl},
		{0x97, 1, 4, 4, (*GBProcessor).sub_a},
		{0x98, 1, 4, 4, (*GBProcessor).sbc_ab},
		{0x99, 1, 4, 4, (*GBProcessor).sbc_ac},
		{0x9a, 1, 4, 4, (*GBProcessor).sbc_ad},
		{0x9b, 1, 4, 4, (*GBProcessor).sbc_ae},
		{0x9c, 1, 4, 4, (*GBProcessor).sbc_ah},
		{0x9d, 1, 4, 4, (*GBProcessor).sbc_al},
		{0x9e, 1, 8, 8, (*GBProcessor).sbc_axhl},
		{0x9f, 1, 4, 4, (*GBProcessor).sbc_aa},
		{0xa0, 1, 4, 4, (*GBProcessor).and_b},
		{0xa1, 1, 4, 4, (*GBProcessor).and_c},
		{0xa2, 1, 4, 4, (*GBProcessor).and_d},
		{0xa3, 1, 4, 4, (*GBProcessor).and_e},
		{0xa4, 1, 4, 4, (*GBProcessor).and_h},
		{0xa5, 1, 4, 4, (*GBProcessor).and_l},
		{0xa6, 1, 8, 8, (*GBProcessor).and_xhl},
		{0xa7, 1, 4, 4, (*GBProcessor).and_a},
		{0xa8, 1, 4, 4, (*GBProcessor).xor_b},
		{0xa9, 1, 4, 4, (*GBProcessor).xor_c},
		{0xaa, 1, 4, 4, (*GBProcessor).xor_d},
		{0xab, 1, 4, 4, (*GBProcessor).xor_e},
		{0xac, 1, 4, 4, (*GBProcessor).xor_h},
		{0xad, 1, 4, 4, (*GBProcessor).xor_l},
		{0xae, 1, 8, 8, (*GBProcessor).xor_xhl},
		{0xaf, 1, 4, 4, (*GBProcessor).xor_a},
		{0xb0, 1, 4, 4, (*GBProcessor).or_b},
		{0xb1, 1, 4, 4, (*GBProcessor).or_c},
		{0xb2, 1, 4, 4, (*GBProcessor).or_d},
		{0xb3, 1, 4, 4, (*GBProcessor).or_e},
		{0xb4, 1, 4, 4, (*GBProcessor).or_h},
		{0xb5, 1, 4, 4, (*GBProcessor).or_l},
		{0xb6, 1, 8, 8, (*GBProcessor).or_xhl},
		{0xb7, 1, 4, 4, (*GBProcessor).or_a},
		{0xb8, 1, 4, 4, (*GBProcessor).cp_b},
		{0xb9, 1, 4, 4, (*GBProcessor).cp_c},
		{0xba, 1, 4, 4, (*GBProcessor).cp_d},
		{0xbb, 1, 4, 4, (*GBProcessor).cp_e},
		{0xbc, 1, 4, 4, (*GBProcessor).cp_h},
		{0xbd, 1, 4, 4, (*GBProcessor).cp_l},
		{0xbe, 1, 8, 8, (*GBProcessor).cp_xhl},
		{0xbf, 1, 4, 4, (*GBProcessor).cp_a},
		{0xc0, 1, 8, 20, (*GBProcessor).ret_nz},
		{0xc1, 1, 12, 12, (*GBProcessor).pop_bc},
		{0xc2, 3, 12, 16, (*GBProcessor).jp_nzi},
		{0xc3, 3, 16, 16, (*GBProcessor).jp_i},
		{0xc4, 3, 12, 24, (*GBProcessor).call_nzi},
		{0xc5, 1, 16, 16, (*GBProcessor).push_bc},
		{0xc6, 2, 8, 8, (*GBProcessor).add_ai},
		{0xc7, 1, 16, 16, (*GBProcessor).rst_00h},
		{0xc8, 1, 8, 20, (*GBProcessor).ret_z},
		{0xc9, 1, 16, 16, (*GBProcessor).ret},
		{0xca, 3, 12, 16, (*GBProcessor).jp_zi},
		{0xcb, 1, 4, 4, (*GBProcessor).prefix_cb},
		{0xcc, 3, 12, 24, (*GBProcessor).call_zi},
		{0xcd, 3, 24, 24, (*GBProcessor).call_i},
		{0xce, 2, 8, 8, (*GBProcessor).adc_ai},
		{0xcf, 1, 16, 16, (*GBProcessor).rst_08h},
		{0xd0, 1, 8, 20, (*GBProcessor).ret_nc},
		{0xd1, 1, 12, 12, (*GBProcessor).pop_de},
		{0xd2, 3, 12, 16, (*GBProcessor).jp_nci},
		{},
		{0xd4, 3, 12, 24, (*GBProcessor).call_nci},
		{0xd5, 1, 16, 16, (*GBProcessor).push_de},
		{0xd6, 2, 8, 8, (*GBProcessor).sub_i},
		{0xd7, 1, 16, 16, (*GBProcessor).rst_10h},
		{0xd8, 1, 8, 20, (*GBProcessor).ret_c},
		{0xd9, 1, 16, 16, (*GBProcessor).reti},
		{0xda, 3, 12, 16, (*GBProcessor).jp_ci},
		{},
		{0xdc, 3, 12, 24, (*GBProcessor).call_ci},
		{},
		{0xde, 2, 8, 8, (*GBProcessor).sbc_ai},
		{0xdf, 1, 16, 16, (*GBProcessor).rst_18h},
		{0xe0, 2, 12, 12, (*GBProcessor).ldh_xia},
		{0xe1, 1, 12, 12, (*GBProcessor).pop_hl},
		{0xe2, 2, 8, 8, (*GBProcessor).ld_xca},
		{},
		{},
		{0xe5, 1, 16, 16, (*GBProcessor).push_hl},
		{0xe6, 2, 8, 8, (*GBProcessor).and_i},
		{0xe7, 1, 16, 16, (*GBProcessor).rst_20h},
		{0xe8, 2, 16, 16, (*GBProcessor).add_spi},
		{0xe9, 1, 4, 4, (*GBProcessor).jp_xhl},
		{0xea, 3, 16, 16, (*GBProcessor).ld_xia},
		{},
		{},
		{},
		{0xee, 2, 8, 8, (*GBProcessor).xor_i},
		{0xef, 1, 16, 16, (*GBProcessor).rst_28h},
		{0xf0, 2, 12, 12, (*GBProcessor).ldh_axi},
		{0xf1, 1, 12, 12, (*GBProcessor).pop_af},
		{0xf2, 2, 8, 8, (*GBProcessor).ld_axc},
		{0xf3, 1, 4, 4, (*GBProcessor).di},
		{},
		{0xf5, 1, 16, 16, (*GBProcessor).push_af},
		{0xf6, 2, 8, 8, (*GBProcessor).or_i},
		{0xf7, 1, 16, 16, (*GBProcessor).rst_30h},
		{0xf8, 2, 12, 12, (*GBProcessor).ld_hlspi},
		{0xf9, 1, 8, 8, (*GBProcessor).ld_sphl},
		{0xfa, 3, 16, 16, (*GBProcessor).ld_axi},
		{0xfb, 1, 4, 4, (*GBProcessor).ei},
		{},
		{},
		{0xfe, 2, 8, 8, (*GBProcessor).cp_i},
		{0xff, 1, 16, 16, (*GBProcessor).rst_38h},
	}
	p.cbOpcodes = []Opcode{
		{0x00, 2, 8, 8, (*GBProcessor).cb_rlc_b},
		{0x01, 2, 8, 8, (*GBProcessor).cb_rlc_c},
		{0x02, 2, 8, 8, (*GBProcessor).cb_rlc_d},
		{0x03, 2, 8, 8, (*GBProcessor).cb_rlc_e},
		{0x04, 2, 8, 8, (*GBProcessor).cb_rlc_h},
		{0x05, 2, 8, 8, (*GBProcessor).cb_rlc_l},
		{0x06, 2, 16, 16, (*GBProcessor).cb_rlc_xhl},
		{0x07, 2, 8, 8, (*GBProcessor).cb_rlc_a},
		{0x08, 2, 8, 8, (*GBProcessor).cb_rrc_b},
		{0x09, 2, 8, 8, (*GBProcessor).cb_rrc_c},
		{0x0a, 2, 8, 8, (*GBProcessor).cb_rrc_d},
		{0x0b, 2, 8, 8, (*GBProcessor).cb_rrc_e},
		{0x0c, 2, 8, 8, (*GBProcessor).cb_rrc_h},
		{0x0d, 2, 8, 8, (*GBProcessor).cb_rrc_l},
		{0x0e, 2, 16, 16, (*GBProcessor).cb_rrc_xhl},
		{0x0f, 2, 8, 8, (*GBProcessor).cb_rrc_a},
		{0x10, 2, 8, 8, (*GBProcessor).cb_rl_b},
		{0x11, 2, 8, 8, (*GBProcessor).cb_rl_c},
		{0x12, 2, 8, 8, (*GBProcessor).cb_rl_d},
		{0x13, 2, 8, 8, (*GBProcessor).cb_rl_e},
		{0x14, 2, 8, 8, (*GBProcessor).cb_rl_h},
		{0x15, 2, 8, 8, (*GBProcessor).cb_rl_l},
		{0x16, 2, 16, 16, (*GBProcessor).cb_rl_xhl},
		{0x17, 2, 8, 8, (*GBProcessor).cb_rl_a},
		{0x18, 2, 8, 8, (*GBProcessor).cb_rr_b},
		{0x19, 2, 8, 8, (*GBProcessor).cb_rr_c},
		{0x1a, 2, 8, 8, (*GBProcessor).cb_rr_d},
		{0x1b, 2, 8, 8, (*GBProcessor).cb_rr_e},
		{0x1c, 2, 8, 8, (*GBProcessor).cb_rr_h},
		{0x1d, 2, 8, 8, (*GBProcessor).cb_rr_l},
		{0x1e, 2, 16, 16, (*GBProcessor).cb_rr_xhl},
		{0x1f, 2, 8, 8, (*GBProcessor).cb_rr_a},
		{0x20, 2, 8, 8, (*GBProcessor).cb_sla_b},
		{0x21, 2, 8, 8, (*GBProcessor).cb_sla_c},
		{0x22, 2, 8, 8, (*GBProcessor).cb_sla_d},
		{0x23, 2, 8, 8, (*GBProcessor).cb_sla_e},
		{0x24, 2, 8, 8, (*GBProcessor).cb_sla_h},
		{0x25, 2, 8, 8, (*GBProcessor).cb_sla_l},
		{0x26, 2, 16, 16, (*GBProcessor).cb_sla_xhl},
		{0x27, 2, 8, 8, (*GBProcessor).cb_sla_a},
		{0x28, 2, 8, 8, (*GBProcessor).cb_sra_b},
		{0x29, 2, 8, 8, (*GBProcessor).cb_sra_c},
		{0x2a, 2, 8, 8, (*GBProcessor).cb_sra_d},
		{0x2b, 2, 8, 8, (*GBProcessor).cb_sra_e},
		{0x2c, 2, 8, 8, (*GBProcessor).cb_sra_h},
		{0x2d, 2, 8, 8, (*GBProcessor).cb_sra_l},
		{0x2e, 2, 16, 16, (*GBProcessor).cb_sra_xhl},
		{0x2f, 2, 8, 8, (*GBProcessor).cb_sra_a},
		{0x30, 2, 8, 8, (*GBProcessor).cb_swap_b},
		{0x31, 2, 8, 8, (*GBProcessor).cb_swap_c},
		{0x32, 2, 8, 8, (*GBProcessor).cb_swap_d},
		{0x33, 2, 8, 8, (*GBProcessor).cb_swap_e},
		{0x34, 2, 8, 8, (*GBProcessor).cb_swap_h},
		{0x35, 2, 8, 8, (*GBProcessor).cb_swap_l},
		{0x36, 2, 16, 16, (*GBProcessor).cb_swap_xhl},
		{0x37, 2, 8, 8, (*GBProcessor).cb_swap_a},
		{0x38, 2, 8, 8, (*GBProcessor).cb_srl_b},
		{0x39, 2, 8, 8, (*GBProcessor).cb_srl_c},
		{0x3a, 2, 8, 8, (*GBProcessor).cb_srl_d},
		{0x3b, 2, 8, 8, (*GBProcessor).cb_srl_e},
		{0x3c, 2, 8, 8, (*GBProcessor).cb_srl_h},
		{0x3d, 2, 8, 8, (*GBProcessor).cb_srl_l},
		{0x3e, 2, 16, 16, (*GBProcessor).cb_srl_xhl},
		{0x3f, 2, 8, 8, (*GBProcessor).cb_srl_a},
		{0x40, 2, 8, 8, (*GBProcessor).cb_bit_0b},
		{0x41, 2, 8, 8, (*GBProcessor).cb_bit_0c},
		{0x42, 2, 8, 8, (*GBProcessor).cb_bit_0d},
		{0x43, 2, 8, 8, (*GBProcessor).cb_bit_0e},
		{0x44, 2, 8, 8, (*GBProcessor).cb_bit_0h},
		{0x45, 2, 8, 8, (*GBProcessor).cb_bit_0l},
		{0x46, 2, 16, 16, (*GBProcessor).cb_bit_0xhl},
		{0x47, 2, 8, 8, (*GBProcessor).cb_bit_0a},
		{0x48, 2, 8, 8, (*GBProcessor).cb_bit_1b},
		{0x49, 2, 8, 8, (*GBProcessor).cb_bit_1c},
		{0x4a, 2, 8, 8, (*GBProcessor).cb_bit_1d},
		{0x4b, 2, 8, 8, (*GBProcessor).cb_bit_1e},
		{0x4c, 2, 8, 8, (*GBProcessor).cb_bit_1h},
		{0x4d, 2, 8, 8, (*GBProcessor).cb_bit_1l},
		{0x4e, 2, 16, 16, (*GBProcessor).cb_bit_1xhl},
		{0x4f, 2, 8, 8, (*GBProcessor).cb_bit_1a},
		{0x50, 2, 8, 8, (*GBProcessor).cb_bit_2b},
		{0x51, 2, 8, 8, (*GBProcessor).cb_bit_2c},
		{0x52, 2, 8, 8, (*GBProcessor).cb_bit_2d},
		{0x53, 2, 8, 8, (*GBProcessor).cb_bit_2e},
		{0x54, 2, 8, 8, (*GBProcessor).cb_bit_2h},
		{0x55, 2, 8, 8, (*GBProcessor).cb_bit_2l},
		{0x56, 2, 16, 16, (*GBProcessor).cb_bit_2xhl},
		{0x57, 2, 8, 8, (*GBProcessor).cb_bit_2a},
		{0x58, 2, 8, 8, (*GBProcessor).cb_bit_3b},
		{0x59, 2, 8, 8, (*GBProcessor).cb_bit_3c},
		{0x5a, 2, 8, 8, (*GBProcessor).cb_bit_3d},
		{0x5b, 2, 8, 8, (*GBProcessor).cb_bit_3e},
		{0x5c, 2, 8, 8, (*GBProcessor).cb_bit_3h},
		{0x5d, 2, 8, 8, (*GBProcessor).cb_bit_3l},
		{0x5e, 2, 16, 16, (*GBProcessor).cb_bit_3xhl},
		{0x5f, 2, 8, 8, (*GBProcessor).cb_bit_3a},
		{0x60, 2, 8, 8, (*GBProcessor).cb_bit_4b},
		{0x61, 2, 8, 8, (*GBProcessor).cb_bit_4c},
		{0x62, 2, 8, 8, (*GBProcessor).cb_bit_4d},
		{0x63, 2, 8, 8, (*GBProcessor).cb_bit_4e},
		{0x64, 2, 8, 8, (*GBProcessor).cb_bit_4h},
		{0x65, 2, 8, 8, (*GBProcessor).cb_bit_4l},
		{0x66, 2, 16, 16, (*GBProcessor).cb_bit_4xhl},
		{0x67, 2, 8, 8, (*GBProcessor).cb_bit_4a},
		{0x68, 2, 8, 8, (*GBProcessor).cb_bit_5b},
		{0x69, 2, 8, 8, (*GBProcessor).cb_bit_5c},
		{0x6a, 2, 8, 8, (*GBProcessor).cb_bit_5d},
		{0x6b, 2, 8, 8, (*GBProcessor).cb_bit_5e},
		{0x6c, 2, 8, 8, (*GBProcessor).cb_bit_5h},
		{0x6d, 2, 8, 8, (*GBProcessor).cb_bit_5l},
		{0x6e, 2, 16, 16, (*GBProcessor).cb_bit_5xhl},
		{0x6f, 2, 8, 8, (*GBProcessor).cb_bit_5a},
		{0x70, 2, 8, 8, (*GBProcessor).cb_bit_6b},
		{0x71, 2, 8, 8, (*GBProcessor).cb_bit_6c},
		{0x72, 2, 8, 8, (*GBProcessor).cb_bit_6d},
		{0x73, 2, 8, 8, (*GBProcessor).cb_bit_6e},
		{0x74, 2, 8, 8, (*GBProcessor).cb_bit_6h},
		{0x75, 2, 8, 8, (*GBProcessor).cb_bit_6l},
		{0x76, 2, 16, 16, (*GBProcessor).cb_bit_6xhl},
		{0x77, 2, 8, 8, (*GBProcessor).cb_bit_6a},
		{0x78, 2, 8, 8, (*GBProcessor).cb_bit_7b},
		{0x79, 2, 8, 8, (*GBProcessor).cb_bit_7c},
		{0x7a, 2, 8, 8, (*GBProcessor).cb_bit_7d},
		{0x7b, 2, 8, 8, (*GBProcessor).cb_bit_7e},
		{0x7c, 2, 8, 8, (*GBProcessor).cb_bit_7h},
		{0x7d, 2, 8, 8, (*GBProcessor).cb_bit_7l},
		{0x7e, 2, 16, 16, (*GBProcessor).cb_bit_7xhl},
		{0x7f, 2, 8, 8, (*GBProcessor).cb_bit_7a},
		{0x80, 2, 8, 8, (*GBProcessor).cb_res_0b},
		{0x81, 2, 8, 8, (*GBProcessor).cb_res_0c},
		{0x82, 2, 8, 8, (*GBProcessor).cb_res_0d},
		{0x83, 2, 8, 8, (*GBProcessor).cb_res_0e},
		{0x84, 2, 8, 8, (*GBProcessor).cb_res_0h},
		{0x85, 2, 8, 8, (*GBProcessor).cb_res_0l},
		{0x86, 2, 16, 16, (*GBProcessor).cb_res_0xhl},
		{0x87, 2, 8, 8, (*GBProcessor).cb_res_0a},
		{0x88, 2, 8, 8, (*GBProcessor).cb_res_1b},
		{0x89, 2, 8, 8, (*GBProcessor).cb_res_1c},
		{0x8a, 2, 8, 8, (*GBProcessor).cb_res_1d},
		{0x8b, 2, 8, 8, (*GBProcessor).cb_res_1e},
		{0x8c, 2, 8, 8, (*GBProcessor).cb_res_1h},
		{0x8d, 2, 8, 8, (*GBProcessor).cb_res_1l},
		{0x8e, 2, 16, 16, (*GBProcessor).cb_res_1xhl},
		{0x8f, 2, 8, 8, (*GBProcessor).cb_res_1a},
		{0x90, 2, 8, 8, (*GBProcessor).cb_res_2b},
		{0x91, 2, 8, 8, (*GBProcessor).cb_res_2c},
		{0x92, 2, 8, 8, (*GBProcessor).cb_res_2d},
		{0x93, 2, 8, 8, (*GBProcessor).cb_res_2e},
		{0x94, 2, 8, 8, (*GBProcessor).cb_res_2h},
		{0x95, 2, 8, 8, (*GBProcessor).cb_res_2l},
		{0x96, 2, 16, 16, (*GBProcessor).cb_res_2xhl},
		{0x97, 2, 8, 8, (*GBProcessor).cb_res_2a},
		{0x98, 2, 8, 8, (*GBProcessor).cb_res_3b},
		{0x99, 2, 8, 8, (*GBProcessor).cb_res_3c},
		{0x9a, 2, 8, 8, (*GBProcessor).cb_res_3d},
		{0x9b, 2, 8, 8, (*GBProcessor).cb_res_3e},
		{0x9c, 2, 8, 8, (*GBProcessor).cb_res_3h},
		{0x9d, 2, 8, 8, (*GBProcessor).cb_res_3l},
		{0x9e, 2, 16, 16, (*GBProcessor).cb_res_3xhl},
		{0x9f, 2, 8, 8, (*GBProcessor).cb_res_3a},
		{0xa0, 2, 8, 8, (*GBProcessor).cb_res_4b},
		{0xa1, 2, 8, 8, (*GBProcessor).cb_res_4c},
		{0xa2, 2, 8, 8, (*GBProcessor).cb_res_4d},
		{0xa3, 2, 8, 8, (*GBProcessor).cb_res_4e},
		{0xa4, 2, 8, 8, (*GBProcessor).cb_res_4h},
		{0xa5, 2, 8, 8, (*GBProcessor).cb_res_4l},
		{0xa6, 2, 16, 16, (*GBProcessor).cb_res_4xhl},
		{0xa7, 2, 8, 8, (*GBProcessor).cb_res_4a},
		{0xa8, 2, 8, 8, (*GBProcessor).cb_res_5b},
		{0xa9, 2, 8, 8, (*GBProcessor).cb_res_5c},
		{0xaa, 2, 8, 8, (*GBProcessor).cb_res_5d},
		{0xab, 2, 8, 8, (*GBProcessor).cb_res_5e},
		{0xac, 2, 8, 8, (*GBProcessor).cb_res_5h},
		{0xad, 2, 8, 8, (*GBProcessor).cb_res_5l},
		{0xae, 2, 16, 16, (*GBProcessor).cb_res_5xhl},
		{0xaf, 2, 8, 8, (*GBProcessor).cb_res_5a},
		{0xb0, 2, 8, 8, (*GBProcessor).cb_res_6b},
		{0xb1, 2, 8, 8, (*GBProcessor).cb_res_6c},
		{0xb2, 2, 8, 8, (*GBProcessor).cb_res_6d},
		{0xb3, 2, 8, 8, (*GBProcessor).cb_res_6e},
		{0xb4, 2, 8, 8, (*GBProcessor).cb_res_6h},
		{0xb5, 2, 8, 8, (*GBProcessor).cb_res_6l},
		{0xb6, 2, 16, 16, (*GBProcessor).cb_res_6xhl},
		{0xb7, 2, 8, 8, (*GBProcessor).cb_res_6a},
		{0xb8, 2, 8, 8, (*GBProcessor).cb_res_7b},
		{0xb9, 2, 8, 8, (*GBProcessor).cb_res_7c},
		{0xba, 2, 8, 8, (*GBProcessor).cb_res_7d},
		{0xbb, 2, 8, 8, (*GBProcessor).cb_res_7e},
		{0xbc, 2, 8, 8, (*GBProcessor).cb_res_7h},
		{0xbd, 2, 8, 8, (*GBProcessor).cb_res_7l},
		{0xbe, 2, 16, 16, (*GBProcessor).cb_res_7xhl},
		{0xbf, 2, 8, 8, (*GBProcessor).cb_res_7a},
		{0xc0, 2, 8, 8, (*GBProcessor).cb_set_0b},
		{0xc1, 2, 8, 8, (*GBProcessor).cb_set_0c},
		{0xc2, 2, 8, 8, (*GBProcessor).cb_set_0d},
		{0xc3, 2, 8, 8, (*GBProcessor).cb_set_0e},
		{0xc4, 2, 8, 8, (*GBProcessor).cb_set_0h},
		{0xc5, 2, 8, 8, (*GBProcessor).cb_set_0l},
		{0xc6, 2, 16, 16, (*GBProcessor).cb_set_0xhl},
		{0xc7, 2, 8, 8, (*GBProcessor).cb_set_0a},
		{0xc8, 2, 8, 8, (*GBProcessor).cb_set_1b},
		{0xc9, 2, 8, 8, (*GBProcessor).cb_set_1c},
		{0xca, 2, 8, 8, (*GBProcessor).cb_set_1d},
		{0xcb, 2, 8, 8, (*GBProcessor).cb_set_1e},
		{0xcc, 2, 8, 8, (*GBProcessor).cb_set_1h},
		{0xcd, 2, 8, 8, (*GBProcessor).cb_set_1l},
		{0xce, 2, 16, 16, (*GBProcessor).cb_set_1xhl},
		{0xcf, 2, 8, 8, (*GBProcessor).cb_set_1a},
		{0xd0, 2, 8, 8, (*GBProcessor).cb_set_2b},
		{0xd1, 2, 8, 8, (*GBProcessor).cb_set_2c},
		{0xd2, 2, 8, 8, (*GBProcessor).cb_set_2d},
		{0xd3, 2, 8, 8, (*GBProcessor).cb_set_2e},
		{0xd4, 2, 8, 8, (*GBProcessor).cb_set_2h},
		{0xd5, 2, 8, 8, (*GBProcessor).cb_set_2l},
		{0xd6, 2, 16, 16, (*GBProcessor).cb_set_2xhl},
		{0xd7, 2, 8, 8, (*GBProcessor).cb_set_2a},
		{0xd8, 2, 8, 8, (*GBProcessor).cb_set_3b},
		{0xd9, 2, 8, 8, (*GBProcessor).cb_set_3c},
		{0xda, 2, 8, 8, (*GBProcessor).cb_set_3d},
		{0xdb, 2, 8, 8, (*GBProcessor).cb_set_3e},
		{0xdc, 2, 8, 8, (*GBProcessor).cb_set_3h},
		{0xdd, 2, 8, 8, (*GBProcessor).cb_set_3l},
		{0xde, 2, 16, 16, (*GBProcessor).cb_set_3xhl},
		{0xdf, 2, 8, 8, (*GBProcessor).cb_set_3a},
		{0xe0, 2, 8, 8, (*GBProcessor).cb_set_4b},
		{0xe1, 2, 8, 8, (*GBProcessor).cb_set_4c},
		{0xe2, 2, 8, 8, (*GBProcessor).cb_set_4d},
		{0xe3, 2, 8, 8, (*GBProcessor).cb_set_4e},
		{0xe4, 2, 8, 8, (*GBProcessor).cb_set_4h},
		{0xe5, 2, 8, 8, (*GBProcessor).cb_set_4l},
		{0xe6, 2, 16, 16, (*GBProcessor).cb_set_4xhl},
		{0xe7, 2, 8, 8, (*GBProcessor).cb_set_4a},
		{0xe8, 2, 8, 8, (*GBProcessor).cb_set_5b},
		{0xe9, 2, 8, 8, (*GBProcessor).cb_set_5c},
		{0xea, 2, 8, 8, (*GBProcessor).cb_set_5d},
		{0xeb, 2, 8, 8, (*GBProcessor).cb_set_5e},
		{0xec, 2, 8, 8, (*GBProcessor).cb_set_5h},
		{0xed, 2, 8, 8, (*GBProcessor).cb_set_5l},
		{0xee, 2, 16, 16, (*GBProcessor).cb_set_5xhl},
		{0xef, 2, 8, 8, (*GBProcessor).cb_set_5a},
		{0xf0, 2, 8, 8, (*GBProcessor).cb_set_6b},
		{0xf1, 2, 8, 8, (*GBProcessor).cb_set_6c},
		{0xf2, 2, 8, 8, (*GBProcessor).cb_set_6d},
		{0xf3, 2, 8, 8, (*GBProcessor).cb_set_6e},
		{0xf4, 2, 8, 8, (*GBProcessor).cb_set_6h},
		{0xf5, 2, 8, 8, (*GBProcessor).cb_set_6l},
		{0xf6, 2, 16, 16, (*GBProcessor).cb_set_6xhl},
		{0xf7, 2, 8, 8, (*GBProcessor).cb_set_6a},
		{0xf8, 2, 8, 8, (*GBProcessor).cb_set_7b},
		{0xf9, 2, 8, 8, (*GBProcessor).cb_set_7c},
		{0xfa, 2, 8, 8, (*GBProcessor).cb_set_7d},
		{0xfb, 2, 8, 8, (*GBProcessor).cb_set_7e},
		{0xfc, 2, 8, 8, (*GBProcessor).cb_set_7h},
		{0xfd, 2, 8, 8, (*GBProcessor).cb_set_7l},
		{0xfe, 2, 16, 16, (*GBProcessor).cb_set_7xhl},
		{0xff, 2, 8, 8, (*GBProcessor).cb_set_7a},
	}
}