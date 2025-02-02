package sudokugo

import (
	"math/rand"
)

func (s *Sudoku) RemoveNumbers() *Sudoku {
	s.board = s.removeNumbers(s.board)
	return s
}

func (s *Sudoku) removeNumbers(board [9][9]int) [9][9]int {
	levels := s.GetAllLevels()
	count := levels[string(s.Level)]
	for i := 0; i < count; i++ {
		row, col := rand.Intn(9), rand.Intn(9)
		for board[row][col] == 0 {
			row, col = rand.Intn(9), rand.Intn(9)
		}
		board[row][col] = 0
	}
	return board
}
