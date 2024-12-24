package validator

import (
	"fmt"
	. "numbat/internal/common"
	"testing"
)

func TestExpressionDoesNotOverWriteVariableType(t *testing.T) {
	code := `
program do
	var b Int64 = 1
	var c Float32 = 1.2
end
`
	src := readsrc(code)
	validation := NewValidation()
	p := validation.Validate(src)

	assertVariableType(t, p, "b", "Int64")
	assertVariableType(t, p, "c", "Float32")
}

func TestInferTypes(t *testing.T) {
	code := `
program do
	var b = 1
	var c = 1.2
	var e = true
	var f = 0xFF
	var g = 1e+22
end
`
	src := readsrc(code)
	validation := NewValidation()
	p := validation.Validate(src)

	assertVariableType(t, p, "b", "Int32")
	assertVariableType(t, p, "c", "Float64")
	assertVariableType(t, p, "e", "Bool")
	assertVariableType(t, p, "f", "Byte")
	assertVariableType(t, p, "g", "Float64")
}

func TestInferTypesErrs(t *testing.T) {
	code := `
program do
	var foo Int32

	var a
	var b = null
	var c = foo
	var d = baz()
end

proc baz Int32 do
	return 0
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewNoExprToInferVariableType(name("a", 5, 6)))
	assertValidationError(t, validation, NewCanNotInferTypeFromNull(loc(6, 10)))
	assertValidationError(t, validation, NewCanNotInferTypeFromOtherVariable(loc(7, 10)))
	assertValidationError(t, validation, NewCanNotInferTypeFromCall(loc(8, 10)))
	assertValidationErrorCount(t, validation, 4)
}

func TestSequences(t *testing.T) {
	code := `
program do
	var a Int32[]
	var b Ascii[]
	var c Float64[10] 
end
`
	validation := NewValidation()
	src := validation.Validate(readsrc(code))

	for _, verr := range validation.errors {
		fmt.Println(verr.Message())
	}

	assertVariableType2(t, src, "a", NewSeqType(NewInt32Type(), 0))
	assertVariableType2(t, src, "b", NewSeqType(NewAsciiType(), 0))
	assertVariableType2(t, src, "c", NewSeqType(NewFloat64Type(), 10))
	assertValidationErrorCount(t, validation, 0)
}

func TestInvalidSequenceSizes(t *testing.T) {
	code := `
program do
	var a Int32[-1]
	var b Ascii[2.3]
	var c Float64[1e+22] 
	var d Bool[0] 
end
`
	validation := NewValidation()
	validation.Validate(readsrc(code))

	assertValidationError(t, validation, NewInvalidSequenceSize(loc(3, 13), "-1"))
	assertValidationError(t, validation, NewInvalidSequenceSize(loc(4, 13), "2.3"))
	assertValidationError(t, validation, NewInvalidSequenceSize(loc(5, 15), "1e+22"))
	assertValidationError(t, validation, NewInvalidSequenceSize(loc(6, 12), "0"))
	assertValidationErrorCount(t, validation, 4)
}
