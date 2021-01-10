package constraints

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

type constraint struct {
	benchfunc            func(b *testing.B)
	verbose              bool
	maxAllocedBytesPerOp *int64
	maxAllocsPerOp       *int64
	maxNsPerOp           *int64
	maxMBPerSec          *float64
	minMBPerSec          *float64
}

type BenchmarkRunner struct {
	c []constraint
}

func NewBR() *BenchmarkRunner {
	return &BenchmarkRunner{[]constraint{}}
}

type optionFunc func(*constraint)

func WithMaxAllocs(i int64) optionFunc {
	return func(c *constraint) {
		c.maxAllocsPerOp = &i
	}
}

func WithMaxAllocatedBytes(i int64) optionFunc {
	return func(c *constraint) {
		c.maxAllocedBytesPerOp = &i
	}
}

func WithMaxNsPerOp(i int64) optionFunc {
	return func(c *constraint) {
		c.maxNsPerOp = &i
	}
}

func WithMaxMBPerSec(f float64) optionFunc {
	return func(c *constraint) {
		c.maxMBPerSec = &f
	}
}

func WithMinMBPerSec(f float64) optionFunc {
	return func(c *constraint) {
		c.minMBPerSec = &f
	}
}

func WithVerbose() optionFunc {
	return func(c *constraint) {
		c.verbose = true
	}
}

func (br *BenchmarkRunner) Append(rr func(b *testing.B), opts ...optionFunc) *BenchmarkRunner {
	c := &constraint{benchfunc: rr}

	for _, opt := range opts {
		opt(c)
	}

	cs := br.c
	cs = append(cs, *c)
	return &BenchmarkRunner{cs}
}

func (br *BenchmarkRunner) Run() (bool, error) {
	var err error
	for _, c := range br.c {
		benchfuncName := runtime.FuncForPC(reflect.ValueOf(c.benchfunc).Pointer()).Name()
		if c.verbose {
			fmt.Printf("Executing : %s\n", benchfuncName)
		}
		res := testing.Benchmark(c.benchfunc)
		if c.verbose {
			fmt.Printf("%s\t%s\n", res.String(), res.MemString())
		}

		if c.maxAllocsPerOp != nil && res.AllocsPerOp() > *c.maxAllocsPerOp {
			err = fmt.Errorf("%w; %s : exceeded max allocations per op", err, benchfuncName)
		}

		if c.maxAllocedBytesPerOp != nil && res.AllocedBytesPerOp() > *c.maxAllocedBytesPerOp {
			err = fmt.Errorf("%w; %s : exceeded max allocated bytes per op", err, benchfuncName)
		}

		if c.maxNsPerOp != nil && res.NsPerOp() > *c.maxNsPerOp {
			err = fmt.Errorf("%w; %s : exceeded max ns per op", err, benchfuncName)
		}

		if c.maxMBPerSec != nil && mbPerSec(res) > *c.maxMBPerSec {
			err = fmt.Errorf("%w; %s : exceeded max MB/s", err, benchfuncName)
		}

		if c.minMBPerSec != nil && mbPerSec(res) < *c.minMBPerSec {
			err = fmt.Errorf("%w; %s : was below min MB/s", err, benchfuncName)
		}
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// mbPerSec returns the "MB/s" metric.
func mbPerSec(r testing.BenchmarkResult) float64 {
	if v, ok := r.Extra["MB/s"]; ok {
		return v
	}
	if r.Bytes <= 0 || r.T <= 0 || r.N <= 0 {
		return 0
	}

	return (float64(r.Bytes) * float64(r.N) / 1e6) / r.T.Seconds()
}
