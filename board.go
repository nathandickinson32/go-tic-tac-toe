package main

import "fmt"

const boardSize = 3

type Board [boardSize][boardSize]string

func NewBoard() Board {
	var board Board
	cell := 1

	for row := range boardSize {
		for col := range boardSize {
			board[row][col] = fmt.Sprintf("%d", cell)
			cell++
		}
	}
	return board
}

func (board Board) AvailableMoves() []int {
	var moves []int
	cell := 1
	for row := range boardSize {
		for col := range boardSize {
			if board[row][col] != "X" && board[row][col] != "O" {
				moves = append(moves, cell)
			}
			cell++
		}
	}
	return moves
}

func (board Board) getCoordinates(position int) (int, int) {
	position--
	return position / boardSize, position % boardSize
}

func (board *Board) MakeMove(position int, player string) error {
	if position < 1 || position > 9 {
		return fmt.Errorf("position must be between 1 and 9")
	}

	row, col := board.getCoordinates(position)
	if board[row][col] == "X" || board[row][col] == "O" {
		return fmt.Errorf("position already taken")
	}

	board[row][col] = player
	return nil
}
