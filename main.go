package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	output := NewConsoleOutput(os.Stdout)

	output.ShowModeSelection()

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	mode := ParseMode(input, output)

	output.ShowNewline()
	StartGameWithMode(os.Stdin, os.Stdout, mode)
}
