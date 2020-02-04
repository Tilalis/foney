package compiler

import (
	"fmt"
)

// BinaryOperator represents BinaryOperator
type BinaryOperator struct {
	Left     AST
	Operator *Token
	Right    AST
}

// Traverse binary operator
func (bo *BinaryOperator) Traverse() (*Instruction, error) {
	right, err := bo.Right.Traverse()

	if err != nil {
		return nil, err
	}

	left, err := bo.Left.Traverse()

	if err != nil {
		return nil, err
	}

	var instruction ByteCodeInstruction = NOP

	typeA, typeB, err := InstructionTypeInfo.Get()

	if err != nil {
		return nil, err
	}

	switch bo.Operator.Type {
	case PLUS:
		switch {
		case typeA == TNUMBER && typeB == TNUMBER:
			instruction = ADDFF
		case typeA == TMONEY && typeB == TMONEY:
			instruction = ADDMM
		}
	case MINUS:
		switch {
		case typeA == TNUMBER && typeB == TNUMBER:
			instruction = SUBFF
		case typeA == TMONEY && typeB == TMONEY:
			instruction = SUBMM
		}
	case MUL:
		switch {
		case typeA == TNUMBER && typeB == TNUMBER:
			instruction = MULFF
		case typeA == TNUMBER && typeB == TMONEY:
			instruction = MULFM
		case typeA == TMONEY && typeB == TNUMBER:
			instruction = MULMF
		}
	case DIV:
		switch {
		case typeA == TNUMBER && typeB == TNUMBER:
			instruction = DIVFF
		case typeA == TMONEY && typeB == TNUMBER:
			instruction = DIVMF
		}
	}

	if instruction == NOP {
		return nil, fmt.Errorf("%w: %v between %v and %v", ErrUnsupportedOperation, bo.Operator.Value, TNUMBER, TMONEY)
	}

	switch instruction {
	case ADDFF, SUBFF, MULFF, DIVFF:
		InstructionTypeInfo.Put(TNUMBER)
	default:
		InstructionTypeInfo.Put(TMONEY)
	}

	left.Append(NewInstruction(instruction, nil, EOF))
	right.Append(left)

	return right, nil
}
