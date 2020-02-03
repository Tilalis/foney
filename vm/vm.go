package vm

import (
	"github.com/Tilalis/foney/compiler"
	"github.com/Tilalis/foney/money"
)

var stack = make([]interface{}, 0)

// VMSymbolTable Global VM Symbol Table
var VMSymbolTable *SymbolTable = NewSymbolTable()

// InstructionHandler type
type InstructionHandler func(interface{}) error

var handlers = map[compiler.ByteCodeInstruction]InstructionHandler{
	compiler.PUSH:  push,
	compiler.PUSHF: push,
	compiler.PUSHM: push,

	compiler.ADDFF: addff,
	compiler.ADDMM: addmm,

	compiler.SUBFF: subff,
	compiler.SUBMM: submm,

	compiler.MULFF: mulff,
	compiler.MULFM: mulfm,
	compiler.MULMF: mulmf,

	compiler.DIVFF: divff,
	compiler.DIVMF: divmf,

	compiler.SET: set,
}

// Execute executes ByteCode
func Execute(instructions []*compiler.Instruction) (interface{}, error) {
	for _, instruction := range instructions {
		handler, ok := handlers[instruction.Instruction]

		if !ok {
			return nil, compiler.ErrUnsupportedOperation
		}

		handler(instruction.Argument)
	}

	return _pop(), nil
}

func push(argument interface{}) error {
	var (
		value = argument
		err   error
	)

	name, ok := argument.(string)

	if ok {
		value, err = VMSymbolTable.Get(name)

		if err != nil {
			return err
		}
	}

	stack = append(stack, value)

	return nil
}

func set(argument interface{}) error {
	name, ok := argument.(string)

	if !ok {
		return compiler.ErrNoSuchSymbol
	}

	stackSize := len(stack)
	element := stack[stackSize-1]

	VMSymbolTable.Set(name, element)

	return nil
}

func addff(_ interface{}) error {
	right := _pop().(float64)
	left := _pop().(float64)

	return push(left + right)
}

func addmm(_ interface{}) error {
	right := _pop().(*money.Money)
	left := _pop().(*money.Money)

	result, err := right.Add(left)

	if err != nil {
		return err
	}

	return push(result)
}

func subff(_ interface{}) error {
	right := _pop().(float64)
	left := _pop().(float64)

	return push(right - left)
}

func submm(_ interface{}) error {
	right := _pop().(*money.Money)
	left := _pop().(*money.Money)

	result, err := right.Sub(left)

	if err != nil {
		return err
	}

	return push(result)
}

func mulff(_ interface{}) error {
	right := _pop().(float64)
	left := _pop().(float64)

	return push(right * left)
}

func mulfm(_ interface{}) error {
	right := _pop().(float64)
	left := _pop().(*money.Money)

	result, err := left.Mul(right)

	if err != nil {
		return err
	}

	return push(result)
}

func mulmf(_ interface{}) error {
	right := _pop().(*money.Money)
	left := _pop().(float64)

	result, err := right.Mul(left)

	if err != nil {
		return err
	}

	return push(result)
}

func divff(_ interface{}) error {
	right := _pop().(float64)
	left := _pop().(float64)

	return push(right / left)
}

func divmf(_ interface{}) error {
	right := _pop().(*money.Money)
	left := _pop().(float64)

	result, err := right.Div(left)

	if err != nil {
		return err
	}

	return push(result)
}

func _pop() interface{} {
	size := len(stack)

	if size == 0 {
		return nil
	}

	result := stack[size-1]
	stack = stack[:size-1]

	return result
}
