package main

type GameStatus int

const (
	InProgress GameStatus = iota
	XWins
	OWins
	Draw
)

type GameRules struct{}

func NewGameRules() *GameRules {
	return &GameRules{}
}

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

func (rules *GameRules) CheckWinner(board Board) string {
	for _, line := range winningLines {
		a := board[line[0][0]][line[0][1]]
		b := board[line[1][0]][line[1][1]]
		c := board[line[2][0]][line[2][1]]

		if a == b && b == c && (a == "X" || a == "O") {
			return a
		}
	}

	return ""
}

func (rules *GameRules) GetGameStatus(board Board) GameStatus {
	if winner := rules.CheckWinner(board); winner != "" {
		if winner == "X" {
			return XWins
		}
		return OWins
	}

	if len(board.AvailableMoves()) == 0 {
		return Draw
	}

	return InProgress
}
