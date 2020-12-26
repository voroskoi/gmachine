// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	// OpHALT terminates Gmachine
	OpHALT uint64 = iota
	// OpNOOP does nothing
	OpNOOP
	// OpINCA increments A
	OpINCA
	// OpDECA decrements A
	OpDECA
	// OpSETA set A to based on current P
	OpSETA
)

// Gmachine represents a G-machine.
type Gmachine struct {
	P      uint64 // P is the program counter, holds the actual program number
	A      uint64 // A is the accumulator to store the restult of last operation
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
		case OpHALT:
			return
		case OpNOOP:
			// do nothing
		case OpINCA:
			gm.A++
		case OpDECA:
			gm.A--
		case OpSETA:
			gm.A = gm.Memory[gm.P]
			gm.P++
		default:
			// not used...
		}
	}
}

// RunProgram loads the program to Gmachine memory, then starts it by calling Run().
func (gm *Gmachine) RunProgram(mem []uint64) {
	for k, v := range mem {
		gm.Memory[k] = v
	}
	gm.Run()
}
