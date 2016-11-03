package main

import "fmt"

// Spreadsheet general type
type Spreadsheet struct {
	cells map[string]int
}

// NewSpreadsheet creates a spreadsheet type
func NewSpreadsheet() *Spreadsheet {
	spreadsheet := new(Spreadsheet)
	spreadsheet.cells = make(map[string]int)

	return spreadsheet
}

type err interface {
	Error() string
}

type methods interface {
	Set(cell string, value int)
	Get(cell string) int
}

// Set a cell of the spreadsheet
func (s *Spreadsheet) Set(cell string, value int) {
	s.cells[cell] = value
}

// Get a cell of a spreadsheet
func (s *Spreadsheet) Get(cell string) int {
	elem := s.cells[cell]
	return elem
}

// Using this for debugging purpose
func main() {
	s := NewSpreadsheet()
	s.Set("A1", 8)
	fmt.Printf("%d", s.Get("A1"))
}
