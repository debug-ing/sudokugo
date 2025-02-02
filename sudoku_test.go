package sudokugo_test

import (
	"sudokugo"
	"testing"
)

func TestNewSudoku(t *testing.T) {
	s := sudokugo.NewSudoku()
	if s == nil {
		t.Errorf("Expected Sudoku instance, got nil")
	}
	if len(s.Get().GetBoard()) != 9 {
		t.Errorf("Expected size 9, got %d", s.Get().GetBoard())
	}
}
