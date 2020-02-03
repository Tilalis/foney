package interpreter

// AST Abstract Syntax Tree
type AST interface {
	Traverse() (interface{}, error)
}

// Value represents values in AST
type Value struct {
	Token *Token
}

// Traverse Value
func (v *Value) Traverse() (interface{}, error) {
	return v.Token.Value, nil
}

// Number represents Number
type Number struct {
	Value
}

// Money represents Money
type Money struct {
	Value
}

// Symbol represents Symbol in AST
type Symbol struct {
	Token *Token
}

// Traverse Symbol
func (s *Symbol) Traverse() (interface{}, error) {
	symbolName := s.Token.Value.(string)
	return globalSymbolTable.Get(symbolName)
}
