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
	if accessible := m.accessible(address); !accessible {
		return 0x00, fmt.Errorf("attempted to read from inaccessible memory at address 0x%04X", address)
	}

	return m.data[address], nil
}

func (m *Mmu) ReadWordAt(address uint16) (uint16, error) {
	if accessible := m.accessible(address); !accessible {
		return 0x0000, fmt.Errorf("attempted to read from inaccessible memory at address 0x%04X", address)
	} else if accessible := m.accessible(address + 1); !accessible {
		return 0x0000, fmt.Errorf("attempted to read from inaccessible memory at address 0x%04X", address+1)
	}

	lowByte := m.data[address]
	highByte := m.data[address+1]

	return uint16(highByte)<<8 | uint16(lowByte), nil
}

func (m *Mmu) WriteByteAt(address uint16, value byte) error {
	if accessible := m.accessible(address); !accessible {
		return fmt.Errorf("attempted to write to inaccessible memory at address 0x%04X", address)
	}

	m.data[address] = value
	return nil
}

func (m *Mmu) accessible(address uint16) bool {
	if address >= 0xE000 && address <= 0xFDFF {
		return false
	} else if address >= 0xFEA0 && address <= 0xFEFF {
		return false
	}

	return true
}
