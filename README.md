# Sudokugo
A simple and efficient Sudoku library for Go. This library provides the essential functionality to generate, solve, and validate Sudoku puzzles. It is designed for developers who want to integrate Sudoku features into their Go applications.

## Features
* Sudoku Generator: Generate random Sudoku puzzles.
* Sudoku Solver: Solve any valid Sudoku puzzle.
* Validation: Check if a given Sudoku puzzle is valid.
* Easy Integration: Simple API for seamless integration into Go applications.

## Installation
  ```sh
  go get github.com/debug-ing/sudokugo
  ```

## Usage

  ```go
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
  ```

## Contribute
We welcome contributions! If you have any ideas, improvements, or bug fixes, feel free to open an issue or a pull request.