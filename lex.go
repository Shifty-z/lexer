package lexer

import (
	"fmt"
	"unicode"
)

type TokenType int32

type Token struct {
	Field rune
	Token TokenType
}

func (tok Token) String() string {
	return fmt.Sprintf("Token: %c Name: %s\n", tok.Field, tokenTypeToName[int32(tok.Token)])
}

const (
	Unknown = iota
	Whitespace
	Number
	ExclamationPoint
	Ampersand
	UnicodeLetter
	Comma
	QuestionMark
	Period
	Colon
	SemiColon
	GreaterThan
	LessThan
	LeftParenthesis
	RightParenthesis
	LeftCurlyBrace
	RightCurlyBrace
	SubtractionSign
	EqualitySign
	Percent
	Newline
	ForwardSlash
	BackwardSlash
	Quote
	Apostrophe
	LeftBracket
	RightBracket
	UtcDateZ
	UtcDateT
)

var tokenTypeToName = map[int32]string{
	Unknown:          "Unknown",
	Whitespace:       "Whitespace",
	Number:           "Number",
	ExclamationPoint: "Exclamation Point",
	Ampersand:        "Ampersand",
	UnicodeLetter:    "Letter",
	Comma:            "Comma",
	QuestionMark:     "Question Mark",
	Period:           "Period",
	Colon:            "Colon",
	SemiColon:        "Semi-Colon",
	GreaterThan:      "Greater Than",
	LessThan:         "Less Than",
	LeftParenthesis:  "Left Parenthesis",
	RightParenthesis: "Right Parenthesis",
	LeftCurlyBrace:   "Left Curly Brace",
	RightCurlyBrace:  "Right Curly Brace",
	SubtractionSign:  "Subtraction Sign",
	EqualitySign:     "Equality Symbol",
	Percent:          "Percent Symbol",
	Newline:          "Newline Character",
	ForwardSlash:     "/",
	BackwardSlash:    "\\",
	Quote:            "Quote",
	Apostrophe:       "Apostrophe",
	LeftBracket:      "Left Bracket",
	RightBracket:     "Right Bracket",
	UtcDateZ:         "UTC Date Z",
	UtcDateT:         "UTC Date T",
}

// Lex - Given an input string, lexically analyzes each character turning it into a recognized lexical token.
func Lex(input string) []Token {
	tokens := make([]Token, 0)

	for _, ch := range input {
		switch {
		case ch == ' ':
			tokens = append(tokens, Token{ch, Whitespace})
		case ch >= '0' && ch <= '9':
			tokens = append(tokens, Token{ch, Number})
		case ch == '!':
			tokens = append(tokens, Token{ch, ExclamationPoint})
		case ch == '&':
			tokens = append(tokens, Token{ch, Ampersand})
		case ch == ',':
			tokens = append(tokens, Token{ch, Comma})
		case ch == '?':
			tokens = append(tokens, Token{ch, QuestionMark})
		case ch == '.':
			tokens = append(tokens, Token{ch, Period})
		case ch == ':':
			tokens = append(tokens, Token{ch, Colon})
		case ch == ';':
			tokens = append(tokens, Token{ch, SemiColon})
		case ch == '>':
			tokens = append(tokens, Token{ch, GreaterThan})
		case ch == '<':
			tokens = append(tokens, Token{ch, LessThan})
		case ch == '(':
			tokens = append(tokens, Token{ch, LeftParenthesis})
		case ch == ')':
			tokens = append(tokens, Token{ch, RightParenthesis})
		case ch == '%':
			tokens = append(tokens, Token{ch, Percent})
		case ch == '\n':
			tokens = append(tokens, Token{ch, Newline})
		case ch == '=':
			tokens = append(tokens, Token{ch, EqualitySign})
		case ch == '-':
			tokens = append(tokens, Token{ch, SubtractionSign})
		case ch == '{':
			tokens = append(tokens, Token{ch, LeftCurlyBrace})
		case ch == '}':
			tokens = append(tokens, Token{ch, RightCurlyBrace})
		case ch == '/':
			tokens = append(tokens, Token{ch, ForwardSlash})
		case ch == '\\':
			tokens = append(tokens, Token{ch, BackwardSlash})
		case ch == '"':
			tokens = append(tokens, Token{ch, Quote})
		case ch == '\'':
			tokens = append(tokens, Token{ch, Apostrophe})
		case ch == '[':
			tokens = append(tokens, Token{ch, LeftBracket})
		case ch == ']':
			tokens = append(tokens, Token{ch, RightBracket})
		case ch == 'Z':
			tokens = append(tokens, Token{ch, UtcDateZ})
		case ch == 'T':
			tokens = append(tokens, Token{ch, UtcDateT})
		case unicode.IsLetter(ch):
			tokens = append(tokens, Token{ch, UnicodeLetter})
		default:
			tokens = append(tokens, Token{ch, Unknown})
		}
	}

	return tokens
}
