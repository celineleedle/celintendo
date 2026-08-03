package cpu

import (
	"log"
)

type Cpu struct {
	registers Registers

	mmu Mmu
}

type Mmu interface {
	ReadByteAt(address uint16) (byte, error)
	WriteByteAt(address uint16, value byte) error
}

func (c *Cpu) Step() int {
	opcode, err := c.mmu.ReadByteAt(c.registers.PC)
	if err != nil {
		log.Fatalf("Error reading opcode at PC: 0x%04X: %v", c.registers.PC, err)
	}

	opcodeFunc, exists := opcodeFuncMap[opcode]
	if !exists {
		log.Fatalf("Unknown opcode: 0x%02X at PC: 0x%04X", opcode, c.registers.PC)
	}

	pc, cycles, err := opcodeFunc(c)
	if err != nil {
		log.Fatalf("Error executing opcode: 0x%02X at PC: 0x%04X: %v", opcode, c.registers.PC, err)
	}
	log.Printf("Cycles executed: %d", cycles)

	c.registers.PC += pc

	return cycles
}

func (c *Cpu) Run() {
	// Loop - FETCH, DECODE, EXECUTE, count cycles then repeat

}
