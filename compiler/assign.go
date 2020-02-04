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

	name := a.Symbol.Token.Value.(string)

	InstructionTypeInfo.PutSymbolType(name, InstructionTypeInfo.Last())
	value.Append(NewInstruction(SET, name))

	return value, nil
}
