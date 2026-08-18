package cpu

import (
	"github.com/celineleedle/celintendo/gb/internal/bus"
)

type Cpu struct {
	registers Registers

	bus *bus.Bus
}

func (c *Cpu) readCycle(address uint16) (byte, error) {
	// advance time
	// c.bus.Tick(4) // 4 t-cycles / 1 m-cycle for a read

	return c.bus.ReadByteAt(address)
}

func (c *Cpu) writeCycle(address uint16, value byte) error {
	// advance time
	// c.bus.Tick(4) // 4 t-cycles / 1 m-cycle for a write

	return c.bus.WriteByteAt(address, value)
}

func (c *Cpu) readWordCycle(address uint16) (uint16, error) {
	lowByte, err := c.readCycle(address) // readcycle handles advancing ticks
	if err != nil {
		return 0, err
	}

	highByte, err := c.readCycle(address + 1)
	if err != nil {
		return 0, err
	}

	return uint16(highByte)<<8 | uint16(lowByte), nil
}

func (c *Cpu) step() error {
	// fetch opcode at PC
	pc := c.registers.PC
	opcode, err := c.readCycle(pc)
	if err != nil {
		return err
	}
	c.registers.PC++ // increment PC to point to next instruction

	err = c.execute(opcode)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cpu) Run() {
	// Loop - FETCH, DECODE, EXECUTE, count cycles then repeat

}
