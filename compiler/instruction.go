package compiler

import "reflect"

// ByteCodeInstruction represents instruction
type ByteCodeInstruction int

// Bytecode
const (
	NOP   ByteCodeInstruction = iota
	START ByteCodeInstruction = iota

	PUSH // PUSH {symbol | value}

	PUSHM // PUSH Money
	PUSHF // PUSH Number

	POP // POP {symbol}
	SET // SET {symbol}

	ADD // everything `op` everything
	SUB
	MULX
	DIVX

	ADDFF // number + number (on top of stack)
	ADDMM // money + money

	SUBFF // number - number
	SUBMM // money - money

	MULFF // number * number
	MULMF // money * number
	MULFM // number * money

	DIVFF // number / number
	DIVMF // money / number

	// Not supported
	CONVERT // money `convert` currency
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
			InstructionTypeInfo.Put(InstructionTypeInfo.Last())
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
	if b.next == nil {
		b.next = a
		a.offset = b.offset + 1
	} else {
		b.next.Append(a)
	}

	return a
}

// Load code to slice
func (b *Instruction) Load() []*Instruction {
	var loaded = make([]*Instruction, 0)

	for ; b != nil; b = b.Next() {
		loaded = append(loaded, b)
	}

	return loaded
}
