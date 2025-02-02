package sudokugo

type Sudoku struct {
	size  int
	Level Level
	board [9][9]int
}

func NewSudoku() *Sudoku {
	return &Sudoku{
		size: 9,
	}
}

func (s *Sudoku) Get() *Sudoku {
	return s
}
