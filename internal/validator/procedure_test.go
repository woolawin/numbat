package validator

import (
	. "numbat/internal/common"
	"testing"
)

func TestParametersAreAddedToProcedure(t *testing.T) {
	code := `
proc apple(a Int32) Bool do
	return true
end
`
	validation := NewValidation()
	src := validation.Validate(readsrc(code))

	proc := src.Procedures[0]

	if proc.Name.Value != "apple" {
		t.Errorf("procedure name not apple")
	}

	expectedType := inoutType(NewInt32Type(), NewBoolType())
	if typesAreNotEqual(proc.Type, expectedType) {
		t.Errorf("incorrect procedure type `%s`\n", proc.Type.String())
	}

	if len(proc.Parameters) != 1 {
		t.Fatal("procedure parameters count not 1")
	}

	param := proc.Parameters[0]

	if param.Name.Value != "a" {
		t.Errorf("parameter name not a")
	}

	if typesAreNotEqual(param.Type, outType(NewInt32Type())) {
		t.Errorf("incorrect parameter type `%s`\n", param.Type.String())
	}
}
