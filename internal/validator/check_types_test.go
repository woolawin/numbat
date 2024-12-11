package validator

import (
	read2 "numbat/internal/read"
	"testing"
)

func TestTypes(t *testing.T) {
	src := &read2.Source{
		Program: &read2.Proc{
			Statements: []read2.Statement{
				{
					Var: &read2.Var{
						Name:    "a",
						VarType: &read2.Type{Out: read2.TypeOut{Name: "Int32"}},
					},
				},
				{
					Var: &read2.Var{
						Name:    "b",
						VarType: &read2.Type{Out: read2.TypeOut{Name: "Str"}},
					},
				},
				{
					Var: &read2.Var{
						Name:    "c",
						VarType: &read2.Type{Out: read2.TypeOut{Name: "FooBar"}},
					},
				},
				{
					Var: &read2.Var{
						Name:    "d",
						VarType: &read2.Type{Out: read2.TypeOut{Name: ""}},
					},
				},
			},
		},
		Procs: []read2.Proc{
			{
				Name:       "get_num",
				Statements: []read2.Statement{},
				ReturnType: &read2.Type{
					Out: read2.TypeOut{Name: "Int32"},
				},
			},
			{
				Name:       "foo",
				Statements: []read2.Statement{},
				ReturnType: &read2.Type{
					Out: read2.TypeOut{Name: "Foo"},
				},
			},
			{
				Name:       "multiply",
				Statements: []read2.Statement{},
				ReturnType: &read2.Type{
					In: []read2.Param{
						{
							Typ: &read2.Type{Out: read2.TypeOut{Name: "Int32"}},
						},
						{
							Typ: &read2.Type{Out: read2.TypeOut{Name: "IntX"}},
						},
					},
					Out: read2.TypeOut{Name: "Int64"},
				},
			},
		},
	}

	validation := NewValidation()
	validation.CheckTypes(src)

	assertValidationError(t, validation, UnknownType{TypeName: "FooBar"})
	assertValidationError(t, validation, UnknownType{TypeName: "Foo"})
	assertValidationError(t, validation, UnknownType{TypeName: "IntX"})
}
