package validator

import (
	"numbat/internal/common"
	"numbat/internal/read"
	"testing"
)

func TestDuplicateProcNames(t *testing.T) {
	code := `
program do
end

proc foo do
end

proc bar do
end

proc foo (a Int32) do

end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, newNameConflict(name("foo", 11, 6), loc(5, 6)))
	assertValidationErrorCount(t, validation, 1)
}

func name(name string, line, column int) read.Name {
	return read.Name{
		Value: name,
		Location: common.Location{
			Line:   line,
			Column: column,
		},
	}
}
