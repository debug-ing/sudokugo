package sudokugo

import "math/rand/v2"

// Level represents the level of the sudoku puzzle
type Level string

const (
	Easy    Level = "easy"
	Medium  Level = "medium"
	Hard    Level = "hard"
	Expert  Level = "expert"
	Master  Level = "master"
	Extreme Level = "extreme"
)

// GetAllLevels returns all levels
func (*Sudoku) GetAllLevels() map[string]int {
	return map[string]int{
		"easy":    rand.IntN(6) + 28,
		"normal":  rand.IntN(10) + 30,
		"hard":    rand.IntN(9) + 38,
		"expert":  rand.IntN(7) + 50,
		"master":  rand.IntN(7) + 55,
		"extreme": rand.IntN(10) + 60,
	}
}

// GetLevel returns the level of the sudoku puzzle
func (s *Sudoku) GetLevel() string {
	return string(s.Level)
}

// SetLevel sets the level of the sudoku puzzle
func (s *Sudoku) SetLevel(level Level) *Sudoku {
	s.Level = level
	return s
}
