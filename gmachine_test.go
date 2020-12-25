package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHALT(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.Run()
	if gm.P != 1 {
		t.Errorf("program counter should be 1, got: %d", gm.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.Memory[0] = gmachine.OpNOOP
	gm.Run()
	if gm.P != 2 {
		t.Errorf("program counter should be 2, got: %d", gm.P)
	}
}
