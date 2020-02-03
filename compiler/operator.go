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
		case typeA == NUMBERTYPE && typeB == NUMBERTYPE:
			instruction = ADDFF
		case typeA == MONEYTYPE && typeB == MONEYTYPE:
			instruction = ADDMM
		}
	case MINUS:
		switch {
		case typeA == NUMBERTYPE && typeB == NUMBERTYPE:
			instruction = SUBFF
		case typeA == MONEYTYPE && typeB == MONEYTYPE:
			instruction = SUBMM
		}
	case MUL:
		switch {
		case typeA == NUMBERTYPE && typeB == NUMBERTYPE:
			instruction = MULFF
		case typeA == NUMBERTYPE && typeB == MONEYTYPE:
			instruction = MULFM
		case typeA == MONEYTYPE && typeB == NUMBERTYPE:
			instruction = MULMF
		}
	case DIV:
		switch {
		case typeA == NUMBERTYPE && typeB == NUMBERTYPE:
			instruction = DIVFF
		case typeA == MONEYTYPE && typeB == NUMBERTYPE:
			instruction = DIVMF
		}
	}

	if instruction == NOP {
		return nil, fmt.Errorf("%w: %v between %v and %v", ErrUnsupportedOperation, bo.Operator.Value, NUMBERTYPE, MONEYTYPE)
	}

	switch instruction {
	case ADDFF, SUBFF, MULFF, DIVFF:
		InstructionTypeInfo.Put(NUMBERTYPE)
	default:
		InstructionTypeInfo.Put(MONEYTYPE)
	}

	left.Append(NewInstruction(instruction, nil))
	right.Append(left)

	return right, nil
}
