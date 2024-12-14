package validator

import (
	"numbat/internal/common"
	"numbat/internal/read"
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

type ProgramingMissing struct {
}

type NameConflict struct {
	Name           read.Name
	ConflictedWith common.Location
}

func newNameConflict(name read.Name, conflictedWith common.Location) NameConflict {
	return NameConflict{Name: name, ConflictedWith: conflictedWith}
}

type UnknownObject struct {
	Name     string
	Location common.Location
}

func newUnknownObject(name string, location common.Location) UnknownObject {
	return UnknownObject{Name: name, Location: location}
}

//type TypeMismatch struct {
//	Location common.Location
//	Provided string
//	Required string
//}
//
//func newTypeMismatch(location common.Location, provided string, required string) TypeMismatch {
//	return TypeMismatch{Location: location, Provided: provided, Required: required}
//}

type ValidationError interface {
}

type Validation struct {
	errors []ValidationError
}

func NewValidation() Validation {
	return Validation{}
}

func Check(src *read.Source) Validation {
	validation := Validation{}
	validation.HasProgram(src)
	validation.InferTypes(src)
	validation.CheckTypesExists(src)
	//checkCollisions
	//checkReferences
	//checkTypeCompatability
	return validation
}

type Object struct {
	Name     read.Name
	Location common.Location
	Type     *read.Type
}

func newProcObject(proc *read.Proc) Object {
	return Object{Name: proc.Name, Location: proc.Location, Type: proc.Type}
}

func newParamObject(param read.Param) Object {
	return Object{Name: param.Name, Location: param.Name.Location, Type: param.Typ}
}

func newVarObject(stmt *read.Var) Object {
	return Object{Name: stmt.Name, Location: stmt.Name.Location, Type: stmt.VarType}
}

func (validation *Validation) Validate(src *read.Source) {
	objects := make(map[string]Object)
	for idx := range src.Procs {
		proc := &src.Procs[idx]
		object, found := objects[proc.Name.Value]
		if found {
			validation.addError(newNameConflict(proc.Name, object.Name.Location))
		} else {
			objects[proc.Name.Value] = newProcObject(proc)
		}
	}

	for idx := range src.Procs {
		proc := &src.Procs[idx]
		validation.validateProc(fork(objects), proc)
	}
}

func (validation *Validation) validateProc(objects map[string]Object, proc *read.Proc) {
	if proc.Type != nil {
		for _, param := range proc.Type.In {
			object, found := objects[param.Name.Value]
			if found {
				validation.addError(newNameConflict(param.Name, object.Name.Location))
			} else {
				objects[param.Name.Value] = newParamObject(param)
			}
		}
	}

	for idx := range proc.Statements {
		validation.validateStatement(objects, &proc.Statements[idx])
	}

}

func (validation *Validation) validateStatement(objects map[string]Object, stmt *read.Statement) {
	if stmt.Var != nil {
		validation.validateVar(objects, stmt.Var)
	}

	if stmt.Call != nil {
		validation.validateCall(objects, stmt.Call)
	}
}

func (validation *Validation) validateVar(objects map[string]Object, stmt *read.Var) {
	// Check if variable name is already taken
	object, found := objects[stmt.Name.Value]
	if found {
		validation.addError(newNameConflict(stmt.Name, object.Name.Location))
	} else {
		objects[stmt.Name.Value] = newVarObject(stmt)
	}

	if len(stmt.Exprs) == 1 {
		validation.validateExpr(objects, &stmt.Exprs[0])
	}
}

func (validation *Validation) validateExpr(objects map[string]Object, expr *read.Expr) {
	if expr.VarName != nil {
		// Check if the RHS variable exists
		_, found := objects[expr.VarName.Value]
		if !found {
			validation.addError(newUnknownObject(expr.VarName.Value, expr.VarName.Location))
		}
	}

	if expr.Call != nil {
		validation.validateCall(objects, expr.Call)
	}
}

func (validation *Validation) validateCall(objects map[string]Object, call *read.Call) {
	// Check if the  procedure exists
	_, found := objects[call.Primary.Value]
	if !found {
		validation.addError(newUnknownObject(call.Primary.Value, call.Primary.Location))
	}

	for idx := range call.Exprs {
		subexpr := &call.Exprs[idx]
		validation.validateExpr(objects, subexpr)
	}
}

func (validation *Validation) HasProgram(src *read.Source) {
	if src.Program == nil {
		validation.addError(ProgramingMissing{})
	}
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
			validation.addError(NoExprToInferVariableType{VarName: stmt.Var.Name.Value})
			continue
		}
		exp := stmt.Var.Exprs[0]

		if exp.Null {
			validation.addError(CanNotInferTypeFromNull{VarName: stmt.Var.Name.Value})
			continue
		}

		if exp.VarName != nil {
			validation.addError(CanNotInferTypeFromOtherVariable{VarName: stmt.Var.Name.Value})
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
			if unicode.IsUpper([]rune(exp.Call.Primary.Value)[0]) && strings.TrimSpace(exp.Call.Secondary.Value) == "" {
				stmt.Var.VarType = TypeOf(exp.Call.Primary.Value)
				continue
			}
			validation.addError(CanNotInferTypeFromCall{VarName: stmt.Var.Name.Value})
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

func (validation *Validation) CheckTypesExists(src *read.Source) {

	if src.Program != nil {
		validation.checkProcTypesExists(*src.Program)
	}

	for _, proc := range src.Procs {
		validation.checkProcTypesExists(proc)
	}
}

func (validation *Validation) checkProcTypesExists(proc read.Proc) {
	if proc.Type != nil {
		validation.checkTypeExists(*proc.Type)
	}
	for _, stmt := range proc.Statements {
		validation.checkStatementTypesExists(stmt)
	}
}

func (validation *Validation) checkStatementTypesExists(statement read.Statement) {
	if statement.Var != nil {
		validation.checkVarTypeExists(*statement.Var)
	}
}

func (validation *Validation) checkVarTypeExists(varStmt read.Var) {
	if varStmt.VarType != nil {
		validation.checkTypeExists(*varStmt.VarType)
	}
}

func (validation *Validation) checkTypeExists(typ read.Type) {
	if isUnknownType(typ.Out.Name) {
		validation.errors = append(validation.errors, UnknownType{TypeName: typ.Out.Name})
	}
	for _, in := range typ.In {
		validation.checkTypeExists(*in.Typ)
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
		common.TypeAscii,
		common.TypeSize,
	}

	for _, t := range types {
		if typeName == t {
			return true
		}
	}

	return false
}

func fork(objects map[string]Object) map[string]Object {
	forked := make(map[string]Object)
	for key, value := range objects {
		forked[key] = value
	}
	return forked
}
