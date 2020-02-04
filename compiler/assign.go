package compiler

// Assign represents assign operation
type Assign struct {
	Symbol *Symbol
	Value  AST
}

// Traverse traverse assign
func (a *Assign) Traverse() (*Instruction, error) {
	value, err := a.Value.Traverse()

	if err != nil {
		return nil, err
	}

	token := a.Symbol.Token
	name := token.Value.(string)

	InstructionTypeInfo.PutSymbolType(name, InstructionTypeInfo.Last())
	value.Append(NewInstruction(SET, name, token.Type))

	return value, nil
}
