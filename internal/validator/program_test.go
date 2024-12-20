package validator

import (
	. "numbat/internal/common"
	"testing"
)

func TestMustContainProgram(t *testing.T) {
	code := `
proc apple(a Int32) do

end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewMissingProgram())
	assertValidationErrorCount(t, validation, 1)
}

func TestCanNotContainMoreThanTwoPrograms(t *testing.T) {
	code := `
program do

end

proc apple(a Int32) do

end

program do

end

program do

end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewProgramRedefined(loc(10, 0)))
	assertValidationError(t, validation, NewProgramRedefined(loc(14, 0)))
	assertValidationErrorCount(t, validation, 2)
}

func loc(line, column int) Location {
	return Location{
		Line:   line,
		Column: column,
	}
}
