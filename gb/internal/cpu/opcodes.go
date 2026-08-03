package cpu

// opcodeFuncMap maps each 1-byte opcode to a handler that executes it and
// returns (pc, cycles):
//   - pc     = the instruction's length in bytes (how far to advance PC)
//   - cycles = the number of T-cycles (clock ticks) the instruction consumed
var opcodeFuncMap = map[uint8]func(*Cpu) (pc uint16, cycles int, err error){
	// 0x00 - 0x0F
	0x00: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // NOP
	0x01: func(c *Cpu) (uint16, int, error) { // LD BC, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, err
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

	// 0x10 - 0x1F
	0x10: func(c *Cpu) (uint16, int, error) { return 2, 4, nil },
	0x11: func(c *Cpu) (uint16, int, error) { // LD DE, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, err
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

	// 0x20 - 0x2F
	0x20: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x21: func(c *Cpu) (uint16, int, error) { // LD HL, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, err
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

	// 0x30 - 0x3F
	0x30: func(c *Cpu) (uint16, int, error) { return 2, 12, nil },
	0x31: func(c *Cpu) (uint16, int, error) { // LD SP, n16
		n16, err := c.mmu.ReadWordAt(c.registers.PC + 1)
		if err != nil {
			return 0, 0, err
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

	// 0x40
	0x40: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // LD B, B
	0x41: func(c *Cpu) (uint16, int, error) { // LD B, C
		c.registers.B = c.registers.C
		return 1, 4, nil
	},
	0x42: func(c *Cpu) (uint16, int, error) { // LD B, D
		c.registers.B = c.registers.D
		return 1, 4, nil
	},
	0x43: func(c *Cpu) (uint16, int, error) { // LD B, E
		c.registers.B = c.registers.E
		return 1, 4, nil
	},
	0x44: func(c *Cpu) (uint16, int, error) { // LD B, H
		c.registers.B = c.registers.H
		return 1, 4, nil
	},
	0x45: func(c *Cpu) (uint16, int, error) { // LD B, L
		c.registers.B = c.registers.L
		return 1, 4, nil
	},
	0x46: func(c *Cpu) (uint16, int, error) { // LD B, [HL]
		value, err := c.mmu.ReadByteAt(c.registers.HL())
		if err != nil {
			return 0, 0, err
		}
		c.registers.B = value
		return 1, 8, nil
	},
	0x47: func(c *Cpu) (uint16, int, error) { // LD B, A
		c.registers.B = c.registers.A
		return 1, 4, nil
	},
	0x48: func(c *Cpu) (uint16, int, error) { // LD C, B
		c.registers.C = c.registers.B
		return 1, 4, nil
	},
	0x49: func(c *Cpu) (uint16, int, error) { return 1, 4, nil }, // LD C, C
	0x4A: func(c *Cpu) (uint16, int, error) { // LD C, D
		c.registers.C = c.registers.D
		return 1, 4, nil
	},
	0x4B: func(c *Cpu) (uint16, int, error) { // LD C, E
		c.registers.C = c.registers.E
		return 1, 4, nil
	},
	0x4C: func(c *Cpu) (uint16, int, error) { // LD C, H
		c.registers.C = c.registers.H
		return 1, 4, nil
	},
	0x4D: func(c *Cpu) (uint16, int, error) { // LD C, L
		c.registers.C = c.registers.L
		return 1, 4, nil
	},
	0x4E: func(c *Cpu) (uint16, int, error) { // LD C, [HL]
		value, err := c.mmu.ReadByteAt(c.registers.HL())
		if err != nil {
			return 0, 0, err
		}
		c.registers.C = value
		return 1, 8, nil
	},
	0x4F: func(c *Cpu) (uint16, int, error) { // LD C, A
		c.registers.C = c.registers.A
		return 1, 4, nil
	},

	// 0x50 - 0x5F  (LD D,r / LD E,r)  — all 1 byte
	0x50: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x51: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x52: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x53: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x54: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x55: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x56: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x57: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x58: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x59: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x5A: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x5B: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x5C: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x5D: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },
	0x5E: func(c *Cpu) (uint16, int, error) { return 1, 8, nil },
	0x5F: func(c *Cpu) (uint16, int, error) { return 1, 4, nil },

	// 0x60 - 0x6F  (LD H,r / LD L,r)  — all 1 byte
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

	// 0x70 - 0x7F  (LD (HL),r / HALT / LD A,r)  — all 1 byte
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

	// 0x80 - 0x8F  (ADD A,r / ADC A,r)  — all 1 byte
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

	// 0x90 - 0x9F  (SUB r / SBC A,r)  — all 1 byte
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

	// 0xA0 - 0xAF  (AND r / XOR r)  — all 1 byte
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

	// 0xB0 - 0xBF  (OR r / CP r)  — all 1 byte
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

	// 0xC0 - 0xCF
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

	// 0xD0 - 0xDF
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

	// 0xE0 - 0xEF
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

	// 0xF0 - 0xFF
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
