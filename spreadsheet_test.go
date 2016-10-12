package main

import "testing"
func TestSingleCell(t *testing.T){
  s := NewSpreadsheet()
  s.Set("A1", 2)
  actual := s.Get("A1")
  expected := 2
  if actual != expected {
    t.Errorf("Test failed, expected: %d, got %d", expected, actual)
  }
}

