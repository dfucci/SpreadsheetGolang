package main

import "fmt"

type Spreadsheet struct {
  cells map[string]int
}

func NewSpreadsheet() *Spreadsheet{
  spreadsheet := new(Spreadsheet)
  spreadsheet.cells = make(map[string]int)
  // spreadsheet.cells["A1"] = 2
  return spreadsheet
}

type err interface {
  Error() string
}
type Methods interface {
  Set(cell string, value int)
  Get(cell string) int
}

func (s *Spreadsheet) Set(cell string, value int) {
  s.cells[cell] = value
}

func (s *Spreadsheet) Get(cell string) int {
  elem := s.cells[cell]
  return elem
}

// Using this for debugging purpose
func main(){
  s := NewSpreadsheet()
  s.Set("A1", 8)
  fmt.Printf("%d", s.Get("A1"))
}
