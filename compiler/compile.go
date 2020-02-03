package compiler

import "strings"

// CompileString method
func CompileString(input string) ([]*Instruction, error) {
	reader := strings.NewReader(input)
	lexer, err := NewLexer(reader)

	if err != nil {
		return nil, err
	}

	parser, err := NewParser(lexer)

	if err != nil {
		return nil, err
	}

	node, err := parser.Parse()

	if err != nil {
		return nil, err
	}

	result, err := node.Traverse()

	if err != nil {
		return nil, err
	}

	return result.Load(), nil
}
