package domain

import (
	"math/rand"
	"time"
)

// CellEmpty is the null cell value
const CellEmpty = "-"

// CellBomb represents the bomb value
const CellBomb = "X"

// Board is literraly the MinesWeeper board
type Board [][]string

// NewBoard is the Board constructor
func NewBoard(size uint, bombs uint) Board {
	board := newEmptyBoard(size)
	board.FillWithBombs(bombs)
	return board
}

// FillWithBombs is the method that fills the board with bombs
func (board Board) FillWithBombs(bombs uint) {
	rows := len(board)
	cols := len(board)
	positions := getRandomPositions(rows*cols, bombs)

	var row, col int

	for _, pos := range positions {
		row = pos / cols
		col = pos - row*rows

		board[row][col] = CellBomb
	}
}

func newEmptyBoard(size uint) Board {
	board := make([][]string, size)
	for row := range board {
		board[row] = make([]string, size)
	}
	for row := range board {
		for col := range board[0] {
			board[row][col] = CellEmpty
		}
	}
	return board
}

func getRandomPositions(size int, n uint) []int {
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(size)

	positions := []int{}

	for _, r := range p[:n] {
		positions = append(positions, r)
	}

	return positions
}
