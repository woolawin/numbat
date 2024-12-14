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

	assertValidationError(t, validation, newUnknownObject("watermelon", loc(9, 20)))
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

	assertValidationError(t, validation, newUnknownObject("blueberry", loc(7, 21)))
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

	assertValidationError(t, validation, newUnknownObject("orange", loc(8, 31)))
	assertValidationError(t, validation, newUnknownObject("banana", loc(9, 13)))
	assertValidationErrorCount(t, validation, 2)
}
