package interpreter

// AST Abstract Syntax Tree
type AST interface {
	Traverse() (interface{}, error)
	Compile(b *Instruction) (*Instruction, error)
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

// Compile Number to Instruction
func (n *Number) Compile(b *Instruction) (*Instruction, error) {
	return b.Append(NewInstruction(PUSHF, n.Token.Value)), nil
}

// Money represents Money
type Money struct {
	Value
}

// Compile Money to Instruction
func (m *Money) Compile(b *Instruction) (*Instruction, error) {
	return b.Append(NewInstruction(PUSHM, m.Token.Value)), nil
}

// Symbol represents Symbol in AST
type Symbol struct {
	Token *Token
}

// Traverse Symbol
func (s *Symbol) Traverse() (interface{}, error) {
	symbolName := s.Token.Value.(string)
	return symbolTable.Get(symbolName)
}

// Compile Symbol
func (s *Symbol) Compile(b *Instruction) (*Instruction, error) {
	symbolName, ok := s.Token.Value.(string)

	if !ok {
		// TODO: fix error
		return nil, ErrSyntaxError
	}

	return b.Append(NewInstruction(PUSH, symbolName)), nil
}
