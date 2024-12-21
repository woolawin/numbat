package validator

import (
	"encoding/json"
	"numbat/internal/common"
	"numbat/internal/read"
	"reflect"
	"testing"
)

func TestInferTypes(t *testing.T) {
	code := `
program do
	var b = 1
	var c = 1.2
	var e = true
	var f = 0xFF
	var g = 1e+22
end
`
	src := readsrc(code)
	validation := NewValidation()
	p := validation.Validate(src)

	assertInferredType(t, p, "b", "Int32")
	assertInferredType(t, p, "c", "Float64")
	assertInferredType(t, p, "e", "Bool")
	assertInferredType(t, p, "f", "Byte")
	assertInferredType(t, p, "g", "Float64")
}

func assertInferredType(t *testing.T, src *common.Source, name string, expected string) {
	for _, stmt := range src.Program.Statements {
		switch stmt := stmt.(type) {
		case common.VariableDeclaration:
			{
				if stmt.Name.Value == name {
					typ := stmt.Type
					actual := typ.String()
					if actual != expected {
						t.Fatalf("Inferred type of %s should be %s but is %s", name, expected, actual)
					}
					return
				}
			}
		}
	}
	t.Fatalf("did not find variable %s", name)
}

func TestInferTypesErrs(t *testing.T) {
	code := `
program do
	var foo Int32

	var a
	var b = null
	var c = foo
	var d = baz()
end

proc baz Int32 do
	return 0
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewNoExprToInferVariableType(name("a", 5, 6)))
	assertValidationError(t, validation, NewCanNotInferTypeFromNull(loc(6, 10)))
	assertValidationError(t, validation, NewCanNotInferTypeFromOtherVariable(loc(7, 10)))
	assertValidationError(t, validation, NewCanNotInferTypeFromCall(loc(8, 10)))
	assertValidationErrorCount(t, validation, 4)
}

func assertValidationError(t *testing.T, validation Validation, err ValidationError) {
	for _, e := range validation.errors {
		if reflect.DeepEqual(e, err) {
			return
		}
	}
	b, _ := json.Marshal(err)
	t.Fatalf("did not find validation error %s", string(b))
}

func assertValidationErrorCount(t *testing.T, validation Validation, expected int) {
	if len(validation.errors) != expected {
		t.Fatalf("did not find %d validation errors, found %d", expected, len(validation.errors))
	}
}

func readsrc(sample string) *read.Source {
	reader := read.NewSourceReader()
	reader.Read(sample, "")
	return reader.Source()
}
