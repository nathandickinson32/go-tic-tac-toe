package main

import "io"

type GameMode int

const (
	HumanVsHuman GameMode = iota
	HumanVsAI
	AIVsAI
)

func ParseMode(input string, output *ConsoleOutput) GameMode {
	switch input {
	case "1":
		output.ShowModeStart("Human vs Human")
		return HumanVsHuman
	case "2":
		output.ShowModeStart("Human vs AI (you are X, AI is O)")
		return HumanVsAI
	case "3":
		output.ShowModeStart("AI vs AI")
		return AIVsAI
	default:
		output.ShowModeStart("Invalid selection, defaulting to Human vs Human")
		return HumanVsHuman
	}
}

func StartGameWithMode(reader io.Reader, writer io.Writer, mode GameMode) {
	rules := NewGameRules()
	output := NewConsoleOutput(writer)

	var playerXInput, playerOInput InputReader

	switch mode {
	case HumanVsHuman:
		sharedInput := NewConsoleInput(reader, output)
		playerXInput = sharedInput
		playerOInput = sharedInput

	case HumanVsAI:
		playerXInput = NewConsoleInput(reader, output)
		playerOInput = NewAIPlayer("O", "X", rules, output)

	case AIVsAI:
		playerXInput = NewAIPlayer("X", "O", rules, output)
		playerOInput = NewAIPlayer("O", "X", rules, output)
	}

	game := NewGame(rules, playerXInput, playerOInput, output)
	game.Start()
}
