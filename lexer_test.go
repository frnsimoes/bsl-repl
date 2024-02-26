package main

import (
	"reflect"
	"testing"
)

func TestTokenizeSymbolOperator(t *testing.T) {
	input := "(+ 1 2)"
	expected := []Token{
		{Type: TokenLParen, Value: "("},
		{Type: TokenSymbol, Value: "+"},
		{Type: TokenNumber, Value: "1"},
		{Type: TokenNumber, Value: "2"},
		{Type: TokenRParen, Value: ")"},
	}

	tokens, err := Tokenize(input)
	if err != nil {
		t.Fatalf("Tokenize returned error: %v", err)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokenize = %v, want %v", tokens, expected)
	}

}

func TestTokenizeDefine(t *testing.T) {
	input := "(define name \"John\")"
	expected := []Token{
		{Type: TokenLParen, Value: "("},
		{Type: TokenKeyword, Value: "define"},
		{Type: TokenSymbol, Value: "name"},
		{Type: TokenWord, Value: "\"John\""},
		{Type: TokenRParen, Value: ")"},
	}

	tokens, err := Tokenize(input)
	if err != nil {
		t.Fatalf("Tokenize returned error: %v", err)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokenize = %v, want %v", tokens, expected)
	}
}
