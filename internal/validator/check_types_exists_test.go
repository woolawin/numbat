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
	validation.CheckTypesExists(readsrc(code))

	assertValidationError(t, validation, UnknownType{TypeName: "FooBar"})
	assertValidationError(t, validation, UnknownType{TypeName: "Foo"})
	assertValidationError(t, validation, UnknownType{TypeName: "IntX"})
	assertValidationErrorCount(t, validation, 3)
}
