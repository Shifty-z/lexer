package lex

import (
	"bufio"
	"fmt"
	"os"
)

// GetTokens - Given a file, lexically analyzes each `character` in the input and converts it to a single LexToken type.
func GetTokens(file *os.File) [][]Token {
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Printf("Scanner error: %q\n", scanner.Err())
		panic("Scanner encountered an error!")
	}

	lexedLines := make([][]Token, 0)
	for _, line := range lines {
		lexedLines = append(lexedLines, Lex(line))
	}

	return lexedLines
}
