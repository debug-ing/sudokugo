package sudokugo

import (
	"fmt"
	"math/rand"
)

// Sudoku represents the sudoku puzzle
func (s *Sudoku) InitBoard() *Sudoku {
	s.fillSudoku()
	return s
}

// Solve sudoku puzzle
func (s *Sudoku) Solve() *Sudoku {
	status := s.solveSudoku()
	if !status {
		panic("sudoku is not solvable")
	}
	return s
}

// SetBoard sets the board
func (s *Sudoku) SetBoard(board [9][9]int) *Sudoku {
	s.board = board
	return s
}

// GetBoard returns the board
func (s *Sudoku) GetBoard() [9][9]int {
	return s.board
}

// Print prints the board
func (s *Sudoku) Print() {
	for _, row := range s.board {
		for _, num := range row {
			if num == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", num)
			}
		}
		fmt.Println()
	}
}

// ssValid checks if the number is valid
func (s *Sudoku) isValid(board [9][9]int, row, col, num int) bool {
	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num ||
			board[boxRow+i/3][boxCol+i%3] == num {
			return false
		}
	}
	return true
}

// fillSudoku fills the sudoku puzzle
func (s *Sudoku) fillSudoku() bool {
	for row := 0; row < s.size; row++ {
		for col := 0; col < s.size; col++ {
			if s.board[row][col] == 0 {
				numbers := rand.Perm(9)
				for _, num := range numbers {
					if s.isValid(s.board, row, col, num+1) {
						s.board[row][col] = num + 1
						if s.fillSudoku() {
							return true
						}
						s.board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// solveSudoku solves the sudoku puzzle
func (s *Sudoku) solveSudoku() bool {
	for row := 0; row < s.size; row++ {
		for col := 0; col < s.size; col++ {
			if s.board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if s.isValid(s.board, row, col, num) {
						s.board[row][col] = num
						if s.solveSudoku() {
							return true
						}
						s.board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// GetStatusBoard returns the status of the board
func (s *Sudoku) GetStatusBoard() bool {
	for _, row := range s.board {
		for _, num := range row {
			if num == 0 {
				return false
			}
		}
	}
	return true
}
