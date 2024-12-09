package validator

import (
	"github.com/google/go-cmp/cmp"
	"numbat/read"
	"testing"
)

func TestTypes(t *testing.T) {
	src := &read.Source{
		Program: &read.Proc{
			Statements: []read.Statement{
				{
					Var: &read.Var{
						Name:    "a",
						VarType: &read.Type{Out: read.TypeOut{Name: "Int32"}},
					},
				},
				{
					Var: &read.Var{
						Name:    "b",
						VarType: &read.Type{Out: read.TypeOut{Name: "Str"}},
					},
				},
				{
					Var: &read.Var{
						Name:    "c",
						VarType: &read.Type{Out: read.TypeOut{Name: "FooBar"}},
					},
				},
				{
					Var: &read.Var{
						Name:    "d",
						VarType: &read.Type{Out: read.TypeOut{Name: ""}},
					},
				},
			},
		},
		Procs: []read.Proc{
			{
				Name:       "get_num",
				Statements: []read.Statement{},
				ReturnType: &read.Type{
					Out: read.TypeOut{Name: "Int32"},
				},
			},
			{
				Name:       "foo",
				Statements: []read.Statement{},
				ReturnType: &read.Type{
					Out: read.TypeOut{Name: "Foo"},
				},
			},
			{
				Name:       "multiply",
				Statements: []read.Statement{},
				ReturnType: &read.Type{
					In: []read.Param{
						{
							Typ: &read.Type{Out: read.TypeOut{Name: "Int32"}},
						},
						{
							Typ: &read.Type{Out: read.TypeOut{Name: "IntX"}},
						},
					},
					Out: read.TypeOut{Name: "Int64"},
				},
			},
		},
	}

	actual := Check(src)

	expected := Validation{
		errors: []ValidationError{
			UnknownType{
				TypeName: "FooBar",
			},
			UnknownType{
				TypeName: "Foo",
			},
			UnknownType{
				TypeName: "IntX",
			},
		},
	}

	assert(t, actual, expected)
}

func assert(t *testing.T, actual Validation, expected Validation) {
	diff := cmp.Diff(actual.errors, expected.errors)
	if diff != "" {
		t.Fatalf(diff)
	}
}
