package compiler

// AST Abstract Syntax Tree
type AST interface {
	Traverse() (*Instruction, error)
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

// Traverse Number to Instruction
func (n *Number) Traverse() (*Instruction, error) {
	return NewInstruction(PUSHF, n.Token.Value), nil
}

// Money represents Money
type Money struct {
	Value
}

// Traverse Money to Instruction
func (m *Money) Traverse() (*Instruction, error) {
	return NewInstruction(PUSHM, m.Token.Value), nil
}

// Symbol represents Symbol in AST
type Symbol struct {
	Token *Token
}

// Traverse Symbol
func (s *Symbol) Traverse() (*Instruction, error) {
	symbolName, ok := s.Token.Value.(string)

	if !ok {
		// TODO: fix error
		return nil, ErrSyntaxError
	}

	return NewInstruction(PUSH, symbolName), nil
}
