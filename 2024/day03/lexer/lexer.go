package lexer

import (
	"strings"
	"unicode"
)

type Predicate[T any] func(value T) bool

type TokenType string

type Token struct {
	Kind TokenType
	Raw  string
}

const (
	LParen    TokenType = "("
	RParen    TokenType = ")"
	Comma     TokenType = ","
	Int       TokenType = "INTEGER"
	EOF       TokenType = "EOF"
	Illegal   TokenType = "ILLEGAL"
	IdentMult TokenType = "mul"
	IdentDo   TokenType = "do"
	IdentDont TokenType = "don't"
)

func createToken(kind TokenType, raw string) Token {
	return Token{
		Kind: kind,
		Raw:  raw,
	}
}

func parseIdent(text string) (Token, bool) {
	if strings.HasSuffix(text, string(IdentMult)) {
		return createToken(IdentMult, text), true
	}
	if strings.HasSuffix(text, string(IdentDo)) {
		return createToken(IdentDo, text), true
	}
	if strings.HasSuffix(text, string(IdentDont)) {
		return createToken(IdentDont, text), true
	}
	return Token{}, false
}

func NewLexer(input string) Lexer {
	lex := Lexer{
		reader: strings.NewReader(input),
	}
	return lex
}

type Lexer struct {
	reader *strings.Reader
	lastCh rune
	ch     rune
}

func (l *Lexer) unreadChar() {
	l.reader.UnreadRune()
	l.ch = l.lastCh
}

func (l *Lexer) readChar() rune {
	l.lastCh = l.ch

	// There are no more runes to read
	if l.reader.Len() == 0 {
		l.ch = '\x00'
		return l.ch
	}

	ch, _, err := l.reader.ReadRune()
	if err != nil {
		panic(err)
	} else {
		l.ch = ch
	}

	return l.ch
}

func (l *Lexer) readCharWhile(predicate Predicate[rune]) string {
	chars := []rune{l.ch}

	for {
		if predicate(l.ch) {
			chars = append(chars, l.readChar())
		} else {
			break
		}
	}

	// TODO: This can be improved to avoid the unread as we incorrectly
	// push the last rune into the chars slice, which must be removed
	l.unreadChar()

	if len(chars) == 1 {
		return string(chars)
	} else {
		return string(chars[0 : len(chars)-1])
	}
}

func (l *Lexer) readIdent() string {
	return l.readCharWhile(func(c rune) bool {
		return c == '\'' || unicode.IsLetter(c)
	})
}

func (l *Lexer) readInt() string {
	return l.readCharWhile(unicode.IsDigit)
}

func (l *Lexer) Tokens() []Token {
	tokens := make([]Token, 0)

	for {
		l.readChar()

		if l.ch == '(' {
			tokens = append(tokens, createToken(LParen, string(l.ch)))
		} else if l.ch == ')' {
			tokens = append(tokens, createToken(RParen, string(l.ch)))
		} else if l.ch == ',' {
			tokens = append(tokens, createToken(Comma, string(l.ch)))
		} else if l.ch == '\x00' {
			break
		} else if unicode.IsLetter(l.ch) {
			text := l.readIdent()
			ident, ok := parseIdent(text)
			if ok {
				tokens = append(tokens, ident)
			} else {
				tokens = append(tokens, createToken(Illegal, string(l.ch)))
			}
		} else if unicode.IsDigit(l.ch) {
			number := l.readInt()
			tokens = append(tokens, createToken(Int, number))
		} else {
			tokens = append(tokens, createToken(Illegal, string(l.ch)))
		}
	}

	return tokens
}
