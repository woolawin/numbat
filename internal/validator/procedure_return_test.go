package validator

import (
	"testing"
)

func TestCanNotReturnOnProcedureWithNoOutType(t *testing.T) {
	code := `
program do
end

proc banana do
	return 1
end
`
	validation := NewValidation()
	validation.Validate(readsrc(code))

	assertValidationError(t, validation, NewDoesNotReturnValue("banana", loc(6, 9)))
	assertValidationErrorCount(t, validation, 1)
}

func TestCanNotReturnAValueFromProgram(t *testing.T) {
	code := `
program do
	return 0
end
`
	validation := NewValidation()
	validation.Validate(readsrc(code))

	assertValidationError(t, validation, NewDoesNotReturnValue("program", loc(3, 9)))
	assertValidationErrorCount(t, validation, 1)
}

func TestCanNotReturnNoValueIfProcedureHasReturnType(t *testing.T) {
	code := `
program do
end

proc banana Int32 do
	return
end
`
	validation := NewValidation()
	validation.Validate(readsrc(code))

	assertValidationError(t, validation, NewReturnValueRequired("banana", loc(6, 2), "Int32"))
	assertValidationErrorCount(t, validation, 1)
}

func TestProcedureWithReturnTypeMustReturnAValue(t *testing.T) {
	code := `
program do
end

proc banana Int32 do
end
`
	validation := NewValidation()
	validation.Validate(readsrc(code))

	assertValidationError(t, validation, NewMissingReturnStatement("banana", loc(5, 6)))
	assertValidationErrorCount(t, validation, 1)
}

func TestProcedureWithReturnTypeDoesNotNeedReturnStatement(t *testing.T) {
	code := `
program do
end

proc banana do
end

proc apple(a Bool) do

end
`
	validation := NewValidation()
	validation.Validate(readsrc(code))

	assertValidationErrorCount(t, validation, 0)
}
