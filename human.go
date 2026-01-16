package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type ConsoleInput struct {
	reader *bufio.Reader
	output *ConsoleOutput
}

func NewConsoleInput(reader io.Reader, output *ConsoleOutput) *ConsoleInput {
	return &ConsoleInput{
		reader: bufio.NewReader(reader),
		output: output,
	}
}

func (ci *ConsoleInput) ReadMove(board Board) (int, error) {
	for {
		line, err := ci.promptAndReadLine()
		if err != nil {
			return 0, err
		}

		position, err := ci.parseInput(line)
		if err != nil {
			ci.output.ShowInvalidInput(err)
			continue
		}

		boardCopy := board
		if err := boardCopy.MakeMove(position, "X"); err != nil {
			ci.output.ShowPositionTaken()
			continue
		}

		return position, nil
	}
}

func (ci *ConsoleInput) promptAndReadLine() (string, error) {
	ci.output.ShowPrompt()
	return ci.reader.ReadString('\n')
}

func (ci *ConsoleInput) parseInput(input string) (int, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, errors.New("Input cannot be empty")
	}

	position, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("Input must be a number")
	}

	if position < 1 || position > 9 {
		return 0, errors.New("Position must be between 1 and 9")
	}

	return position, nil
}
