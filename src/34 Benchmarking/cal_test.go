// calculator_test.go
package main

import (
   "testing"
)

// Benchmark function
func BenchmarkAdd(b *testing.B) {
   // b.N is a value provided by the testing framework, representing the number of iterations to run
   for i := 0; i < b.N; i++ {
       // Call the function you want to benchmark
       _ = Add(3, 4)
   }
}