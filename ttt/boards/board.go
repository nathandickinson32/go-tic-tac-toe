package boards

import "fmt"

const (
	boardSize   = 3
	MinPosition = 1
	MaxPosition = 9
	PlayerX     = "X"
	PlayerO     = "O"
	EmptyCell   = ""
	TotalCells  = boardSize * boardSize
)

type Board [boardSize][boardSize]string

type GameStatus int

const (
	InProgress GameStatus = iota
	XWins
	OWins
	Draw
)

var winningLines = [8][3][2]int{
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

func NewBoard() Board {
	var board Board

	for row := range boardSize {
		for col := range boardSize {
			cell := boardSize*(row) + col + MinPosition // duplicated below
			board[row][col] = fmt.Sprintf("%d", cell)
		}
	}
	return board
}

func (board Board) IsPositionValid(position int) bool {
	if position < MinPosition || position > MaxPosition {
		return false
	}
	row, col := board.getCoordinates(position)
	return board[row][col] != PlayerX && board[row][col] != PlayerO // duplicated in fn below
}

func (board Board) AvailableMoves() []int {
	var moves []int
	for row := range boardSize {
		for col := range boardSize {
			if board[row][col] != PlayerX && board[row][col] != PlayerO {
				cell := boardSize*(row) + col + MinPosition // maybe needs its own fn
				moves = append(moves, cell)
			}
		}
	}
	return moves
}

func (board Board) getCoordinates(position int) (int, int) {
	adjustedPosition := position - MinPosition
	return adjustedPosition / boardSize, adjustedPosition % boardSize
}

func (board *Board) MakeMove(position int, player string) error {
	if position < MinPosition || position > MaxPosition {
		return fmt.Errorf("position must be between %d and %d", MinPosition, MaxPosition)
	}

	row, col := board.getCoordinates(position)
	if board[row][col] == PlayerX || board[row][col] == PlayerO {
		return fmt.Errorf("position already taken")
	}

	board[row][col] = player
	return nil
}

func tokenAtCell(board Board, cell [2]int) string {
	row := cell[0]
	col := cell[1]
	return board[row][col]
}

func (board Board) CheckWinner() string {
	for _, line := range winningLines {
		// firstPosition := board[line[0][0]][line[0][1]]
		firstPosition := tokenAtCell(board, line[0])
		secondPosition := board[line[1][0]][line[1][1]]
		thirdPosition := board[line[2][0]][line[2][1]]

		// What am I looking at?
		// This could be a function returning a boolean
		if firstPosition == secondPosition &&
			secondPosition == thirdPosition &&
			(firstPosition == PlayerX || firstPosition == PlayerO) {
			return firstPosition
		}
	}

	return EmptyCell
}

func (board Board) GetGameStatus() GameStatus {
	if winner := board.CheckWinner(); winner != EmptyCell {
		if winner == PlayerX {
			return XWins
		}
		return OWins
	}

	if len(board.AvailableMoves()) == 0 {
		return Draw
	}

	return InProgress
}
