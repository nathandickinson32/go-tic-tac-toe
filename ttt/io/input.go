package io

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

const (
	HumanChoice = "1"
	AIChoice    = "2"
	YesShort    = "y"
	YesLong     = "yes"
	NoShort     = "n"
	NoLong      = "no"
	EmptyInput  = ""
)

type PlayerType int

const (
	Human PlayerType = iota
	AI
)

func ReadPlayerType(reader *bufio.Reader, output io.Writer) (PlayerType, error) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return Human, err
		}

		playerType, err := parsePlayerType(line)
		if err != nil {
			ShowInvalidInput(output, err)
			continue
		}

		return playerType, nil
	}
}

func parsePlayerType(input string) (PlayerType, error) {
	input = strings.TrimSpace(input)

	if input == EmptyInput {
		return Human, errors.New("Input cannot be empty")
	}

	switch input {
	case HumanChoice:
		return Human, nil
	case AIChoice:
		return AI, nil
	default:
		return Human, errors.New("Invalid choice. Enter 1 for Human or 2 for AI")
	}
}

func ReadPlayAgain(reader *bufio.Reader, output io.Writer) (bool, error) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		playAgain, err := parsePlayAgain(line)
		if err != nil {
			ShowInvalidInput(output, err)
			continue
		}

		return playAgain, nil
	}
}

func parsePlayAgain(input string) (bool, error) {
	input = strings.TrimSpace(strings.ToLower(input))

	if input == EmptyInput {
		return false, errors.New("Input cannot be empty")
	}

	switch input {
	case YesShort, YesLong:
		return true, nil
	case NoShort, NoLong:
		return false, nil
	default:
		return false, errors.New("Invalid input. Enter 'y' for yes or 'n' for no")
	}
}
