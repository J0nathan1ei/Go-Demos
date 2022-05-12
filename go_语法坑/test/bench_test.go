// hello project main.go
package test

import (
	"testing"
)

// go test -bench=. -run=^Benchmark$
func Benchmark(t *testing.B) {
	for i := 0; i < t.N; i++ {
	}
}
