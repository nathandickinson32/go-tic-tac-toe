package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func parseInput(input string) (int, error) {
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

func getUserMove(board *Board, reader io.Reader, writer io.Writer) int {
	bufReader := bufio.NewReader(reader)
	var bufWriter *bufio.Writer
	if writer != nil {
		bufWriter = bufio.NewWriter(writer)
		defer bufWriter.Flush()
	}
	return getValidUserMove(board, bufReader, bufWriter)
}

func getValidUserMove(board *Board, bufReader *bufio.Reader, bufWriter *bufio.Writer) int {
	for {
		if bufWriter != nil {
			displayPrompt(bufWriter)
		}

		line, err := bufReader.ReadString('\n')
		if err != nil {
			return 0
		}

		position, parseErr := parseInput(line)

		if parseErr != nil {
			if bufWriter != nil {
				displayInvalidInput(parseErr, bufWriter)
			}
			continue
		}

		row := (position - 1) / 3
		col := (position - 1) % 3

		if board[row][col] == "X" || board[row][col] == "O" {
			if bufWriter != nil {
				displayPositionTaken(bufWriter)
			}
			continue
		}

		return position
	}
}
