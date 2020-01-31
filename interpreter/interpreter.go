package interpreter

import "strings"

// InterpretString interpretes string, roses are red, sky is blue
func InterpretString(input string) (interface{}, error) {
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

	return result, nil
}

// InterpretStringVM interpretes with VirtualMachine
func InterpretStringVM(input string) (interface{}, error) {
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

	bytecode := &Instruction{
		Instruction: START,
	}

	_, err = node.Compile(bytecode)

	if err != nil {
		return nil, err
	}

	vm := NewVM(bytecode)

	result, err := vm.Execute()

	if err != nil {
		return nil, err
	}

	return result, nil
}
