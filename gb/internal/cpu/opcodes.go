package cpu

// opcodeFuncMap maps each 1-byte opcode to a handler that executes it and
// returns (pc, cycles):
//   - pc     = the instruction's length in bytes (how far to advance PC)
//   - cycles = the number of T-cycles (clock ticks) the instruction consumed
//
// Length rule: 1 byte base, +1 for a d8/a8/r8 operand (=2), +2 for a
// d16/a16 operand (=3).
//
// Cycle convention: T-cycles (NOP = 4). For conditional instructions the
// comment shows "taken/untaken"; the stub returns the taken value for now.
// (The length never changes with the branch — only the cycles do.)
//
// Operand notation: d8 = 8-bit immediate, d16 = 16-bit immediate,
// a8/a16 = 8/16-bit address, r8 = signed 8-bit offset, (HL) = byte in
// memory at address HL. Opcodes marked "illegal" do not exist on hardware.
//
// 0xCB is the prefix byte: a CB instruction is 2 bytes (prefix + sub-opcode).
// When hit, read the next byte and dispatch it through a separate CB-opcode
// table (rotates/shifts/BIT/SET/RES).
var opcodeFuncMap = map[uint8]func(*Cpu) (pc int, cycles int){
	// 0x00 - 0x0F
	0x00: func(c *Cpu) (int, int) { return 1, 4 },  // NOP
	0x01: func(c *Cpu) (int, int) { return 3, 12 }, // LD BC,d16
	0x02: func(c *Cpu) (int, int) { return 1, 8 },  // LD (BC),A
	0x03: func(c *Cpu) (int, int) { return 1, 8 },  // INC BC
	0x04: func(c *Cpu) (int, int) { return 1, 4 },  // INC B
	0x05: func(c *Cpu) (int, int) { return 1, 4 },  // DEC B
	0x06: func(c *Cpu) (int, int) { return 2, 8 },  // LD B,d8
	0x07: func(c *Cpu) (int, int) { return 1, 4 },  // RLCA
	0x08: func(c *Cpu) (int, int) { return 3, 20 }, // LD (a16),SP
	0x09: func(c *Cpu) (int, int) { return 1, 8 },  // ADD HL,BC
	0x0A: func(c *Cpu) (int, int) { return 1, 8 },  // LD A,(BC)
	0x0B: func(c *Cpu) (int, int) { return 1, 8 },  // DEC BC
	0x0C: func(c *Cpu) (int, int) { return 1, 4 },  // INC C
	0x0D: func(c *Cpu) (int, int) { return 1, 4 },  // DEC C
	0x0E: func(c *Cpu) (int, int) { return 2, 8 },  // LD C,d8
	0x0F: func(c *Cpu) (int, int) { return 1, 4 },  // RRCA

	// 0x10 - 0x1F
	0x10: func(c *Cpu) (int, int) { return 2, 4 },  // STOP (encoded as 0x10 0x00, 2 bytes)
	0x11: func(c *Cpu) (int, int) { return 3, 12 }, // LD DE,d16
	0x12: func(c *Cpu) (int, int) { return 1, 8 },  // LD (DE),A
	0x13: func(c *Cpu) (int, int) { return 1, 8 },  // INC DE
	0x14: func(c *Cpu) (int, int) { return 1, 4 },  // INC D
	0x15: func(c *Cpu) (int, int) { return 1, 4 },  // DEC D
	0x16: func(c *Cpu) (int, int) { return 2, 8 },  // LD D,d8
	0x17: func(c *Cpu) (int, int) { return 1, 4 },  // RLA
	0x18: func(c *Cpu) (int, int) { return 2, 12 }, // JR r8
	0x19: func(c *Cpu) (int, int) { return 1, 8 },  // ADD HL,DE
	0x1A: func(c *Cpu) (int, int) { return 1, 8 },  // LD A,(DE)
	0x1B: func(c *Cpu) (int, int) { return 1, 8 },  // DEC DE
	0x1C: func(c *Cpu) (int, int) { return 1, 4 },  // INC E
	0x1D: func(c *Cpu) (int, int) { return 1, 4 },  // DEC E
	0x1E: func(c *Cpu) (int, int) { return 2, 8 },  // LD E,d8
	0x1F: func(c *Cpu) (int, int) { return 1, 4 },  // RRA

	// 0x20 - 0x2F
	0x20: func(c *Cpu) (int, int) { return 2, 12 }, // JR NZ,r8 (12/8)
	0x21: func(c *Cpu) (int, int) { return 3, 12 }, // LD HL,d16
	0x22: func(c *Cpu) (int, int) { return 1, 8 },  // LD (HL+),A
	0x23: func(c *Cpu) (int, int) { return 1, 8 },  // INC HL
	0x24: func(c *Cpu) (int, int) { return 1, 4 },  // INC H
	0x25: func(c *Cpu) (int, int) { return 1, 4 },  // DEC H
	0x26: func(c *Cpu) (int, int) { return 2, 8 },  // LD H,d8
	0x27: func(c *Cpu) (int, int) { return 1, 4 },  // DAA
	0x28: func(c *Cpu) (int, int) { return 2, 12 }, // JR Z,r8 (12/8)
	0x29: func(c *Cpu) (int, int) { return 1, 8 },  // ADD HL,HL
	0x2A: func(c *Cpu) (int, int) { return 1, 8 },  // LD A,(HL+)
	0x2B: func(c *Cpu) (int, int) { return 1, 8 },  // DEC HL
	0x2C: func(c *Cpu) (int, int) { return 1, 4 },  // INC L
	0x2D: func(c *Cpu) (int, int) { return 1, 4 },  // DEC L
	0x2E: func(c *Cpu) (int, int) { return 2, 8 },  // LD L,d8
	0x2F: func(c *Cpu) (int, int) { return 1, 4 },  // CPL

	// 0x30 - 0x3F
	0x30: func(c *Cpu) (int, int) { return 2, 12 }, // JR NC,r8 (12/8)
	0x31: func(c *Cpu) (int, int) { return 3, 12 }, // LD SP,d16
	0x32: func(c *Cpu) (int, int) { return 1, 8 },  // LD (HL-),A
	0x33: func(c *Cpu) (int, int) { return 1, 8 },  // INC SP
	0x34: func(c *Cpu) (int, int) { return 1, 12 }, // INC (HL)
	0x35: func(c *Cpu) (int, int) { return 1, 12 }, // DEC (HL)
	0x36: func(c *Cpu) (int, int) { return 2, 12 }, // LD (HL),d8
	0x37: func(c *Cpu) (int, int) { return 1, 4 },  // SCF
	0x38: func(c *Cpu) (int, int) { return 2, 12 }, // JR C,r8 (12/8)
	0x39: func(c *Cpu) (int, int) { return 1, 8 },  // ADD HL,SP
	0x3A: func(c *Cpu) (int, int) { return 1, 8 },  // LD A,(HL-)
	0x3B: func(c *Cpu) (int, int) { return 1, 8 },  // DEC SP
	0x3C: func(c *Cpu) (int, int) { return 1, 4 },  // INC A
	0x3D: func(c *Cpu) (int, int) { return 1, 4 },  // DEC A
	0x3E: func(c *Cpu) (int, int) { return 2, 8 },  // LD A,d8
	0x3F: func(c *Cpu) (int, int) { return 1, 4 },  // CCF

	// 0x40 - 0x4F  (LD B,r / LD C,r)  — all 1 byte
	0x40: func(c *Cpu) (int, int) { return 1, 4 }, // LD B,B
	0x41: func(c *Cpu) (int, int) { return 1, 4 }, // LD B,C
	0x42: func(c *Cpu) (int, int) { return 1, 4 }, // LD B,D
	0x43: func(c *Cpu) (int, int) { return 1, 4 }, // LD B,E
	0x44: func(c *Cpu) (int, int) { return 1, 4 }, // LD B,H
	0x45: func(c *Cpu) (int, int) { return 1, 4 }, // LD B,L
	0x46: func(c *Cpu) (int, int) { return 1, 8 }, // LD B,(HL)
	0x47: func(c *Cpu) (int, int) { return 1, 4 }, // LD B,A
	0x48: func(c *Cpu) (int, int) { return 1, 4 }, // LD C,B
	0x49: func(c *Cpu) (int, int) { return 1, 4 }, // LD C,C
	0x4A: func(c *Cpu) (int, int) { return 1, 4 }, // LD C,D
	0x4B: func(c *Cpu) (int, int) { return 1, 4 }, // LD C,E
	0x4C: func(c *Cpu) (int, int) { return 1, 4 }, // LD C,H
	0x4D: func(c *Cpu) (int, int) { return 1, 4 }, // LD C,L
	0x4E: func(c *Cpu) (int, int) { return 1, 8 }, // LD C,(HL)
	0x4F: func(c *Cpu) (int, int) { return 1, 4 }, // LD C,A

	// 0x50 - 0x5F  (LD D,r / LD E,r)  — all 1 byte
	0x50: func(c *Cpu) (int, int) { return 1, 4 }, // LD D,B
	0x51: func(c *Cpu) (int, int) { return 1, 4 }, // LD D,C
	0x52: func(c *Cpu) (int, int) { return 1, 4 }, // LD D,D
	0x53: func(c *Cpu) (int, int) { return 1, 4 }, // LD D,E
	0x54: func(c *Cpu) (int, int) { return 1, 4 }, // LD D,H
	0x55: func(c *Cpu) (int, int) { return 1, 4 }, // LD D,L
	0x56: func(c *Cpu) (int, int) { return 1, 8 }, // LD D,(HL)
	0x57: func(c *Cpu) (int, int) { return 1, 4 }, // LD D,A
	0x58: func(c *Cpu) (int, int) { return 1, 4 }, // LD E,B
	0x59: func(c *Cpu) (int, int) { return 1, 4 }, // LD E,C
	0x5A: func(c *Cpu) (int, int) { return 1, 4 }, // LD E,D
	0x5B: func(c *Cpu) (int, int) { return 1, 4 }, // LD E,E
	0x5C: func(c *Cpu) (int, int) { return 1, 4 }, // LD E,H
	0x5D: func(c *Cpu) (int, int) { return 1, 4 }, // LD E,L
	0x5E: func(c *Cpu) (int, int) { return 1, 8 }, // LD E,(HL)
	0x5F: func(c *Cpu) (int, int) { return 1, 4 }, // LD E,A

	// 0x60 - 0x6F  (LD H,r / LD L,r)  — all 1 byte
	0x60: func(c *Cpu) (int, int) { return 1, 4 }, // LD H,B
	0x61: func(c *Cpu) (int, int) { return 1, 4 }, // LD H,C
	0x62: func(c *Cpu) (int, int) { return 1, 4 }, // LD H,D
	0x63: func(c *Cpu) (int, int) { return 1, 4 }, // LD H,E
	0x64: func(c *Cpu) (int, int) { return 1, 4 }, // LD H,H
	0x65: func(c *Cpu) (int, int) { return 1, 4 }, // LD H,L
	0x66: func(c *Cpu) (int, int) { return 1, 8 }, // LD H,(HL)
	0x67: func(c *Cpu) (int, int) { return 1, 4 }, // LD H,A
	0x68: func(c *Cpu) (int, int) { return 1, 4 }, // LD L,B
	0x69: func(c *Cpu) (int, int) { return 1, 4 }, // LD L,C
	0x6A: func(c *Cpu) (int, int) { return 1, 4 }, // LD L,D
	0x6B: func(c *Cpu) (int, int) { return 1, 4 }, // LD L,E
	0x6C: func(c *Cpu) (int, int) { return 1, 4 }, // LD L,H
	0x6D: func(c *Cpu) (int, int) { return 1, 4 }, // LD L,L
	0x6E: func(c *Cpu) (int, int) { return 1, 8 }, // LD L,(HL)
	0x6F: func(c *Cpu) (int, int) { return 1, 4 }, // LD L,A

	// 0x70 - 0x7F  (LD (HL),r / HALT / LD A,r)  — all 1 byte
	0x70: func(c *Cpu) (int, int) { return 1, 8 }, // LD (HL),B
	0x71: func(c *Cpu) (int, int) { return 1, 8 }, // LD (HL),C
	0x72: func(c *Cpu) (int, int) { return 1, 8 }, // LD (HL),D
	0x73: func(c *Cpu) (int, int) { return 1, 8 }, // LD (HL),E
	0x74: func(c *Cpu) (int, int) { return 1, 8 }, // LD (HL),H
	0x75: func(c *Cpu) (int, int) { return 1, 8 }, // LD (HL),L
	0x76: func(c *Cpu) (int, int) { return 1, 4 }, // HALT
	0x77: func(c *Cpu) (int, int) { return 1, 8 }, // LD (HL),A
	0x78: func(c *Cpu) (int, int) { return 1, 4 }, // LD A,B
	0x79: func(c *Cpu) (int, int) { return 1, 4 }, // LD A,C
	0x7A: func(c *Cpu) (int, int) { return 1, 4 }, // LD A,D
	0x7B: func(c *Cpu) (int, int) { return 1, 4 }, // LD A,E
	0x7C: func(c *Cpu) (int, int) { return 1, 4 }, // LD A,H
	0x7D: func(c *Cpu) (int, int) { return 1, 4 }, // LD A,L
	0x7E: func(c *Cpu) (int, int) { return 1, 8 }, // LD A,(HL)
	0x7F: func(c *Cpu) (int, int) { return 1, 4 }, // LD A,A

	// 0x80 - 0x8F  (ADD A,r / ADC A,r)  — all 1 byte
	0x80: func(c *Cpu) (int, int) { return 1, 4 }, // ADD A,B
	0x81: func(c *Cpu) (int, int) { return 1, 4 }, // ADD A,C
	0x82: func(c *Cpu) (int, int) { return 1, 4 }, // ADD A,D
	0x83: func(c *Cpu) (int, int) { return 1, 4 }, // ADD A,E
	0x84: func(c *Cpu) (int, int) { return 1, 4 }, // ADD A,H
	0x85: func(c *Cpu) (int, int) { return 1, 4 }, // ADD A,L
	0x86: func(c *Cpu) (int, int) { return 1, 8 }, // ADD A,(HL)
	0x87: func(c *Cpu) (int, int) { return 1, 4 }, // ADD A,A
	0x88: func(c *Cpu) (int, int) { return 1, 4 }, // ADC A,B
	0x89: func(c *Cpu) (int, int) { return 1, 4 }, // ADC A,C
	0x8A: func(c *Cpu) (int, int) { return 1, 4 }, // ADC A,D
	0x8B: func(c *Cpu) (int, int) { return 1, 4 }, // ADC A,E
	0x8C: func(c *Cpu) (int, int) { return 1, 4 }, // ADC A,H
	0x8D: func(c *Cpu) (int, int) { return 1, 4 }, // ADC A,L
	0x8E: func(c *Cpu) (int, int) { return 1, 8 }, // ADC A,(HL)
	0x8F: func(c *Cpu) (int, int) { return 1, 4 }, // ADC A,A

	// 0x90 - 0x9F  (SUB r / SBC A,r)  — all 1 byte
	0x90: func(c *Cpu) (int, int) { return 1, 4 }, // SUB B
	0x91: func(c *Cpu) (int, int) { return 1, 4 }, // SUB C
	0x92: func(c *Cpu) (int, int) { return 1, 4 }, // SUB D
	0x93: func(c *Cpu) (int, int) { return 1, 4 }, // SUB E
	0x94: func(c *Cpu) (int, int) { return 1, 4 }, // SUB H
	0x95: func(c *Cpu) (int, int) { return 1, 4 }, // SUB L
	0x96: func(c *Cpu) (int, int) { return 1, 8 }, // SUB (HL)
	0x97: func(c *Cpu) (int, int) { return 1, 4 }, // SUB A
	0x98: func(c *Cpu) (int, int) { return 1, 4 }, // SBC A,B
	0x99: func(c *Cpu) (int, int) { return 1, 4 }, // SBC A,C
	0x9A: func(c *Cpu) (int, int) { return 1, 4 }, // SBC A,D
	0x9B: func(c *Cpu) (int, int) { return 1, 4 }, // SBC A,E
	0x9C: func(c *Cpu) (int, int) { return 1, 4 }, // SBC A,H
	0x9D: func(c *Cpu) (int, int) { return 1, 4 }, // SBC A,L
	0x9E: func(c *Cpu) (int, int) { return 1, 8 }, // SBC A,(HL)
	0x9F: func(c *Cpu) (int, int) { return 1, 4 }, // SBC A,A

	// 0xA0 - 0xAF  (AND r / XOR r)  — all 1 byte
	0xA0: func(c *Cpu) (int, int) { return 1, 4 }, // AND B
	0xA1: func(c *Cpu) (int, int) { return 1, 4 }, // AND C
	0xA2: func(c *Cpu) (int, int) { return 1, 4 }, // AND D
	0xA3: func(c *Cpu) (int, int) { return 1, 4 }, // AND E
	0xA4: func(c *Cpu) (int, int) { return 1, 4 }, // AND H
	0xA5: func(c *Cpu) (int, int) { return 1, 4 }, // AND L
	0xA6: func(c *Cpu) (int, int) { return 1, 8 }, // AND (HL)
	0xA7: func(c *Cpu) (int, int) { return 1, 4 }, // AND A
	0xA8: func(c *Cpu) (int, int) { return 1, 4 }, // XOR B
	0xA9: func(c *Cpu) (int, int) { return 1, 4 }, // XOR C
	0xAA: func(c *Cpu) (int, int) { return 1, 4 }, // XOR D
	0xAB: func(c *Cpu) (int, int) { return 1, 4 }, // XOR E
	0xAC: func(c *Cpu) (int, int) { return 1, 4 }, // XOR H
	0xAD: func(c *Cpu) (int, int) { return 1, 4 }, // XOR L
	0xAE: func(c *Cpu) (int, int) { return 1, 8 }, // XOR (HL)
	0xAF: func(c *Cpu) (int, int) { return 1, 4 }, // XOR A

	// 0xB0 - 0xBF  (OR r / CP r)  — all 1 byte
	0xB0: func(c *Cpu) (int, int) { return 1, 4 }, // OR B
	0xB1: func(c *Cpu) (int, int) { return 1, 4 }, // OR C
	0xB2: func(c *Cpu) (int, int) { return 1, 4 }, // OR D
	0xB3: func(c *Cpu) (int, int) { return 1, 4 }, // OR E
	0xB4: func(c *Cpu) (int, int) { return 1, 4 }, // OR H
	0xB5: func(c *Cpu) (int, int) { return 1, 4 }, // OR L
	0xB6: func(c *Cpu) (int, int) { return 1, 8 }, // OR (HL)
	0xB7: func(c *Cpu) (int, int) { return 1, 4 }, // OR A
	0xB8: func(c *Cpu) (int, int) { return 1, 4 }, // CP B
	0xB9: func(c *Cpu) (int, int) { return 1, 4 }, // CP C
	0xBA: func(c *Cpu) (int, int) { return 1, 4 }, // CP D
	0xBB: func(c *Cpu) (int, int) { return 1, 4 }, // CP E
	0xBC: func(c *Cpu) (int, int) { return 1, 4 }, // CP H
	0xBD: func(c *Cpu) (int, int) { return 1, 4 }, // CP L
	0xBE: func(c *Cpu) (int, int) { return 1, 8 }, // CP (HL)
	0xBF: func(c *Cpu) (int, int) { return 1, 4 }, // CP A

	// 0xC0 - 0xCF
	0xC0: func(c *Cpu) (int, int) { return 1, 20 }, // RET NZ (20/8)
	0xC1: func(c *Cpu) (int, int) { return 1, 12 }, // POP BC
	0xC2: func(c *Cpu) (int, int) { return 3, 16 }, // JP NZ,a16 (16/12)
	0xC3: func(c *Cpu) (int, int) { return 3, 16 }, // JP a16
	0xC4: func(c *Cpu) (int, int) { return 3, 24 }, // CALL NZ,a16 (24/12)
	0xC5: func(c *Cpu) (int, int) { return 1, 16 }, // PUSH BC
	0xC6: func(c *Cpu) (int, int) { return 2, 8 },  // ADD A,d8
	0xC7: func(c *Cpu) (int, int) { return 1, 16 }, // RST 00H
	0xC8: func(c *Cpu) (int, int) { return 1, 20 }, // RET Z (20/8)
	0xC9: func(c *Cpu) (int, int) { return 1, 16 }, // RET
	0xCA: func(c *Cpu) (int, int) { return 3, 16 }, // JP Z,a16 (16/12)
	0xCB: func(c *Cpu) (int, int) { return 2, 4 },  // PREFIX CB (read next byte, dispatch CB table)
	0xCC: func(c *Cpu) (int, int) { return 3, 24 }, // CALL Z,a16 (24/12)
	0xCD: func(c *Cpu) (int, int) { return 3, 24 }, // CALL a16
	0xCE: func(c *Cpu) (int, int) { return 2, 8 },  // ADC A,d8
	0xCF: func(c *Cpu) (int, int) { return 1, 16 }, // RST 08H

	// 0xD0 - 0xDF
	0xD0: func(c *Cpu) (int, int) { return 1, 20 }, // RET NC (20/8)
	0xD1: func(c *Cpu) (int, int) { return 1, 12 }, // POP DE
	0xD2: func(c *Cpu) (int, int) { return 3, 16 }, // JP NC,a16 (16/12)
	0xD3: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xD4: func(c *Cpu) (int, int) { return 3, 24 }, // CALL NC,a16 (24/12)
	0xD5: func(c *Cpu) (int, int) { return 1, 16 }, // PUSH DE
	0xD6: func(c *Cpu) (int, int) { return 2, 8 },  // SUB d8
	0xD7: func(c *Cpu) (int, int) { return 1, 16 }, // RST 10H
	0xD8: func(c *Cpu) (int, int) { return 1, 20 }, // RET C (20/8)
	0xD9: func(c *Cpu) (int, int) { return 1, 16 }, // RETI
	0xDA: func(c *Cpu) (int, int) { return 3, 16 }, // JP C,a16 (16/12)
	0xDB: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xDC: func(c *Cpu) (int, int) { return 3, 24 }, // CALL C,a16 (24/12)
	0xDD: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xDE: func(c *Cpu) (int, int) { return 2, 8 },  // SBC A,d8
	0xDF: func(c *Cpu) (int, int) { return 1, 16 }, // RST 18H

	// 0xE0 - 0xEF
	0xE0: func(c *Cpu) (int, int) { return 2, 12 }, // LDH (a8),A
	0xE1: func(c *Cpu) (int, int) { return 1, 12 }, // POP HL
	0xE2: func(c *Cpu) (int, int) { return 1, 8 },  // LD (C),A
	0xE3: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xE4: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xE5: func(c *Cpu) (int, int) { return 1, 16 }, // PUSH HL
	0xE6: func(c *Cpu) (int, int) { return 2, 8 },  // AND d8
	0xE7: func(c *Cpu) (int, int) { return 1, 16 }, // RST 20H
	0xE8: func(c *Cpu) (int, int) { return 2, 16 }, // ADD SP,r8
	0xE9: func(c *Cpu) (int, int) { return 1, 4 },  // JP (HL)
	0xEA: func(c *Cpu) (int, int) { return 3, 16 }, // LD (a16),A
	0xEB: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xEC: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xED: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xEE: func(c *Cpu) (int, int) { return 2, 8 },  // XOR d8
	0xEF: func(c *Cpu) (int, int) { return 1, 16 }, // RST 28H

	// 0xF0 - 0xFF
	0xF0: func(c *Cpu) (int, int) { return 2, 12 }, // LDH A,(a8)
	0xF1: func(c *Cpu) (int, int) { return 1, 12 }, // POP AF
	0xF2: func(c *Cpu) (int, int) { return 1, 8 },  // LD A,(C)
	0xF3: func(c *Cpu) (int, int) { return 1, 4 },  // DI
	0xF4: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xF5: func(c *Cpu) (int, int) { return 1, 16 }, // PUSH AF
	0xF6: func(c *Cpu) (int, int) { return 2, 8 },  // OR d8
	0xF7: func(c *Cpu) (int, int) { return 1, 16 }, // RST 30H
	0xF8: func(c *Cpu) (int, int) { return 2, 12 }, // LD HL,SP+r8
	0xF9: func(c *Cpu) (int, int) { return 1, 8 },  // LD SP,HL
	0xFA: func(c *Cpu) (int, int) { return 3, 16 }, // LD A,(a16)
	0xFB: func(c *Cpu) (int, int) { return 1, 4 },  // EI
	0xFC: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xFD: func(c *Cpu) (int, int) { return 1, 0 },  // illegal
	0xFE: func(c *Cpu) (int, int) { return 2, 8 },  // CP d8
	0xFF: func(c *Cpu) (int, int) { return 1, 16 }, // RST 38H
}
