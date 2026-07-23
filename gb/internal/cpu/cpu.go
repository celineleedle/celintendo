package cpu

type Registers struct {
	A, F uint8 // accumulator & flags
	B, C uint8
	D, E uint8
	H, L uint8
	SP   uint16 // stack pointer
	PC   uint16 // program counter
}

func (r *Registers) AF() uint16 { return uint16(r.A)<<8 | uint16(r.F) }
func (r *Registers) BC() uint16 { return uint16(r.B)<<8 | uint16(r.C) }
func (r *Registers) DE() uint16 { return uint16(r.D)<<8 | uint16(r.E) }
func (r *Registers) HL() uint16 { return uint16(r.H)<<8 | uint16(r.L) }

const (
	ZeroFlag        uint8 = 1 << 7
	SubtractionFlag uint8 = 1 << 6
	HalfCarryFlag   uint8 = 1 << 5
	CarryFlag       uint8 = 1 << 4
)
