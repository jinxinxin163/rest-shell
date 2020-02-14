// one_test.go
package test

import (
	"testing"
)

//
func BenchmarkTest2(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
