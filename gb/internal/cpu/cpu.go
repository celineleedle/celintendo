package cpu

import (
	"log"

	"github.com/celineleedle/celintendo/gb/internal/bus"
)

type Cpu struct {
	registers Registers

	bus *bus.Bus
}

func (c *Cpu) Step() int {
	opcode, err := c.bus.ReadByteAt(c.registers.PC)
	if err != nil {
		log.Fatalf("Error reading opcode at PC: 0x%04X: %v", c.registers.PC, err)
	}

	pc, cycles, err := c.handleOpcode(opcode)
	if err != nil {
		log.Fatalf("Error executing opcode: 0x%02X at PC: 0x%04X: %v", opcode, c.registers.PC, err)
	}
	log.Printf("Cycles executed: %d", cycles)
	log.Printf("PC increment: %d", pc)

	c.registers.PC += pc

	return cycles
}

func (c *Cpu) Run() {
	// Loop - FETCH, DECODE, EXECUTE, count cycles then repeat

}
