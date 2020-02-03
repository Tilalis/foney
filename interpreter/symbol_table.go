package interpreter

import (
	"fmt"
)

// SymbolTable is just SymbolTable, like it's name says and that's all.
// Yeah, really. That's simple. Just what it's name says.
// I'm not kidding, that's true! This is a great example of self-descriptive name.
// Self-descriptive name is that name which describes itself, like the term "self-descriptive".
//
// Well, maybe one can argue that `SymbolTable` is not good enough to describe the purpose of this type inside of interpreter package...
// But what fucking else could this type be?
type SymbolTable struct {
	values map[string]interface{}
}

var globalSymbolTable *SymbolTable = NewSymbolTable()

// NewSymbolTable returns symbol table
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		values: make(map[string]interface{}),
	}
}

// GetGlobalSymbolTable returns SymbolTable
// It returns same instance every time, if one is not costructed, then it custructs it.
// Well, maybe here comments are useful
func GetGlobalSymbolTable() *SymbolTable {
	return globalSymbolTable
}

// Get returns Symbol Value from table
// Returns NoSuchSymbol error is no name is present in SymbolTable
func (s *SymbolTable) Get(name string) (interface{}, error) {
	value, ok := s.values[name]

	if !ok {
		return nil, fmt.Errorf("%w '%s'", ErrNoSuchSymbol, name)
	}

	return value, nil
}

// Set Sets Symbol
func (s *SymbolTable) Set(name string, value interface{}) {
	s.values[name] = value
}
