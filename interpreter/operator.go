package interpreter

import (
	"fmt"
	money "foney/money"
)

// BinaryOperator represents BinaryOperator
type BinaryOperator struct {
	Left     AST
	Operator *Token
	Right    AST
}

func (bo *BinaryOperator) moneyAndMoney(left, right *money.Money) (*money.Money, error) {
	switch bo.Operator.Type {
	case PLUS:
		return left.Add(right)
	case MINUS:
		return left.Sub(right)
	}
	return nil, fmt.Errorf("%w: %v between %v and %v", ErrOperationNotDefined, bo.Operator.Value, left, right)
}

func (bo *BinaryOperator) moneyAndNumber(left *money.Money, right float64) (*money.Money, error) {
	switch bo.Operator.Type {
	case MUL:
		return left.Mul(right)
	case DIV:
		return left.Div(right)
	}

	return nil, fmt.Errorf("%w: %v between %v and %v", ErrOperationNotDefined, bo.Operator.Value, left, right)
}

func (bo *BinaryOperator) numberAndNumber(left, right float64) (float64, error) {
	switch bo.Operator.Type {
	case PLUS:
		return left + right, nil
	case MINUS:
		return left - right, nil
	case MUL:
		return left * right, nil
	case DIV:
		if right == 0 {
			return 0, ErrDivisionByZero
		}

		return left / right, nil
	}

	return 0, fmt.Errorf("%w: %v between %v and %v", ErrOperationNotDefined, bo.Operator.Value, left, right)
}

// Traverse implements AST interface
func (bo *BinaryOperator) Traverse() (interface{}, error) {
	leftValue, err := bo.Left.Traverse()
	if err != nil {
		return nil, err
	}

	rightValue, err := bo.Right.Traverse()
	if err != nil {
		return nil, err
	}

	var (
		leftMoney, rightMoney *money.Money
		leftFloat, rightFloat float64

		leftIsMoney, rightIsMoney bool
		leftIsFloat, rightIsFloat bool
	)

	leftMoney, leftIsMoney = leftValue.(*money.Money)
	if !leftIsMoney {
		leftFloat, leftIsFloat = leftValue.(float64)
		if !leftIsFloat {
			return nil, fmt.Errorf("%w: %T", ErrUnsupportedType, leftValue)
		}
	}

	rightMoney, rightIsMoney = rightValue.(*money.Money)
	if !rightIsMoney {
		rightFloat, rightIsFloat = rightValue.(float64)
		if !rightIsFloat {
			return nil, fmt.Errorf("%w: %T", ErrUnsupportedType, rightValue)
		}
	}

	if leftIsMoney && rightIsMoney {
		return bo.moneyAndMoney(leftMoney, rightMoney)
	}

	if leftIsMoney && rightIsFloat {
		return bo.moneyAndNumber(leftMoney, rightFloat)
	}

	if leftIsFloat && rightIsMoney {
		return bo.moneyAndNumber(rightMoney, leftFloat)
	}

	if leftIsFloat && rightIsFloat {
		return bo.numberAndNumber(leftFloat, rightFloat)
	}

	return nil, fmt.Errorf("%w: %v between %v and %v", ErrUnsupportedOperation, bo.Operator.Value, leftValue, rightValue)
}

// Compile binary operator
func (bo *BinaryOperator) Compile(b *Instruction) (*Instruction, error) {
	// Order is important!
	b, err := bo.Right.Compile(b)

	if err != nil {
		return nil, err
	}

	b, err = bo.Left.Compile(b)

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

	return b.Append(NewInstruction(instruction, nil)), nil
}
