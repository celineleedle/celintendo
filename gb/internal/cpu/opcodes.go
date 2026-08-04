package cpu

import "fmt"

func (c *Cpu) handleOpcode(opcode byte) (uint16, int, error) {
	blockBits := opcode >> 6

	if opcode == 0xCB {
		return c.handleCBOpcode(opcode)
	}

	switch blockBits {
	case 0: // block 0
		if opcode == 0x00 { // NOP
			return 1, 4, nil
		}

	case 1: // block 1
		err := c.loadRegToReg(opcode)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x%02X: %v", opcode, err)
		}
		return 1, 4, nil
	case 2:
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

var opcodeFuncMap = map[uint8]func(*Cpu) (pc uint16, cycles int, err error){
	0x00: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // NOP
	0x01: func(c *Cpu) (uint16, int, error) { // LD BC, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x01: %v", err)
		}
		c.registers.SetBC(n16)
		return 3, 12, nil
	},
	0x02: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x03: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x04: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x05: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x06: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0x07: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x08: func(c *Cpu) (uint16, int, error) { return 3, 20, nil },
	0x09: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x0A: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x0B: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x0C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x0D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x0E: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0x0F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0x10: func(c *Cpu) (uint16, int, error) { return 2, 4, nil },
	0x11: func(c *Cpu) (uint16, int, error) { // LD DE, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x11: %v", err)
		}
		c.registers.SetDE(n16)
		return 3, 12, nil
	},
	0x12: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x13: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x14: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x15: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x16: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0x17: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x18: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x19: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x1A: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x1B: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x1C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x1D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x1E: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0x1F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0x20: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x21: func(c *Cpu) (uint16, int, error) { // LD HL, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x21: %v", err)
		}
		c.registers.SetHL(n16)
		return 3, 12, nil
	},
	0x22: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x23: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x24: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x25: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x26: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0x27: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x28: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x29: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x2A: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x2B: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x2C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x2D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x2E: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0x2F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0x30: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x31: func(c *Cpu) (uint16, int, error) { // LD SP, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x31: %v", err)
		}
		c.registers.SP = n16
		return 3, 12, nil
	},
	0x32: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x33: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x34: func(c *Cpu) (uint16, int, error) { return 1, 12, nil },
	0x35: func(c *Cpu) (uint16, int, error) { return 1, 12, nil },
	0x36: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x37: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x38: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x39: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x3A: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x3B: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x3C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x3D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x3E: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0x3F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0x40: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // LD B, B
	0x41: func(c *Cpu) (uint16, int, error) { // LD B, C
		err := c.loadRegToReg(0x41)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x41: %v", err)
		}
		return 1, 4, nil
	},
	0x42: func(c *Cpu) (uint16, int, error) { // LD B, D
		err := c.loadRegToReg(0x42)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x42: %v", err)
		}
		return 1, 4, nil
	},
	0x43: func(c *Cpu) (uint16, int, error) { // LD B, E
		err := c.loadRegToReg(0x43)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x43: %v", err)
		}
		return 1, 4, nil
	},
	0x44: func(c *Cpu) (uint16, int, error) { // LD B, H
		err := c.loadRegToReg(0x44)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x44: %v", err)
		}
		return 1, 4, nil
	},
	0x45: func(c *Cpu) (uint16, int, error) { // LD B, L
		err := c.loadRegToReg(0x45)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x45: %v", err)
		}
		return 1, 4, nil
	},
	0x46: func(c *Cpu) (uint16, int, error) { // LD B, [HL]
		err := c.loadRegToReg(0x46)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x46: %v", err)
		}
		return 1, 8, nil
	},
	0x47: func(c *Cpu) (uint16, int, error) { // LD B, A
		err := c.loadRegToReg(0x47)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x47: %v", err)
		}
		return 1, 4, nil
	},
	0x48: func(c *Cpu) (uint16, int, error) { // LD C, B
		err := c.loadRegToReg(0x48)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x48: %v", err)
		}
		return 1, 4, nil
	},
	0x49: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // LD C, C
	0x4A: func(c *Cpu) (uint16, int, error) { // LD C, D
		err := c.loadRegToReg(0x4A)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x4A: %v", err)
		}
		return 1, 4, nil
	},
	0x4B: func(c *Cpu) (uint16, int, error) { // LD C, E
		err := c.loadRegToReg(0x4B)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x4B: %v", err)
		}
		return 1, 4, nil
	},
	0x4C: func(c *Cpu) (uint16, int, error) { // LD C, H
		err := c.loadRegToReg(0x4C)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x4C: %v", err)
		}
		return 1, 4, nil
	},
	0x4D: func(c *Cpu) (uint16, int, error) { // LD C, L
		err := c.loadRegToReg(0x4D)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x4D: %v", err)
		}
		return 1, 4, nil
	},
	0x4E: func(c *Cpu) (uint16, int, error) { // LD C, [HL]
		err := c.loadRegToReg(0x4E)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x4E: %v", err)
		}
		return 1, 8, nil
	},
	0x4F: func(c *Cpu) (uint16, int, error) { // LD C, A
		err := c.loadRegToReg(0x4F)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x4F: %v", err)
		}
		return 1, 4, nil
	},

	0x50: func(c *Cpu) (uint16, int, error) { // LD D, B
		err := c.loadRegToReg(0x50)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x50: %v", err)
		}
		return 1, 4, nil
	},
	0x51: func(c *Cpu) (uint16, int, error) { // LD D, C
		err := c.loadRegToReg(0x51)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x51: %v", err)
		}
		return 1, 4, nil
	},
	0x52: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // LD D, D
	0x53: func(c *Cpu) (uint16, int, error) { // LD D, E
		err := c.loadRegToReg(0x53)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x53: %v", err)
		}
		return 1, 4, nil
	},
	0x54: func(c *Cpu) (uint16, int, error) { // LD D, H
		err := c.loadRegToReg(0x54)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x54: %v", err)
		}
		return 1, 4, nil
	},
	0x55: func(c *Cpu) (uint16, int, error) { // LD D, L
		err := c.loadRegToReg(0x55)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x55: %v", err)
		}
		return 1, 4, nil
	},
	0x56: func(c *Cpu) (uint16, int, error) { // LD D, [HL]
		err := c.loadRegToReg(0x56)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x56: %v", err)
		}
		return 1, 8, nil
	},
	0x57: func(c *Cpu) (uint16, int, error) { // LD D, A
		err := c.loadRegToReg(0x57)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x57: %v", err)
		}
		return 1, 4, nil
	},
	0x58: func(c *Cpu) (uint16, int, error) { // LD E, B
		err := c.loadRegToReg(0x58)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x58: %v", err)
		}
		return 1, 4, nil
	},
	0x59: func(c *Cpu) (uint16, int, error) { // LD E, C
		err := c.loadRegToReg(0x59)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x59: %v", err)
		}
		return 1, 4, nil
	},
	0x5A: func(c *Cpu) (uint16, int, error) { // LD E, D
		err := c.loadRegToReg(0x5A)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x5A: %v", err)
		}
		return 1, 4, nil
	},
	0x5B: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // LD E, E
	0x5C: func(c *Cpu) (uint16, int, error) { // LD E, H
		err := c.loadRegToReg(0x5C)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x5C: %v", err)
		}
		return 1, 4, nil
	},
	0x5D: func(c *Cpu) (uint16, int, error) { // LD E, L
		err := c.loadRegToReg(0x5D)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x5D: %v", err)
		}
		return 1, 4, nil
	},
	0x5E: func(c *Cpu) (uint16, int, error) { // LD E, [HL]
		err := c.loadRegToReg(0x5E)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x5E: %v", err)
		}
		return 1, 8, nil
	},
	0x5F: func(c *Cpu) (uint16, int, error) { // LD E, A
		err := c.loadRegToReg(0x5F)
		if err != nil {
			return 0, 0, fmt.Errorf("failed opcode 0x5F: %v", err)
		}
		return 1, 4, nil
	},

	0x60: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x61: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x62: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x63: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x64: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x65: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x66: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x67: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x68: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x69: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x6A: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x6B: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x6C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x6D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x6E: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x6F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0x70: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x71: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x72: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x73: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x74: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x75: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x76: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x77: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x78: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x79: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x7A: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x7B: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x7C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x7D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x7E: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x7F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0x80: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x81: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x82: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x83: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x84: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x85: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x86: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x87: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x88: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x89: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x8A: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x8B: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x8C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x8D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x8E: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x8F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0x90: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x91: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x92: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x93: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x94: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x95: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x96: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x97: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x98: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x99: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x9A: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x9B: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x9C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x9D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x9E: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x9F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0xA0: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA1: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA2: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA3: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA4: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA5: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA6: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0xA7: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA8: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xA9: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xAA: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xAB: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xAC: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xAD: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xAE: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0xAF: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0xB0: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB1: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB2: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB3: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB4: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB5: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB6: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0xB7: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB8: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xB9: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xBA: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xBB: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xBC: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xBD: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xBE: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0xBF: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	0xC0: func(c *Cpu) (uint16, int, error) { return 1, 20, nil },
	0xC1: func(c *Cpu) (uint16, int, error) { return 1, 12, nil },
	0xC2: func(c *Cpu) (uint16, int, error) { return 3, 16, nil },
	0xC3: func(c *Cpu) (uint16, int, error) { return 3, 16, nil },
	0xC4: func(c *Cpu) (uint16, int, error) { return 3, 24, nil },
	0xC5: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xC6: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xC7: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xC8: func(c *Cpu) (uint16, int, error) { return 1, 20, nil },
	0xC9: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xCA: func(c *Cpu) (uint16, int, error) { return 3, 16, nil },
	0xCB: func(c *Cpu) (uint16, int, error) { return 2, 4, nil },
	0xCC: func(c *Cpu) (uint16, int, error) { return 3, 24, nil },
	0xCD: func(c *Cpu) (uint16, int, error) { return 3, 24, nil },
	0xCE: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xCF: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },

	0xD0: func(c *Cpu) (uint16, int, error) { return 1, 20, nil },
	0xD1: func(c *Cpu) (uint16, int, error) { return 1, 12, nil },
	0xD2: func(c *Cpu) (uint16, int, error) { return 3, 16, nil },
	0xD3: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xD4: func(c *Cpu) (uint16, int, error) { return 3, 24, nil },
	0xD5: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xD6: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xD7: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xD8: func(c *Cpu) (uint16, int, error) { return 1, 20, nil },
	0xD9: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xDA: func(c *Cpu) (uint16, int, error) { return 3, 16, nil },
	0xDB: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xDC: func(c *Cpu) (uint16, int, error) { return 3, 24, nil },
	0xDD: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xDE: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xDF: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },

	0xE0: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0xE1: func(c *Cpu) (uint16, int, error) { return 1, 12, nil },
	0xE2: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0xE3: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xE4: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xE5: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xE6: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xE7: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xE8: func(c *Cpu) (uint16, int, error) { return 2, 16, nil },
	0xE9: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xEA: func(c *Cpu) (uint16, int, error) { return 3, 16, nil },
	0xEB: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xEC: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xED: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xEE: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xEF: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },

	0xF0: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0xF1: func(c *Cpu) (uint16, int, error) { return 1, 12, nil },
	0xF2: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0xF3: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xF4: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xF5: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xF6: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xF7: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
	0xF8: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0xF9: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0xFA: func(c *Cpu) (uint16, int, error) { return 3, 16, nil },
	0xFB: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0xFC: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xFD: func(c *Cpu) (uint16, int, error) { return 1, 0, nil },
	0xFE: func(c *Cpu) (uint16, int, error) { return 2, 8, nil },
	0xFF: func(c *Cpu) (uint16, int, error) { return 1, 16, nil },
}
