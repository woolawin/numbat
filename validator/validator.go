package validator

import (
	"numbat/common"
	"numbat/read"
	"strings"
	"unicode"
)

type UnknownType struct {
	TypeName string
}

type UnInferrableVar struct {
	VarName string
}

type NoExprToInferVariableType struct {
	VarName string
}

type CanNotInferTypeFromNull struct {
	VarName string
}

type CanNotInferTypeFromOtherVariable struct {
	VarName string
}

type CanNotInferTypeFromCall struct {
	VarName string
}

type ValidationError interface {
}

type Validation struct {
	errors []ValidationError
}

func NewValidation() Validation {
	return Validation{}
}

func (validation *Validation) InferTypes(src *read.Source) {
	validation.inferProcTypes(src.Program)
	for idx := range src.Procs {
		proc := &src.Procs[idx]
		validation.inferProcTypes(proc)
	}
}

func (validation *Validation) inferProcTypes(proc *read.Proc) {
	for idx := range proc.Statements {
		stmt := &proc.Statements[idx]
		if stmt.Var == nil {
			continue
		}
		if stmt.Var.VarType != nil {
			continue
		}
		if len(stmt.Var.Exprs) == 0 {
			validation.addError(NoExprToInferVariableType{VarName: stmt.Var.Name})
			continue
		}
		exp := stmt.Var.Exprs[0]

		if exp.Unit == common.UnitFloat64 {
			stmt.Var.VarType = TypeOf(common.TypeFloat64)
			continue
		}

		if exp.Unit == common.UnitFloat32 {
			stmt.Var.VarType = TypeOf(common.TypeFloat32)
			continue
		}

		if exp.Unit == common.UnitInt32 {
			stmt.Var.VarType = TypeOf(common.TypeInt32)
			continue
		}

		if exp.Unit == common.UnitInt64 {
			stmt.Var.VarType = TypeOf(common.TypeInt64)
			continue
		}

		if exp.Unit == common.UnitUint32 {
			stmt.Var.VarType = TypeOf(common.TypeUint32)
			continue
		}

		if exp.Unit == common.UnitUint64 {
			stmt.Var.VarType = TypeOf(common.TypeUint64)
			continue
		}

		if exp.Null {
			validation.addError(CanNotInferTypeFromNull{VarName: stmt.Var.Name})
			continue
		}

		if exp.VarName != nil {
			validation.addError(CanNotInferTypeFromOtherVariable{VarName: stmt.Var.Name})
			continue
		}

		if exp.Hex != nil {
			stmt.Var.VarType = TypeOf(common.TypeByte)
			continue
		}

		if exp.Str != nil {
			stmt.Var.VarType = TypeOf(common.TypeStr)
			continue
		}

		if exp.Boolean != nil {
			stmt.Var.VarType = TypeOf(common.TypeBool)
			continue
		}

		if exp.Call != nil {
			if unicode.IsUpper([]rune(exp.Call.Primary)[0]) && strings.TrimSpace(exp.Call.Secondary) == "" {
				stmt.Var.VarType = TypeOf(exp.Call.Primary)
				continue
			}
			validation.addError(CanNotInferTypeFromCall{VarName: stmt.Var.Name})
			continue
		}

		if exp.Number == nil {
			// return error
			continue
		}

		if strings.Contains(*exp.Number, "e+") {
			stmt.Var.VarType = TypeOf(common.TypeFloat64)
			continue
		}

		if strings.Contains(*exp.Number, ".") {
			stmt.Var.VarType = TypeOf(common.TypeFloat64)
			continue
		}
		stmt.Var.VarType = TypeOf(common.TypeInt32)
		continue
	}
}

func TypeOf(name string) *read.Type {
	return &read.Type{Out: read.TypeOut{Name: name}}
}

func Check(source *read.Source) Validation {
	validation := Validation{}

	if source.Program != nil {
		validation.checkProcTypes(*source.Program)
	}

	for _, proc := range source.Procs {
		validation.checkProcTypes(proc)
	}

	return validation
}

func (validation *Validation) checkProcTypes(proc read.Proc) {
	if proc.ReturnType != nil {
		validation.checkType(*proc.ReturnType)
	}
	for _, stmt := range proc.Statements {
		validation.checkStatementTypes(stmt)
	}
}

func (validation *Validation) checkStatementTypes(statement read.Statement) {
	if statement.Var != nil {
		validation.checkVarType(*statement.Var)
	}
}

func (validation *Validation) checkVarType(varStmt read.Var) {
	if varStmt.VarType != nil {
		validation.checkType(*varStmt.VarType)
	}
}

func (validation *Validation) checkType(typ read.Type) {
	if isUnknownType(typ.Out.Name) {
		validation.errors = append(validation.errors, UnknownType{TypeName: typ.Out.Name})
	}
	for _, in := range typ.In {
		validation.checkType(*in.Typ)
	}
}

func (validation *Validation) addError(verr ValidationError) {
	validation.errors = append(validation.errors, verr)
}

func isUnknownType(name string) bool {
	if name == "" {
		return false
	}
	return !isBuiltInType(name)
}

func isBuiltInType(typeName string) bool {
	var types = [...]string{
		common.TypeInt32,
		common.TypeInt64,
		common.TypeUint32,
		common.TypeUint64,
		common.TypeFloat32,
		common.TypeFloat64,
		common.TypeBool,
		common.TypeStr,
		common.TypeStr,
	}

	for _, t := range types {
		if typeName == t {
			return true
		}
	}

	return false
}
