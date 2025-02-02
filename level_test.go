package sudokugo_test

import (
	"sudokugo"
	"testing"
)

func TestSetAndGetLevel(t *testing.T) {
	s := sudokugo.NewSudoku()
	s.SetLevel(sudokugo.Hard)
	if s.GetLevel() != "hard" {
		t.Errorf("Expected level 'hard', got %s", s.GetLevel())
	}
}

func TestGetAllLevel(t *testing.T) {
	s := sudokugo.NewSudoku()
	all := s.GetAllLevels()
	if len(all) == 0 {
		t.Errorf("Expected more than 0 levels, got %d", len(all))
	}
}
