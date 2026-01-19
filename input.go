package main

import (
	"bufio"
	"errors"
	"strings"
)

type ConsoleInput struct {
	reader *bufio.Reader
	output *ConsoleOutput
}

func NewConsoleInput(reader *bufio.Reader, output *ConsoleOutput) *ConsoleInput {
	return &ConsoleInput{
		reader: reader,
		output: output,
	}
}

func (ci *ConsoleInput) promptAndReadLine() (string, error) {
	return ci.reader.ReadString('\n')
}

func (ci *ConsoleInput) parsePlayerType(input string) (PlayerType, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return Human, errors.New("Input cannot be empty")
	}

	switch input {
	case "1":
		return Human, nil
	case "2":
		return AI, nil
	default:
		return Human, errors.New("Invalid choice. Enter 1 for Human or 2 for AI")
	}
}

func (ci *ConsoleInput) ReadPlayerType() (PlayerType, error) {
	for {
		line, err := ci.promptAndReadLine()
		if err != nil {
			return Human, err
		}

		playerType, err := ci.parsePlayerType(line)
		if err != nil {
			ci.output.ShowInvalidInput(err)
			continue
		}

		return playerType, nil
	}
}

func (ci *ConsoleInput) parseFirstPlayer(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "X", errors.New("Input cannot be empty")
	}

	switch input {
	case "1":
		return "X", nil
	case "2":
		return "O", nil
	default:
		return "X", errors.New("Invalid choice. Enter 1 for Player X or 2 for Player O")
	}
}

func (ci *ConsoleInput) ReadFirstPlayer() (string, error) {
	for {
		line, err := ci.promptAndReadLine()
		if err != nil {
			return "X", err
		}

		firstPlayer, err := ci.parseFirstPlayer(line)
		if err != nil {
			ci.output.ShowInvalidInput(err)
			continue
		}

		return firstPlayer, nil
	}
}

func (ci *ConsoleInput) parsePlayAgain(input string) (bool, error) {
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "" {
		return false, errors.New("Input cannot be empty")
	}

	switch input {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, errors.New("Invalid input. Enter 'y' for yes or 'n' for no")
	}
}

func (ci *ConsoleInput) ReadPlayAgain() (bool, error) {
	for {
		line, err := ci.promptAndReadLine()
		if err != nil {
			return false, err
		}

		playAgain, err := ci.parsePlayAgain(line)
		if err != nil {
			ci.output.ShowInvalidInput(err)
			continue
		}

		return playAgain, nil
	}
}
