package main

import (
	"fmt"

	"github.com/debug-ing/sudokugo"
)

func main() {
	data := sudokugo.NewSudoku().SetLevel(sudokugo.Extreme).InitBoard().RemoveNumbers().GetBoard()
	fmt.Print("Sudoku Board\n", data)
	result := sudokugo.NewSudoku().SetBoard(data).Solve().GetStatusBoard()
	fmt.Println(result)
}
