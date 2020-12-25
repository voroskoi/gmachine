// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	// HALT terminates Gmachine
	HALT uint64 = iota
)

// Gmachine represents a G-machine.
type Gmachine struct {
	P      uint64 // P is the program counter, holds the actual program number
	Memory []uint64
}

// New creates and initializes a new G-machine.
func New() *Gmachine {
	gm := Gmachine{}
	gm.Memory = make([]uint64, DefaultMemSize)
	return &gm
}

// Run starts Gmachine.
func (gm *Gmachine) Run() {
	for gm.P < DefaultMemSize {
		gm.P++
		switch gm.Memory[gm.P-1] {
		case HALT:
			return
		default:
		}
	}
}
