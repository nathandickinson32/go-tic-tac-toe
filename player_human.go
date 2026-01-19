package main

import (
	"errors"
	"strconv"
	"strings"
)

func (ci *ConsoleInput) isPositionAvailable(board Board, position int) bool {
	boardCopy := board
	return boardCopy.MakeMove(position, "X") == nil
}

var (
	ErrEmptyInput    = errors.New("Input cannot be empty")
	ErrNotNumber     = errors.New("Input must be a number")
	ErrOutOfRange    = errors.New("Position must be between 1 and 9")
	ErrInvalidChoice = errors.New("Invalid choice")
)

func (ci *ConsoleInput) parseInput(input string) (int, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, ErrEmptyInput
	}

	position, err := strconv.Atoi(input)
	if err != nil {
		return 0, ErrNotNumber
	}

	if position < 1 || position > 9 {
		return 0, ErrOutOfRange
	}

	return position, nil
}

func (input *ConsoleInput) getValidPosition() (int, error) {
	for {
		input.output.ShowPrompt()
		line, err := input.reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		position, err := input.parseInput(line)
		if err != nil {
			input.output.ShowInvalidInput(err)
			continue
		}

		return position, nil
	}
}

func (ci *ConsoleInput) ReadMove(board Board) (int, error) {
	for {
		position, err := ci.getValidPosition()
		if err != nil {
			return 0, err
		}

		if ci.isPositionAvailable(board, position) {
			return position, nil
		}

		ci.output.ShowPositionTaken()
	}
}
