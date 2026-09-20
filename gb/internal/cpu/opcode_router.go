package cpu

import "fmt"

func (c *Cpu) execute(opcode byte) error {
	if opcode == 0x00 { // NOP
		return nil
	}

	if opcode == 0xCB { // CB prefix
		return c.handleCBOpcode(opcode)
	}

	var err error

	blockBits := opcode >> 6
	switch blockBits {
	case 0:
		err = c.blockZeroOpcodeHandler(opcode)
	case 1:
		err = c.blockOneOpcodeHandler(opcode)
	case 2:
		err = c.blockTwoOpcodeHandler(opcode)
	case 3:
	}

	if err != nil {
		return fmt.Errorf("opcode: 0x%02X %w", opcode, err)
	}
	return nil
}

func (c *Cpu) handleCBOpcode(opcode byte) error {
	blockBits := opcode >> 6

	switch blockBits {

	}

	return fmt.Errorf("unhandled CB opcode: 0x%02X", opcode)
}
