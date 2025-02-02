package sudokugo

type Sudoku struct {
	size  int
	Level Level
	board [9][9]int
}

// NewSudoku creates a new sudoku instance
func NewSudoku() *Sudoku {
	return &Sudoku{
		size: 9,
	}
}

// Get returns the Sudoku instance
func (s *Sudoku) Get() *Sudoku {
	return s
}
