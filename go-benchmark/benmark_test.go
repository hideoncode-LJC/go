package benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for n := 0 ; n < b.N ; n ++ {
		fib(30)
	}
	fmt.Println("b.N", b.N)
}