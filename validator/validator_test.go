package validator

import "testing"

func TestProgramCanNotBeNull(t *testing.T) {
	code := `
proc main do
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.HasProgram(src)

	assertValidationError(t, validation, ProgramingMissing{})
}
