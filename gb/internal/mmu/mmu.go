package mmu

import "fmt"

// memory map
// 0000-3FFF | 16 KiB ROM bank 00 | From cartridge, usually a fixed bank
// 4000-7FFF | 16 KiB ROM Bank 01-NN | From cartridge, switchable bank via mapper (if any)
// 8000-9FFF | 8 KiB Video RAM (VRAM) | In CGB mode, switchable bank 0/1
// A000-BFFF | 8 KiB External RAM | From cartridge, switchable bank if any
// C000-CFFF | 4 KiB Work RAM (WRAM)
// D000-DFFF | 4 KiB Work RAM (WRAM) | In CGB mode, switchable bank 1-7
// E000-FDFF | Echo RAM (mirror of C000-DDFF) | Nintendo says use of this area is prohibited.
// FE00-FE9F | Object attribute memory (OAM)
// FEA0-FEFF | Not Usable | Nintendo says use of this area is prohibited.
// FF00-FF7F | I/O Registers
// FF80-FFFE | High RAM (HRAM)
// FFFF-FFFF | Interrupt Enable register (IE)
type Mmu struct {
	data [0x10000]byte
}

func (m *Mmu) ReadByteAt(address uint16) (byte, error) {
	if address >= 0xE000 && address <= 0xFDFF {
		return 0x00, fmt.Errorf("attempted to read from echo RAM at address 0x%04X", address)
	} else if address >= 0xFEA0 && address <= 0xFEFF {
		return 0x00, fmt.Errorf("attempted to read from unusable memory at address 0x%04X", address)
	}

	return m.data[address], nil
}

func (m *Mmu) ReadNext8(address uint16) (byte, error) {
	// is this technically the correct check
	// since it's the next byte
	if address >= 0xE000 && address <= 0xFDFF {
		return 0x00, fmt.Errorf("attempted to read from echo RAM at address 0x%04X", address)
		// same question below
	} else if address >= 0xFEA0 && address <= 0xFEFF {
		return 0x00, fmt.Errorf("attempted to read from unusable memory at address 0x%04X", address)
	}

	return m.data[address+1], nil
}

func (m *Mmu) ReadNext16(address uint16) (uint16, error) {
	low, err := m.ReadNext8(address)
	if err != nil {
		return 0, err
	}

	high, err := m.ReadNext8(address + 1)
	if err != nil {
		return 0, err
	}

	return uint16(high)<<8 | uint16(low), nil
}

func (m *Mmu) WriteByteAt(address uint16, value byte) error {
	if address >= 0xE000 && address <= 0xFEFF {
		return fmt.Errorf("attempted to write to echo RAM at address 0x%04X", address)
	} else if address >= 0xFEA0 && address <= 0xFEFF {
		return fmt.Errorf("attempted to write to unusable memory at address 0x%04X", address)
	}

	m.data[address] = value
	return nil
}
