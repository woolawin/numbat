package validator

import (
	"testing"
)

func TestTypes(t *testing.T) {
	code := `
program do
	var a Int32
	var b Bool
	var c FooBar
end

proc foo(f Foo) do

end

proc multiply(num Int32) IntX do

end
`
	validation := NewValidation()
	validation.Validate(readsrc(code))

	assertValidationError(t, validation, NewUnknownType(name("FooBar", 5, 8)))
	assertValidationError(t, validation, NewUnknownType(name("Foo", 8, 12)))
	assertValidationError(t, validation, NewUnknownType(name("IntX", 12, 26)))
	assertValidationErrorCount(t, validation, 3)
}
