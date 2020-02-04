package vm

import (
	"fmt"
	"reflect"

	"github.com/Tilalis/foney/compiler"
)

// SymbolTable is just SymbolTable, like it's name says and that's all.
type SymbolTable struct {
	values map[string]interface{}
	types  map[string]compiler.TypeInfo
}

// NewSymbolTable returns symbol table
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		values: make(map[string]interface{}),
		types:  make(map[string]compiler.TypeInfo),
	}
}

// Get returns Symbol Value from table
// Returns NoSuchSymbol error is no name is present in SymbolTable
func (s *SymbolTable) Get(name string) (interface{}, error) {
	value, ok := s.values[name]

	if !ok {
		return nil, fmt.Errorf("%w '%s'", compiler.ErrNoSuchSymbol, name)
	}

	return value, nil
}

// GetType gets type
func (s *SymbolTable) GetType(name string) (compiler.TypeInfo, error) {
	value, ok := s.types[name]

	if !ok {
		return 0, fmt.Errorf("%w '%s'", compiler.ErrNoSuchSymbol, name)
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
		s.types[name] = compiler.TNUMBER
	case "Money":
		s.types[name] = compiler.TMONEY
	default:
		s.types[name] = compiler.TDYNAMIC
	}
}
