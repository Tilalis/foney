package interpreter

import "reflect"

import "fmt"

// ByteCodeInstruction represents instruction
type ByteCodeInstruction int

// Bytecode
const (
	NOP   ByteCodeInstruction = iota
	START ByteCodeInstruction = iota

	PUSH ByteCodeInstruction = iota // PUSH {symbol | value}

	PUSHM ByteCodeInstruction = iota // PUSH Money
	PUSHF ByteCodeInstruction = iota // PUSH Number

	POP ByteCodeInstruction = iota // POP {symbol}
	SET ByteCodeInstruction = iota // SET {symbol}

	ADD  ByteCodeInstruction = iota // everything `op` everything
	SUB  ByteCodeInstruction = iota
	MULX ByteCodeInstruction = iota
	DIVX ByteCodeInstruction = iota

	ADDFF ByteCodeInstruction = iota // number + number (on top of stack)
	ADDMM ByteCodeInstruction = iota // money + money

	SUBFF ByteCodeInstruction = iota // number - number
	SUBMM ByteCodeInstruction = iota // money - money

	MULFF ByteCodeInstruction = iota // number * number
	MULMF ByteCodeInstruction = iota // money * number
	MULFM ByteCodeInstruction = iota // number * money

	DIVFF ByteCodeInstruction = iota // number / number
	DIVMF ByteCodeInstruction = iota // money / number

	// Not supported
	CONVERT ByteCodeInstruction = iota // money `convert` currency
)

func (b ByteCodeInstruction) String() string {
	instructionName := map[ByteCodeInstruction]string{
		NOP:   "NOP",
		START: "START",

		PUSH:  "PUSH",
		PUSHF: "PUSHF",
		PUSHM: "PUSHM",

		POP: "POP",

		ADD:   "ADD",
		ADDFF: "ADDFF",
		ADDMM: "ADDMM",

		SUB: "SUB",

		SUBFF: "SUBFF",
		SUBMM: "SUBMM",

		MULFF: "MULFF",
		MULFM: "MULFM",
		MULMF: "MULMF",

		DIVFF: "DIVFF",
		DIVMF: "DIVMF",

		MULX: "MUL",
		DIVX: "DIV",

		SET: "SET",
	}[b]

	return instructionName
}

// TypeInfo type
type TypeInfo int

// TypeInfo
const (
	NUMBERTYPE TypeInfo = iota
	MONEYTYPE  TypeInfo = iota
	DYNAMIC    TypeInfo = iota
)

func (ti TypeInfo) String() string {
	name := map[TypeInfo]string{
		NUMBERTYPE: "Number",
		MONEYTYPE:  "Money",
		DYNAMIC:    "<Dynamic>",
	}[ti]

	return name
}

// TypeInfoList type
type TypeInfoList struct {
	list []TypeInfo
	size int
}

// Get TypeInfo
func (til *TypeInfoList) Get() (a, b TypeInfo, err error) {
	if til.size < 2 {
		err = fmt.Errorf("%w of TypeInfoList", ErrEmptyStack)
		return
	}

	a, b = til.list[til.size-1], til.list[til.size-2]

	til.list = til.list[:til.size-2]
	til.size = til.size - 2

	return
}

// Put TypeInfo
func (til *TypeInfoList) Put(info TypeInfo) {
	til.list = append(til.list, info)
	til.size++
}

// InstructionTypeInfo ???
var InstructionTypeInfo = &TypeInfoList{
	list: make([]TypeInfo, 0),
	size: 0,
}

// Instruction struct
type Instruction struct {
	Instruction ByteCodeInstruction
	Argument    interface{}
	next        *Instruction
	offset      int
}

// NewInstruction creates instruction
func NewInstruction(byteCode ByteCodeInstruction, argument interface{}) *Instruction {
	if argument != nil {
		var typeOf = reflect.Indirect(reflect.ValueOf(argument)).Type()
		var name = typeOf.Name()

		switch name {
		case "float64":
			InstructionTypeInfo.Put(NUMBERTYPE)
		case "Money":
			InstructionTypeInfo.Put(MONEYTYPE)
		case "string": // Symbol
			// TODO: Error
			symbolType, _ := symbolTable.GetType(argument.(string))
			InstructionTypeInfo.Put(symbolType)
		}
	}

	return &Instruction{
		Instruction: byteCode,
		Argument:    argument,
		next:        nil,
	}
}

// Next Instruction
func (b *Instruction) Next() *Instruction {
	return b.next
}

// Append bytecode instruction
func (b *Instruction) Append(a *Instruction) *Instruction {
	b.next = a
	a.offset = b.offset + 1
	return a
}

// Load code to slice
func (b *Instruction) Load() []*Instruction {
	var loaded = make([]*Instruction, 0)

	for b := b.Next(); b != nil; b = b.Next() {
		loaded = append(loaded, b)
	}

	return loaded
}
