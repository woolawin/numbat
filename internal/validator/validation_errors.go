package validator

import (
	"fmt"
	. "numbat/internal/common"
)

type ValidationError interface {
	Message() string
}

type MissingProgram struct {
}

func NewMissingProgram() MissingProgram {
	return MissingProgram{}
}

func (verr MissingProgram) Message() string {
	return "missing program"
}

type ProgramRedefined struct {
	Location Location
}

func NewProgramRedefined(location Location) ProgramRedefined {
	return ProgramRedefined{location}
}

func (verr ProgramRedefined) Message() string {
	return fmt.Sprintf("(%d,%d) program redefined", verr.Location.Line, verr.Location.Column)
}

type UnknownType struct {
	TypeName Name
}

func NewUnknownType(name Name) UnknownType {
	return UnknownType{name}
}

func (verr UnknownType) Message() string {
	return fmt.Sprintf("unknown type `%s`", verr.TypeName.Value)
}

type UnInferrableVar struct {
	VarName string
}

func (verr UnInferrableVar) Message() string {
	return fmt.Sprintf("can not infer type for variable: %s", verr.VarName)
}

type NoExprToInferVariableType struct {
	VarName Name
}

func NewNoExprToInferVariableType(varName Name) NoExprToInferVariableType {
	return NoExprToInferVariableType{VarName: varName}
}

func (verr NoExprToInferVariableType) Message() string {
	return fmt.Sprintf("can not infer type for variable `%s` without expression", verr.VarName.Value)
}

type CanNotInferTypeFromNull struct {
	Location Location
}

func NewCanNotInferTypeFromNull(location Location) CanNotInferTypeFromNull {
	return CanNotInferTypeFromNull{Location: location}
}

func (verr CanNotInferTypeFromNull) Message() string {
	return fmt.Sprintf("(%d,%d) can not infer type from null", verr.Location.Line, verr.Location.Column)
}

type CanNotInferTypeFromOtherVariable struct {
	Location Location
}

func NewCanNotInferTypeFromOtherVariable(location Location) CanNotInferTypeFromOtherVariable {
	return CanNotInferTypeFromOtherVariable{Location: location}
}

func (verr CanNotInferTypeFromOtherVariable) Message() string {
	return fmt.Sprintf("(%d,%d) not permitted to infer type from other variable", verr.Location.Line, verr.Location.Column)
}

type CanNotInferTypeFromCall struct {
	Location Location
}

func NewCanNotInferTypeFromCall(location Location) CanNotInferTypeFromCall {
	return CanNotInferTypeFromCall{Location: location}
}

func (verr CanNotInferTypeFromCall) Message() string {
	return fmt.Sprintf("(%d,%d) not permitted to infer type from procedure call", verr.Location.Line, verr.Location.Column)
}

type ProgramingMissing struct {
}

func (verr ProgramingMissing) Message() string {
	return "No program object found"
}

type NameConflict struct {
	Name           Name
	ConflictedWith Location
}

func (verr NameConflict) Message() string {
	return fmt.Sprintf("Object name conflicts with name on %d,%d", verr.ConflictedWith.Line, verr.ConflictedWith.Column)
}

func NewNameConflict(name Name, conflictedWith Location) NameConflict {
	return NameConflict{Name: name, ConflictedWith: conflictedWith}
}

type UnknownObject struct {
	Name     string
	Location Location
}

func (verr UnknownObject) Message() string {
	return fmt.Sprintf("(%d,%d) unknown object: %s", verr.Location.Line, verr.Location.Column, verr.Name)
}

func NewUnknownObject(name string, location Location) UnknownObject {
	return UnknownObject{Name: name, Location: location}
}

type IncompatibleType struct {
	Location Location
	Actual   Type
	Expected Type
}

func NewIncompatibleType(actual Type, expected Type, location Location) IncompatibleType {
	return IncompatibleType{Location: location, Actual: actual, Expected: expected}
}

func (verr IncompatibleType) Message() string {
	return fmt.Sprintf(
		"(%d,%d) can not use type `%s` for `%s`",
		verr.Location.Line,
		verr.Location.Column,
		verr.Actual.GetOut().Value,
		verr.Expected.GetOut().Value,
	)
}

type IncorrectArgumentCount struct {
	ProcName string
	Location Location
	Actual   int
	Expected int
}

func NewIncorrectArgumentCount(name string, location Location, actual, expected int) IncorrectArgumentCount {
	return IncorrectArgumentCount{ProcName: name, Location: location, Actual: actual, Expected: expected}
}

func (verr IncorrectArgumentCount) Message() string {
	return fmt.Sprintf(
		"(%d,%d) incorrect number of arguments to `%s`, %d expected, %d provided",
		verr.Location.Line,
		verr.Location.Column,
		verr.ProcName,
		verr.Expected,
		verr.Actual,
	)
}
