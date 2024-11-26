package lexer

import (
	"bufio"
	"fmt"
	"github.com/Shifty-z/lexer/lex"
	"os"
)

// GetTokens - Given a file, lexically analyzes each `character` in the input and converts it to a single LexToken type.
func GetTokens(file *os.File) [][]lex.Token {
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Printf("Scanner error: %q\n", scanner.Err())
		panic("Scanner encountered an error!")
	}

	lexedLines := make([][]lex.Token, 0)
	for _, line := range lines {
		lexedLines = append(lexedLines, lex.Lex(line))
	}

	return lexedLines
}
