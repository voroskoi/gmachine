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
	wantP := uint64(1)
	if gm.P != wantP {
		t.Errorf("program counter should be %d, got: %d", wantP, gm.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.RunProgram([]uint64{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})
	wantP := uint64(2)
	if gm.P != wantP {
		t.Errorf("program counter should be %d, got: %d", wantP, gm.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.Memory[0] = gmachine.OpINCA
	gm.Run()
	wantA := uint64(1)
	if gm.A != wantA {
		t.Errorf("accumulator (0++) should be %d, got %d", wantA, gm.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.A = 2
	gm.Memory[0] = gmachine.OpDECA
	gm.Run()
	wantA := uint64(1)
	if gm.A != 1 {
		t.Errorf("accumulator (2--) should be %d, got %d", wantA, gm.A)
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
	wantA := uint64(1)
	if gm.A != wantA {
		t.Errorf("2-3 should be %d, got %d", wantA, gm.A)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	gm.Memory[0] = gmachine.OpSETA
	gm.Memory[1] = 66
	wantA := uint64(66)
	wantP := uint64(3) // 0: OpSETA, 1: 66, 2: actual, 3: gm.P++
	gm.Run()
	if gm.A != wantA {
		t.Errorf("accumulator should be %d, got %d", wantA, gm.A)
	}
	if gm.P != wantP {
		t.Errorf("program counter should be %d, got %d", wantP, gm.P)
	}
}

var subtractTestcases = []struct {
	base uint64
	want uint64
}{
	{2, 0},
	{5, 3},
	{7, 5},
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	gm := gmachine.New()
	for _, tc := range subtractTestcases {
		gm.Memory[0] = gmachine.OpSETA
		gm.Memory[1] = tc.base
		gm.Memory[2] = gmachine.OpDECA
		gm.Memory[3] = gmachine.OpDECA
		gm.Run()
		wantA := tc.want
		if gm.A != wantA {
			t.Errorf("Substraction result should be %d, got %d", gm.A, wantA)
		}
		gm.P = 0
	}
}
