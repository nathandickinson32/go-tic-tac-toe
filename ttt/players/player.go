package players

import (
	"bufio"
	"io"
	"ttt/boards"
	tttio "ttt/io"
)

type Player interface {
	ReadMove(board boards.Board) (int, error)
}

func CreatePlayer(
	playerType tttio.PlayerType,
	symbol string,
	opponentSymbol string,
	reader *bufio.Reader,
	output io.Writer,
) Player {
	if playerType == tttio.AI {
		return NewAIPlayer(symbol, opponentSymbol)
	}
	return NewHumanPlayer(reader, output)
}
