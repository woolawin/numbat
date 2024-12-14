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

func TestProcParamNameConflictsWithEachOther(t *testing.T) {
	code := `
program do
end

proc foo(bar Int32, bar Int64, baz Bool) do
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, newNameConflict(name("bar", 5, 21), loc(5, 10)))
	assertValidationErrorCount(t, validation, 1)
}

func TestProcParamNameConflictsWithOtherProc(t *testing.T) {
	code := `
program do
end

proc foo(bar Int32) do
end

proc bar do
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, newNameConflict(name("bar", 5, 10), loc(8, 6)))
	assertValidationErrorCount(t, validation, 1)
}

func TestVarConflictsWithParamAndProc(t *testing.T) {
	code := `
program do
end

proc foo(apple Int32) do
	var banana Int64
	var orange Int32	// <-- conflicts with orange parameter
	var pear Bool
	var apple Float32  	// <-- conflicts with apple proc
end

proc orange do
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, newNameConflict(name("orange", 7, 6), loc(12, 6)))
	assertValidationError(t, validation, newNameConflict(name("apple", 9, 6), loc(5, 10)))
	assertValidationErrorCount(t, validation, 2)
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
