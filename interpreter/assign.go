package interpreter

import "fmt"

// SymbolTable is just SymbolTable, like it's name says and that's all.
// Yeah, really. That's simple. Just what it's name says.
// I'm not kidding, that's true! This is a great example of self-descriptive name.
// Self-descriptive name is that name which describes itself, like the term "self-descriptive".
//
// Well, maybe one can argue that `SymbolTable` is not good enough to describe the purpose of this type inside of interpreter package...
// But what fucking else could this type be?
type SymbolTable map[string]interface{}

var symbolTable SymbolTable = make(map[string]interface{})

// GetSymbolTable returns SymbolTable
// It returns same instance every time, if one is not costructed, then it custructs it.
// Well, maybe here comments are useful
func GetSymbolTable() SymbolTable {
	return symbolTable
}

// Get returns Symbol Value from table
// Returns NoSuchSymbol error is no name is present in SymbolTable
func (s *SymbolTable) Get(name string) (interface{}, error) {
	value, ok := symbolTable[name]

	if !ok {
		return nil, fmt.Errorf("%w '%s'", ErrNoSuchSymbol, name)
	}

	return value, nil
}

// Set Sets Symbol
func (s *SymbolTable) Set(name string, value interface{}) {
	symbolTable[name] = value
}

// Assign represents assign operation
type Assign struct {
	Symbol *Symbol
	Value  AST
}

// Traverse traverse assign
func (a *Assign) Traverse() (interface{}, error) {
	name := a.Symbol.Token.Value.(string)
	value, err := a.Value.Traverse()

	if err != nil {
		return nil, err
	}

	symbolTable.Set(name, value)

	return value, nil
}
