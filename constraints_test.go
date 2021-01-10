package constraints

import (
	"testing"
)

var benchfunc = func(b *testing.B) {
	b.SetBytes(512)
	for i := 0; i < b.N; i++ {
		memAllocRepro(nil)
	}
}

func TestMaxAllocations(t *testing.T) {
	br1 := NewBR().Append(benchfunc, WithMaxAllocs(0), WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail WithMaxAllocs(0)")
	}

	br2 := NewBR().Append(benchfunc, WithMaxAllocs(5), WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail WithMaxAllocs(5)")
	}
}

func TestMaxNsPerOp(t *testing.T) {
	br1 := NewBR().Append(benchfunc, WithMaxNsPerOp(5), WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail WithMaxNsPerOp(5)")
	}

	br2 := NewBR().Append(benchfunc, WithMaxNsPerOp(1_000_000), WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail WithMaxNsPerOp(1_000_000)")
	}
}

func TestMaxAllocatedBytes(t *testing.T) {
	br1 := NewBR().Append(benchfunc, WithMaxAllocatedBytes(1), WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail WithMaxAllocatedBytes(1)")
	}

	br2 := NewBR().Append(benchfunc, WithMaxAllocatedBytes(5000), WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail WithMaxAllocatedBytes(5000)")
	}
}

func TestMinMaxMBPerSec(t *testing.T) {
	br1 := NewBR().Append(benchfunc, WithMaxMBPerSec(1), WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail WithMaxMBPerSec(1)")
	}

	br2 := NewBR().Append(benchfunc, WithMaxMBPerSec(1_000_000), WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail WithMaxMBPerSec(1_000_000)")
	}

	br3 := NewBR().Append(benchfunc, WithMinMBPerSec(1_000_000), WithVerbose())
	_, err3 := br3.Run()
	if err3 == nil {
		t.Errorf("Expected to fail WithMinMBPerSec(1_000_000)")
	}

	br4 := NewBR().Append(benchfunc, WithMinMBPerSec(1), WithVerbose())
	_, err4 := br4.Run()
	if err4 != nil {
		t.Errorf("Not expected to fail WithMinMBPerSec(1)")
	}
}

func memAllocRepro(values []int) *[]int {
	for {
		break
	}
	return &values
}
