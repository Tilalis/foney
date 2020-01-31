package interpreter

import (
	"fmt"
	"reflect"
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
	types  map[string]TypeInfo
}

var symbolTable SymbolTable = SymbolTable{
	values: make(map[string]interface{}),
	types:  make(map[string]TypeInfo),
}

// GetSymbolTable returns SymbolTable
// It returns same instance every time, if one is not costructed, then it custructs it.
// Well, maybe here comments are useful
func GetSymbolTable() SymbolTable {
	return symbolTable
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

// GetType gets type
func (s *SymbolTable) GetType(name string) (TypeInfo, error) {
	value, ok := s.types[name]

	if !ok {
		return 0, fmt.Errorf("%w '%s'", ErrNoSuchSymbol, name)
	}

	return value, nil
}

// Set Sets Symbol
func (s *SymbolTable) Set(name string, value interface{}) {
	s.values[name] = value
	var typeOf = reflect.Indirect(reflect.ValueOf(value)).Type()
	var typeName = typeOf.Name()

	switch typeName {
	case "float64":
		s.types[name] = NUMBERTYPE
	case "Money":
		s.types[name] = MONEYTYPE
	default:
		s.types[name] = DYNAMIC
	}
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

// Compile compiles
func (a *Assign) Compile(b *Instruction) (*Instruction, error) {
	b, err := a.Value.Compile(b)

	if err != nil {
		return nil, err
	}

	name := a.Symbol.Token.Value.(string)

	return b.Append(NewInstruction(SET, name)), nil

	// return nil, ErrUnsoppertedOperatrion
}
