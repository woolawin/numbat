package validator

import (
	"testing"
)

func TestRHSVariableUnknown(t *testing.T) {
	code := `
program do
end

proc foo do
	var banana Int32
	var orange Int32 = banana
	var pear Int32
	var apple Int32 = watermelon 	// <-- watermelon is not a know object
	var grape Int32 = pear
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewUnknownObject("watermelon", loc(9, 20)))
	assertValidationErrorCount(t, validation, 1)
}

func TestRHSProcUnknown(t *testing.T) {
	code := `
program do
end

proc foo do
	var banana Int32 = strawberry()
	var orange Int32 = blueberry()
	var pear Int32
end

proc strawberry Int32 do

end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewUnknownObject("blueberry", loc(7, 21)))
	assertValidationErrorCount(t, validation, 1)
}

func TestProcCallWithUnknownVariable(t *testing.T) {
	code := `
program do
end

proc foo do
	var pear Int32
	var berry Int32 = strawberry(pear, true)
	var apple Int32 = strawberry(orange, true)		// <-- orange is not defined
	strawberry(banana, false)						// <-- banana is not defined
	strawberry(apple, false)
end

proc strawberry(item Int32, repeat Bool) Int32 do
	return 1
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewUnknownObject("orange", loc(8, 31)))
	assertValidationError(t, validation, NewUnknownObject("banana", loc(9, 13)))
	assertValidationErrorCount(t, validation, 2)
}

func TestCanNotForwardReferenceVariable(t *testing.T) {
	code := `
program do
	var banana Int32 = orange
	var orange = 1
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewUnknownObject("orange", loc(3, 21)))
	assertValidationErrorCount(t, validation, 1)
}

func TestCanReferenceParameter(t *testing.T) {
	code := `
program do
end

proc foo(apple Int64) do
	var pear Int64 = apple
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationErrorCount(t, validation, 0)
}

func TestCanReferenceProcedures(t *testing.T) {
	code := `
program do
	foo(4)
end

proc grape Int32 do

end

proc foo(apple Int32) do
	var pear Int32 = grape()
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationErrorCount(t, validation, 0)
}
