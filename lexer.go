package main

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	TokenLParen  TokenType = "LPAREN"
	TokenRParen  TokenType = "RPAREN"
	TokenNumber  TokenType = "NUMBER"
	TokenWord    TokenType = "WORD"
	TokenSymbol  TokenType = "SYMBOL"
	TokenKeyword TokenType = "KEYWORD"
)

type Token struct {
	Type  TokenType
	Value string
}

func getType(char rune) TokenType {
	switch {
	case char == '(':
		return TokenLParen
	case char == ')':
		return TokenRParen
	case isOperator(char):
		return TokenSymbol
	}
	return TokenNumber
}

func isOperator(token rune) bool {
	return strings.ContainsRune("+-*/", token)
}

func isKeyword(currentToken string) bool {
	return currentToken == "define"
}

func isWord(currentToken string) bool {
	return currentToken[0] == '"' && currentToken[len(currentToken)-1] == '"'
}

func createToken(currentToken string) Token {
	switch {
	case isKeyword(currentToken):
		return Token{Type: TokenKeyword, Value: currentToken}
	case isWord(currentToken):
		return Token{Type: TokenWord, Value: currentToken}
	default:
		return Token{Type: TokenSymbol, Value: currentToken}
	}
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	var currentToken string

	for _, char := range input {
		fixedRunes := char == '(' || char == ')' || isOperator(char) || unicode.IsDigit(char)
		switch {
		case fixedRunes:
			if currentToken != "" {
				tokens = append(tokens, createToken(currentToken))
				currentToken = ""
			}
			tokens = append(tokens, Token{Type: getType(char), Value: string(char)})
		case unicode.IsSpace(char):
			if currentToken != "" {
				tokens = append(tokens, createToken(currentToken))
				currentToken = ""
			}
		default:
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, createToken(currentToken))
	}

	return tokens, nil
}
