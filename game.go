package main

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
) *Game {
	return &Game{
		board:         NewBoard(),
		rules:         rules,
		playerXInput:  playerXInput,
		playerOInput:  playerOInput,
		output:        output,
		currentPlayer: "X",
	}
}

func (game *Game) Start() {
	game.output.ShowWelcome()
	game.output.ShowBoard(game.board)
	game.playTurns()
}

func (game *Game) playTurns() {
	for {
		game.output.ShowPlayerTurn(game.currentPlayer)

		input := game.getInputForCurrentPlayer()
		position, err := input.ReadMove(game.board)
		if err != nil {
			return
		}

		if err := game.board.MakeMove(position, game.currentPlayer); err != nil {
			return
		}
		game.output.ShowBoard(game.board)

		if game.checkEndCondition() {
			break
		}

		game.switchPlayer()
	}
}

func (game *Game) getInputForCurrentPlayer() InputReader {
	if game.currentPlayer == "X" {
		return game.playerXInput
	}
	return game.playerOInput
}

func (game *Game) checkEndCondition() bool {
	status := game.rules.GetGameStatus(game.board)
	if status == InProgress {
		return false
	}

	game.displayEndResult(status)
	return true
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
