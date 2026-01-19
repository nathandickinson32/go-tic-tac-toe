package main

import (
	"bufio"
	"os"
)

type Game struct {
	board         Board
	rules         *GameRules
	playerXInput  InputReader
	playerOInput  InputReader
	output        *ConsoleOutput
	currentPlayer string
}

func NewGame(
	rules *GameRules,
	playerXInput InputReader,
	playerOInput InputReader,
	output *ConsoleOutput,
	firstPlayer string) *Game {
	return &Game{
		board:         NewBoard(),
		rules:         rules,
		playerXInput:  playerXInput,
		playerOInput:  playerOInput,
		output:        output,
		currentPlayer: firstPlayer,
	}
}

func (game *Game) getCurrentPlayerInput() InputReader {
	if game.currentPlayer == "X" {
		return game.playerXInput
	}
	return game.playerOInput
}

func (game *Game) switchPlayer() {
	if game.currentPlayer == "X" {
		game.currentPlayer = "O"
	} else {
		game.currentPlayer = "X"
	}
}

func (game *Game) displayEndResult(status GameStatus) {
	switch status {
	case XWins:
		game.output.ShowWinner("X")
	case OWins:
		game.output.ShowWinner("O")
	case Draw:
		game.output.ShowDraw()
	}
}

func (game *Game) playTurns() {
	for {
		game.output.ShowPlayerTurn(game.currentPlayer)

		position, err := game.getCurrentPlayerInput().ReadMove(game.board)
		if err != nil {
			break
		}

		if err := game.board.MakeMove(position, game.currentPlayer); err != nil {
			break
		}

		game.output.ShowBoard(game.board)

		if status := game.rules.GetGameStatus(game.board); status != InProgress {
			game.displayEndResult(status)
			break
		}

		game.switchPlayer()
	}
}

func (game *Game) PlayGame() {
	game.output.ShowWelcome()
	game.output.ShowBoard(game.board)
	game.playTurns()
}

func buildGame(reader *bufio.Reader, input *ConsoleInput, output *ConsoleOutput) *Game {
	output.ShowPlayerTypeSelection("X")
	playerXType, _ := input.ReadPlayerType()

	output.ShowPlayerTypeSelection("O")
	playerOType, _ := input.ReadPlayerType()

	output.ShowFirstPlayerSelection()
	firstPlayer, _ := input.ReadFirstPlayer()

	output.ShowNewline()

	rules := NewGameRules()
	playerXInput := createPlayerInput(playerXType, "X", "O", reader, output, rules)
	playerOInput := createPlayerInput(playerOType, "O", "X", reader, output, rules)

	return NewGame(rules, playerXInput, playerOInput, output, firstPlayer)
}

func startGame(output *ConsoleOutput) {
	reader := bufio.NewReader(os.Stdin)
	input := NewConsoleInput(reader, output)

	for {
		output.ShowNewline()

		game := buildGame(reader, input, output)
		game.PlayGame()

		output.ShowPlayAgainPrompt()
		playAgain, err := input.ReadPlayAgain()

		if err != nil || !playAgain {
			output.ShowGoodbye()
			break
		}
	}
}
