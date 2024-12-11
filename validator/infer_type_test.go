package validator

import (
	"numbat/read"
	"testing"
)

func TestInferTypes(t *testing.T) {
	code := `
program do
	var a = "string"
	var b = 1
	var c = 1.2
	var d = #i64 2
	var e = true
	var f = 0xFF
	var g = 1e+22
	var h = #f32 3
	var i = #i32 0
	var j = #f64 1
	var u32 = #u32 1
	var u64 = #u64 1
	var k Int32 = #f64 4
	var l = (Link)
	var m = (Metric: "request_count")
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.InferTypes(src)

	assertInferredType(t, src, "a", "Str")
	assertInferredType(t, src, "b", "Int32")
	assertInferredType(t, src, "c", "Float64")
	assertInferredType(t, src, "d", "Int64")
	assertInferredType(t, src, "e", "Bool")
	assertInferredType(t, src, "f", "Byte")
	assertInferredType(t, src, "g", "Float64")
	assertInferredType(t, src, "h", "Float32")
	assertInferredType(t, src, "i", "Int32")
	assertInferredType(t, src, "j", "Float64")
	assertInferredType(t, src, "u32", "Uint32")
	assertInferredType(t, src, "u64", "Uint64")
	assertInferredType(t, src, "k", "Int32")
	assertInferredType(t, src, "l", "Link")
	assertInferredType(t, src, "m", "Metric")
}

func assertInferredType(t *testing.T, src *read.Source, name string, expected string) {
	for _, stmt := range src.Program.Statements {
		if stmt.Var.Name == name {
			if stmt.Var.VarType == nil {
				t.Fatalf("%s: expected type", name)
			}
			actual := stmt.Var.VarType.Out.Name
			if actual != expected {
				t.Fatalf("Inferred type of %s should be %s but is %s", name, expected, actual)
			}
			return
		}
	}
	t.Fatalf("did not find variable %s", name)
}

func TestInferTypesErrs(t *testing.T) {
	code := `
program do
	var a
	var b = null
	var c = foo
	var d = (baz bar)
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.InferTypes(src)

	assertValidationError(t, validation, NoExprToInferVariableType{VarName: "a"})
	assertValidationError(t, validation, CanNotInferTypeFromNull{VarName: "b"})
	assertValidationError(t, validation, CanNotInferTypeFromOtherVariable{VarName: "c"})
	assertValidationError(t, validation, CanNotInferTypeFromCall{VarName: "d"})
}

func assertValidationError(t *testing.T, validation Validation, err ValidationError) {
	for _, e := range validation.errors {
		if e == err {
			return
		}
	}
	t.Fatalf("did not find validation error %s", err)
}

func readsrc(sample string) *read.Source {
	listener := read.NewListener()
	listener.Exec(sample)
	return listener.Source()
}
