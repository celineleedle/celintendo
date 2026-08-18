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
	zeroBitIndex        = 7
	subtractionBitIndex = 6
	halfCarryBitIndex   = 5
	carryBitIndex       = 4
)

func (r *Registers) getFlag(bitIndex int) bool {
	return (r.F&(uint8(1)<<bitIndex))>>bitIndex == 1
}

func (r *Registers) setFlag(bitIndex int, value bool) {
	if value {
		r.F |= uint8(1) << bitIndex
	} else {
		r.F &= ^(uint8(1) << bitIndex)
	}
}

func (r *Registers) setAF(value uint16) {
	r.A = uint8(value >> 8)
	r.F = uint8(value & 0xF0) // flag register uses bits 7, 6, 5, 4 only
}

func (r *Registers) setBC(value uint16) {
	r.B = uint8(value >> 8)
	r.C = uint8(value & 0xFF)
}

func (r *Registers) setDE(value uint16) {
	r.D = uint8(value >> 8)
	r.E = uint8(value & 0xFF)
}

func (r *Registers) setHL(value uint16) {
	r.H = uint8(value >> 8)
	r.L = uint8(value & 0xFF)
}
