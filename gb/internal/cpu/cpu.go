package cpu

import (
	"log"

	"github.com/celineleedle/celintendo/gb/internal/bus"
)

type Cpu struct {
	registers Registers

	bus *bus.Bus
}

func (c *Cpu) ReadCycle(address uint16) (byte, error) {
	// advance time
	// c.bus.Tick(4) // 4 t-cycles / 1 m-cycle for a read

	return c.bus.ReadByteAt(address)
}

func (c *Cpu) WriteCycle(address uint16, value byte) error {
	// advance time
	// c.bus.Tick(4) // 4 t-cycles / 1 m-cycle for a write

	return c.bus.WriteByteAt(address, value)
}

func (c *Cpu) Step() int {
	// fetch opcode at PC
	opcode, err := c.ReadCycle(c.registers.PC)
	if err != nil {
		log.Fatalf("Error reading opcode at PC: 0x%04X: %v", c.registers.PC, err)
	}
	c.registers.PC++ // increment PC to point to next instruction

	err = c.execute(opcode)
	if err != nil {
		log.Fatalf("Error executing opcode: 0x%02X at PC: 0x%04X: %v", opcode, c.registers.PC, err)
	}
	// TODO: add a debug boolean flag and gate these logs behind it.
	log.Printf("Cycles executed: %d", cycles)
	log.Printf("PC increment: %d", pc)

	c.registers.PC += pc

	return cycles
}

func (c *Cpu) Run() {
	// Loop - FETCH, DECODE, EXECUTE, count cycles then repeat

}
