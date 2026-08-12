package cpu

import "fmt"

func (c *Cpu) execute(opcode byte) (uint16, int, error) {
	if opcode == 0x00 { // NOP
		return 1, 4, nil
	}

	if opcode == 0xCB { // CB prefix
		return c.handleCBOpcode(opcode)
	}

	blockBits := opcode >> 6
	switch blockBits {
	case 0:
		return c.blockZeroOpcode(opcode)
	case 1:
		return c.blockOneOpcodeHandler(opcode)
	case 2:
		return c.blockTwoOpcodeHandler(opcode)
	case 3:
	}

	return 0, 0, fmt.Errorf("unhandled opcode: 0x%02X", opcode)
}

func (c *Cpu) handleCBOpcode(opcode byte) (uint16, int, error) {
	blockBits := opcode >> 6

	switch blockBits {

	}

	return 0, 0, fmt.Errorf("unhandled CB opcode: 0x%02X", opcode)
}

func (c *Cpu) blockZeroOpcode(opcode byte) (uint16, int, error) {
	if opcode == 0x00 { // NOP
		return 1, 4, nil
	}

	// TODO
	lastThreeBits := opcode & 0b00000111
	if lastThreeBits == 0b000 {
		// jr imm8
		// jr cond imm8
	} else if lastThreeBits == 0b111 {
		// rlca
		// rrca
		// rla
		// rra
		// daa
		// cpl
		// scf
		// ccf
	} else if lastThreeBits == 0b011 {

	}

	// lastFourBits := opcode & 0b00001111
	// if lastFourBits ==
	return 1, 4, nil
}

func (c *Cpu) blockOneOpcodeHandler(opcode byte) (uint16, int, error) {
	err := c.loadRegToReg(opcode)
	if err != nil {
		return 0, 0, fmt.Errorf("failed opcode 0x%02X: %v", opcode, err)
	}
	return 1, 4, nil
}

func (c *Cpu) blockTwoOpcodeHandler(opcode byte) (uint16, int, error) {
	actionBits := (opcode >> 3) & 0b00111
	// operand := opcode & 0b00000111

	if actionBits == 0 {
		err := c.addToRegister(opcode)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x%02X: %v", opcode, err)
		}
		return 1, 4, nil
	}

	return 1, 4, nil
}
