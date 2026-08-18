package cpu

import "fmt"

func (c *Cpu) execute(opcode byte) error {
	if opcode == 0x00 { // NOP
		return nil
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

	return fmt.Errorf("unhandled opcode: 0x%02X", opcode)
}

func (c *Cpu) handleCBOpcode(opcode byte) error {
	blockBits := opcode >> 6

	switch blockBits {

	}

	return fmt.Errorf("unhandled CB opcode: 0x%02X", opcode)
}

func (c *Cpu) blockZeroOpcode(opcode byte) error {
	if opcode == 0x00 { // NOP
		return nil
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
	return nil
}

func (c *Cpu) blockOneOpcodeHandler(opcode byte) error {
	err := c.loadRegisterToRegister(opcode)
	if err != nil {
		return fmt.Errorf("failed opcode 0x%02X: %v", opcode, err)
	}
	return nil
}

func (c *Cpu) blockTwoOpcodeHandler(opcode byte) error {
	actionBits := (opcode >> 3) & 0b00111

	var err error

	switch actionBits {
	case 0:
		err = c.addRegisterToA(opcode)
	case 1:
		err = c.addCarryRegisterToA(opcode)
	case 2:
		err = c.subRegisterFromA(opcode)
	default:
		err = fmt.Errorf("unhandled opcode: 0x%02X", opcode)
	}

	return err
}
