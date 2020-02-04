package compiler

import "fmt"

// InstructionTypeInfo ???
var InstructionTypeInfo = &TypeInfoList{
	list: make([]TypeInfo, 0),
	vars: make(map[string]TypeInfo, 0),
	size: 0,
}

// TypeInfoList type
type TypeInfoList struct {
	list []TypeInfo
	vars map[string]TypeInfo
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

// Last TypeInfo
func (til *TypeInfoList) Last() TypeInfo {
	if til.size < 1 {
		return DYNAMIC
	}

	return til.list[til.size-1]
}

// GetSymbolType get symbol type
func (til *TypeInfoList) GetSymbolType(name string) (TypeInfo, error) {
	typeInfo, ok := til.vars[name]

	if !ok {
		return DYNAMIC, ErrNoSuchSymbol
	}

	return typeInfo, nil
}

// PutSymbolType put symbol type
func (til *TypeInfoList) PutSymbolType(name string, ti TypeInfo) {
	til.vars[name] = ti
}

// TypeInfo type
type TypeInfo int

// TypeInfo
const (
	NUMBERTYPE TypeInfo = iota
	MONEYTYPE
	DYNAMIC
)

func (ti TypeInfo) String() string {
	name := map[TypeInfo]string{
		NUMBERTYPE: "Number",
		MONEYTYPE:  "Money",
		DYNAMIC:    "<Dynamic>",
	}[ti]

	return name
}
