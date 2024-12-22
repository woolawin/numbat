package validator

import (
	. "numbat/internal/common"
	"numbat/internal/read"
	"reflect"
	"testing"
)

func typesAreNotEqual(a, b InOutType) bool {
	if a.Out != b.Out {
		return true
	}
	if len(a.In) != len(b.In) {
		return true
	}

	var inTypeAreNotEqual = func(a, b InType) bool {
		return typesAreNotEqual(*a.Type, *b.Type)
	}

	for i := range a.In {
		if inTypeAreNotEqual(a.In[i], b.In[i]) {
			return true
		}
	}
	return false
}

func inoutType(in AtomicType, out AtomicType) InOutType {
	return NewInOutType(
		[]InType{
			NewInType(
				outType(in),
				Name{Value: "", Location: Location{}},
				nil,
			),
		},
		NewSuperAtomicType(out),
	)
}

func outType(out AtomicType) InOutType {
	return NewInOutType(nil, NewSuperAtomicType(out))
}

func assertVariableType(t *testing.T, src *Source, name string, expected string) {
	for _, stmt := range src.Program.Statements {
		switch stmt := stmt.(type) {
		case VariableDeclaration:
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

func assertValidationError(t *testing.T, validation Validation, err ValidationError) {
	for _, e := range validation.errors {
		if reflect.DeepEqual(e, err) {
			return
		}
	}
	t.Fatalf("did not find validation error: %s", err.Message())
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
