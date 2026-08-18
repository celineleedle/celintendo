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
		return err
	}

	sum := c.registers.A + addValue

	c.registers.setFlag(zeroBitIndex, sum == 0)
	c.registers.setFlag(subtractionBitIndex, false)
	c.registers.setFlag(halfCarryBitIndex, c.registers.A&0x0F+addValue&0x0F > 0x0F)
	c.registers.setFlag(carryBitIndex, uint16(c.registers.A)+uint16(addValue) > 0xFF)

	c.registers.A = sum
	return nil
}

func (c *Cpu) addCarryRegisterToA(opcode byte) error {
	srcRegister := opcode & 0b00000111
	addValue, err := c.loadRegister(srcRegister)
	if err != nil {
		return err
	}

	carry := 0
	if c.registers.getFlag(carryBitIndex) {
		carry = 1
	}

	sum := c.registers.A + addValue + uint8(carry)

	c.registers.setFlag(zeroBitIndex, sum == 0)
	c.registers.setFlag(subtractionBitIndex, false)
	c.registers.setFlag(halfCarryBitIndex, c.registers.A&0x0F+addValue&0x0F+uint8(carry) > 0x0F)
	c.registers.setFlag(carryBitIndex, uint16(c.registers.A)+uint16(addValue)+uint16(carry) > 0xFF)

	c.registers.A = sum
	return nil
}

func (c *Cpu) subRegisterFromA(opcode byte) error {
	srcRegister := opcode & 0b00000111
	subValue, err := c.loadRegister(srcRegister)
	if err != nil {
		return err
	}

	diff := c.registers.A - subValue

	c.registers.setFlag(zeroBitIndex, diff == 0)
	c.registers.setFlag(subtractionBitIndex, true)
	c.registers.setFlag(halfCarryBitIndex, c.registers.A&0x0F < subValue&0x0F)
	c.registers.setFlag(carryBitIndex, c.registers.A < subValue)

	c.registers.A = diff
	return nil
}

func (c *Cpu) subCarryRegisterFromA(opcode byte) error {
	srcRegister := opcode & 0b00000111
	subValue, err := c.loadRegister(srcRegister)
	if err != nil {
		return err
	}

	carry := 0
	if c.registers.getFlag(carryBitIndex) {
		carry = 1
	}

	diff := c.registers.A - subValue - uint8(carry)

	c.registers.setFlag(zeroBitIndex, diff == 0)
	c.registers.setFlag(subtractionBitIndex, true)
	c.registers.setFlag(halfCarryBitIndex, c.registers.A&0x0F < (subValue&0x0F+uint8(carry)))
	c.registers.setFlag(carryBitIndex, uint16(c.registers.A) < uint16(subValue+uint8(carry)))

	c.registers.A = diff
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
			return 0, err
		}
		return value, nil
	case 7:
		return c.registers.A, nil
	default:
		return 0, fmt.Errorf("load register %d", register)
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
			return err
		}
	case 7:
		c.registers.A = value
	default:
		return fmt.Errorf("store register %d", register)
	}

	return nil
}
