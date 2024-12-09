package validator

import (
	"numbat/read"
)

type UnknownType struct {
	TypeName string
}

type ValidationError interface {
}

type Validation struct {
	errors []ValidationError
}

func Check(source *read.Source) Validation {
	validation := Validation{}

	if source.Program != nil {
		checkProc(&validation, *source.Program)
	}

	for _, proc := range source.Procs {
		checkProc(&validation, proc)
	}

	return validation
}

func checkProc(validation *Validation, proc read.Proc) {
	if proc.ReturnType != nil {
		checkType(validation, *proc.ReturnType)
	}
	for _, stmt := range proc.Statements {
		checkStatement(validation, stmt)
	}
}

func checkStatement(validation *Validation, statement read.Statement) {
	if statement.Var != nil {
		checkVar(validation, *statement.Var)
	}
}

func checkVar(validation *Validation, varStmt read.Var) {
	if varStmt.VarType != nil {
		checkType(validation, *varStmt.VarType)
	}
}

func checkType(validation *Validation, typ read.Type) {
	if isUnknownType(typ.Out.Name) {
		validation.errors = append(validation.errors, UnknownType{TypeName: typ.Out.Name})
	}
	for _, in := range typ.In {
		checkType(validation, *in.Typ)
	}
}

func isUnknownType(name string) bool {
	if name == "" {
		return false
	}
	return !isBuiltInType(name)
}

func isBuiltInType(typeName string) bool {
	var types = [...]string{
		"Int16",
		"Int32",
		"Int64",
		"Float32",
		"Float64",
		"Bool",
		"Str",
		"Byte",
	}

	for _, t := range types {
		if typeName == t {
			return true
		}
	}

	return false
}
