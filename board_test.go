package sudokugo_test

import (
	"testing"

	"github.com/debug-ing/sudokugo"
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
