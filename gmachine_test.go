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
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
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

func TestINCA(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.Memory[0] = gmachine.OpINCA
	gm.Run()
	if gm.A != 1 {
		t.Errorf("accumulator should be 1, got %d", gm.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.A = 2
	gm.Memory[0] = gmachine.OpDECA
	gm.Run()
	if gm.A != 1 {
		t.Errorf("accumulator should be 1, got %d", gm.A)
	}
}

func Test3DEC2(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.A = 3
	for i := 0; i < 2; i++ {
		gm.Memory[i] = gmachine.OpDECA
	}
	gm.Run()
	if gm.A != 1 {
		t.Errorf("2-3 should be 1, got %d", gm.A)
	}
}
