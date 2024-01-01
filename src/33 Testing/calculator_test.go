// calculator_test.go
package main

import (
   "testing"
   "fmt"
)

func TestAdd(t *testing.T) {
   result := Add(3, 4)
   expected := 7
   if result != expected {
      t.Errorf("Expected %d, but got %d", expected, result)
   }
   fmt.Println("Test Done")
}