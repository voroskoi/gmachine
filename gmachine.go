// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

// Gmachine represents a G-machine.
type Gmachine struct {
	P      uint64
	Memory []uint64
}

// New creates and initializes a new G-machine.
func New() *Gmachine {
	gm := Gmachine{}
	gm.Memory = make([]uint64, DefaultMemSize)
	return &gm
}
