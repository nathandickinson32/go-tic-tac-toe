package game

import (
	"bufio"
	"io"
	"os"
	"ttt/boards"
	tttio "ttt/io"
	"ttt/players"
)

type Game struct {
	board         boards.Board
	playerX       players.Player
	playerO       players.Player
	output        io.Writer
	currentPlayer string
}

func NewGame(
	playerX players.Player,
	playerO players.Player,
	output io.Writer) *Game {
	return &Game{
		board:         boards.NewBoard(),
		playerX:       playerX,
		playerO:       playerO,
		output:        output,
		currentPlayer: boards.PlayerX,
	}
}

func (game *Game) getCurrentPlayer() players.Player {
	if game.currentPlayer == boards.PlayerX {
		return game.playerX
	}
	return game.playerO
}

func (game *Game) switchPlayer() {
	if game.currentPlayer == boards.PlayerX {
		game.currentPlayer = boards.PlayerO
	} else {
		game.currentPlayer = boards.PlayerX
	}
}

func (game *Game) displayEndResult(status boards.GameStatus) {
	switch status {
	case boards.XWins:
		tttio.ShowWinner(game.output, boards.PlayerX)
	case boards.OWins:
		tttio.ShowWinner(game.output, boards.PlayerO)
	case boards.Draw:
		tttio.ShowDraw(game.output)
	}
}

func (game *Game) playTurns() {
	for {
		tttio.ShowPlayerTurn(game.output, game.currentPlayer)

		position, err := game.getCurrentPlayer().ReadMove(game.board)
		if err != nil {
			break
		}

		if err := game.board.MakeMove(position, game.currentPlayer); err != nil {
			break
		}

		tttio.ShowBoard(game.output, game.board)

		if status := game.board.GetGameStatus(); status != boards.InProgress {
			game.displayEndResult(status)
			break
		}

		game.switchPlayer()
	}
}

func (game *Game) PlayGame() {
	tttio.ShowWelcome(game.output)
	tttio.ShowBoard(game.output, game.board)
	game.playTurns()
}

func BuildGame(reader *bufio.Reader, output io.Writer) *Game {
	tttio.ShowPlayerTypeSelection(output, boards.PlayerX)
	playerXType, _ := tttio.ReadPlayerType(reader, output)

	tttio.ShowPlayerTypeSelection(output, boards.PlayerO)
	playerOType, _ := tttio.ReadPlayerType(reader, output)

	tttio.ShowNewline(output)

	playerX := players.CreatePlayer(playerXType, boards.PlayerX, boards.PlayerO, reader, output)
	playerO := players.CreatePlayer(playerOType, boards.PlayerO, boards.PlayerX, reader, output)

	return NewGame(playerX, playerO, output)
}

func StartGame() {
	reader := bufio.NewReader(os.Stdin)
	output := os.Stdout

	for {
		tttio.ShowNewline(output)

		game := BuildGame(reader, output)
		game.PlayGame()

		tttio.ShowPlayAgainPrompt(output)
		playAgain, err := tttio.ReadPlayAgain(reader, output)

		if err != nil || !playAgain {
			tttio.ShowGoodbye(output)
			break
		}
	}
}
