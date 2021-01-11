package constraints_test

import (
	"testing"

	c "github.com/tpaschalis/go-benchmark-constraints"
)

var benchfunc = func(b *testing.B) {
	b.SetBytes(512)
	for i := 0; i < b.N; i++ {
		memAllocRepro(nil)
	}
}

func TestMaxAllocations(t *testing.T) {
	br1 := c.NewBR().Append(benchfunc, c.WithMaxAllocs(0), c.WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail c.WithMaxAllocs(0)")
	}

	br2 := c.NewBR().Append(benchfunc, c.WithMaxAllocs(5), c.WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail c.WithMaxAllocs(5)")
	}
}

func TestMaxNsPerOp(t *testing.T) {
	br1 := c.NewBR().Append(benchfunc, c.WithMaxNsPerOp(5), c.WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail c.WithMaxNsPerOp(5)")
	}

	br2 := c.NewBR().Append(benchfunc, c.WithMaxNsPerOp(1_000_000), c.WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail c.WithMaxNsPerOp(1_000_000)")
	}
}

func TestMaxAllocatedBytes(t *testing.T) {
	br1 := c.NewBR().Append(benchfunc, c.WithMaxAllocatedBytes(1), c.WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail c.WithMaxAllocatedBytes(1)")
	}

	br2 := c.NewBR().Append(benchfunc, c.WithMaxAllocatedBytes(5000), c.WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail c.WithMaxAllocatedBytes(5000)")
	}
}

func TestMinMaxMBPerSec(t *testing.T) {
	br1 := c.NewBR().Append(benchfunc, c.WithMaxMBPerSec(1), c.WithVerbose())
	_, err1 := br1.Run()
	if err1 == nil {
		t.Errorf("Expected to fail c.WithMaxMBPerSec(1)")
	}

	br2 := c.NewBR().Append(benchfunc, c.WithMaxMBPerSec(1_000_000), c.WithVerbose())
	_, err2 := br2.Run()
	if err2 != nil {
		t.Errorf("Not expected to fail c.WithMaxMBPerSec(1_000_000)")
	}

	br3 := c.NewBR().Append(benchfunc, c.WithMinMBPerSec(1_000_000), c.WithVerbose())
	_, err3 := br3.Run()
	if err3 == nil {
		t.Errorf("Expected to fail c.WithMinMBPerSec(1_000_000)")
	}

	br4 := c.NewBR().Append(benchfunc, c.WithMinMBPerSec(1), c.WithVerbose())
	_, err4 := br4.Run()
	if err4 != nil {
		t.Errorf("Not expected to fail c.WithMinMBPerSec(1)")
	}
}

func memAllocRepro(values []int) *[]int {
	for {
		break
	}
	return &values
}
