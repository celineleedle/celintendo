package cpu

import "fmt"

func (c *Cpu) loadRegisterToRegister(opcode byte) error {
	dstRegister := (opcode >> 3) & 0b00111
	srcRegister := opcode & 0b00000111

	loadValue, err := c.loadRegister(srcRegister)
	if err != nil {
		return err
	}
	err = c.storeRegister(dstRegister, loadValue)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cpu) addRegisterToA(opcode byte) error {
	srcRegister := opcode & 0b00000111
	addValue, err := c.loadRegister(srcRegister)
	if err != nil {
		return fmt.Errorf("failed to load register: %v", err)
	}

	c.registers.SetFlag(ZeroBitIndex, c.registers.A+addValue == 0)
	c.registers.SetFlag(SubtractionBitIndex, false)
	c.registers.SetFlag(HalfCarryBitIndex, uint16(c.registers.A&0x0F)+uint16(addValue&0x0F) > 0x0F)
	c.registers.SetFlag(CarryBitIndex, uint16(c.registers.A)+uint16(addValue) > 0xFF)

	c.registers.A += addValue
	return nil
}

func (c *Cpu) loadRegister(register byte) (uint8, error) {
	switch register {
	case 0:
		return c.registers.B, nil
	case 1:
		return c.registers.C, nil
	case 2:
		return c.registers.D, nil
	case 3:
		return c.registers.E, nil
	case 4:
		return c.registers.H, nil
	case 5:
		return c.registers.L, nil
	case 6:
		value, err := c.readCycle(c.registers.HL()) // memory access, advance time
		if err != nil {
			return 0, fmt.Errorf("invalid address in register HL: %v", err)
		}
		return value, nil
	case 7:
		return c.registers.A, nil
	default:
		return 0, fmt.Errorf("invalid register: %d", register)
	}
}

func (c *Cpu) storeRegister(register byte, value uint8) error {
	switch register {
	case 0:
		c.registers.B = value
	case 1:
		c.registers.C = value
	case 2:
		c.registers.D = value
	case 3:
		c.registers.E = value
	case 4:
		c.registers.H = value
	case 5:
		c.registers.L = value
	case 6:
		err := c.writeCycle(c.registers.HL(), value) // writing to memory, advance time
		if err != nil {
			return fmt.Errorf("invalid address in register HL: %w", err)
		}
	case 7:
		c.registers.A = value
	default:
		return fmt.Errorf("invalid register: %d", register)
	}

	return nil
}
