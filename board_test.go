package sudokugo_test

import (
	"sudokugo"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	s := sudokugo.NewSudoku().SetLevel(sudokugo.Easy).InitBoard()
	s.RemoveNumbers()
	if s.GetStatusBoard() {
		t.Errorf("Expected board to be incomplete after removing numbers")
	}
	s.Solve()
	if !s.GetStatusBoard() {
		t.Errorf("Expected board to be solved")
	}
}
