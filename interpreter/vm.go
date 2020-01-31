package interpreter

import "foney/money"

var stack = make([]interface{}, 0)

// VirtualMachine represents vm
type VirtualMachine struct {
	ip *Instruction
	sp int
}

// NewVM returns new VM
func NewVM(start *Instruction) *VirtualMachine {
	return &VirtualMachine{
		ip: start,
		sp: 0,
	}
}

// InstructionHandler type
type InstructionHandler func(interface{}) error

var handlers = map[ByteCodeInstruction]InstructionHandler{
	PUSH:  push,
	PUSHF: push,
	PUSHM: push,

	ADDFF: addff,
	ADDMM: addmm,

	SUBFF: subff,
	SUBMM: submm,

	MULFF: mulff,
	MULFM: mulfm,
	MULMF: mulmf,

	DIVFF: divff,
	DIVMF: divmf,

	SET: set,
}

// Execute executes ByteCode
func (vm *VirtualMachine) Execute() (interface{}, error) {
	for instruction := vm.ip.Next(); instruction != nil; instruction = instruction.Next() {
		handler, ok := handlers[instruction.Instruction]

		if !ok {
			return nil, ErrUnsupportedOperation
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
		value, err = symbolTable.Get(name)

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
		return ErrNoSuchSymbol
	}

	stackSize := len(stack)
	element := stack[stackSize-1]

	symbolTable.Set(name, element)

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
