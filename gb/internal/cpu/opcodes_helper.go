package cpu

import "fmt"

func (c *Cpu) loadRegToReg(opcode uint16) error {
	storeRegister := (opcode >> 3) & 0b00111
	loadRegister := opcode & 0b00000111

	loadValue, err := c.loadRegister(uint8(loadRegister))
	if err != nil {
		return fmt.Errorf("failed to load register: %v", err)
	}
	err = c.storeRegister(uint8(storeRegister), loadValue)
	if err != nil {
		return fmt.Errorf("failed to store register: %v", err)
	}

	return nil
}

func (c *Cpu) loadRegister(register uint8) (uint8, error) {
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
		value, err := c.mmu.ReadByteAt(c.registers.HL())
		if err != nil {
			return 0, fmt.Errorf("failed to read from memory at address 0x%04X: %v", c.registers.HL(), err)
		}
		return value, nil
	case 7:
		return c.registers.A, nil
	default:
		return 0, fmt.Errorf("invalid register: %d", register)
	}
}

func (c *Cpu) storeRegister(register uint8, value uint8) error {
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
		err := c.mmu.WriteByteAt(c.registers.HL(), value)
		if err != nil {
			return fmt.Errorf("failed to write to memory at address 0x%04X: %v", c.registers.HL(), err)
		}
	case 7:
		c.registers.A = value
	default:
		return fmt.Errorf("invalid register: %d", register)
	}

	return nil
}
