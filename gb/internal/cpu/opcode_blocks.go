package cpu

import "fmt"

func (c *Cpu) blockOneOpcodeHandler(opcode byte) error {
	err := c.loadRegisterToRegister(opcode)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cpu) blockTwoOpcodeHandler(opcode byte) error {
	var err error

	actionBits := (opcode >> 3) & 0b00111
	switch actionBits {
	case 0:
		err = c.addRegister(opcode)
	case 1:
		err = c.addCarryRegister(opcode)
	case 2:
		err = c.subRegister(opcode)
	case 3:
		err = c.subCarryRegister(opcode)
	case 4:
		err = c.andRegister(opcode)
	case 5:
		err = c.xorRegister(opcode)
	case 6:
		err = c.orRegister(opcode)
	case 7:
		err = c.compareRegister(opcode)
	default:
		err = fmt.Errorf("unknown opcode")
	}

	return err
}
