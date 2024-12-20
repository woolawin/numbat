package validator

import (
	. "numbat/internal/common"
	"testing"
)

func TestInvalidNumberOfArguments(t *testing.T) {
	code := `
program do
	apple(true)
	pear()
end

proc apple(a Int32) do

end

proc pear(a Bool) do

end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewIncompatibleType(NewBoolType(), NewInt32Type(), loc(3, 8)))
	assertValidationError(t, validation, NewIncorrectArgumentCount("pear", loc(4, 2), 0, 1))
	assertValidationErrorCount(t, validation, 2)
}
