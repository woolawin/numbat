package validator

import (
	"numbat/internal/common"
	"testing"
)

func TestProgramCanNotBeNull(t *testing.T) {
	code := `
proc main do
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, &ProgramingMissing{})
	assertValidationErrorCount(t, validation, 1)
}

func loc(line, column int) common.Location {
	return common.Location{
		Line:   line,
		Column: column,
	}
}
