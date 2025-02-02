package sudokugo_test

import (
	"testing"

	"github.com/debug-ing/sudokugo"
)

func TestRemoveNumbers(t *testing.T) {
	s := sudokugo.NewSudoku().SetLevel(sudokugo.Easy).InitBoard().RemoveNumbers().GetBoard()
	same := false
	for i := range s {
		for j := range s[i] {
			if s[i][j] == 0 {
				same = true
				break
			}
		}
	}
	if !same {
		t.Errorf("Expected board to be modified after removing numbers")
	}
}
