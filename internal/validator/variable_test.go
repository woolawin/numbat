package validator

import (
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
