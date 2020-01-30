package interpreter

// InterpretString interpretes string, roses are red, sky is blue
func InterpretString(input string) (interface{}, error) {
	lexer, err := NewLexer(input)

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

	return result, nil
}
