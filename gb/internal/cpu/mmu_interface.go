package cpu

type Mmu interface {
	ReadByteAt(address uint16) (byte, error)
	ReadNext8(address uint16) (byte, error)
	ReadNext16(address uint16) (uint16, error)

	WriteByteAt(address uint16, value byte) error
}
