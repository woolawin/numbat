package validator

import (
	"fmt"
	"numbat/internal/common"
)

type ValidationError interface {
	Message() string
}

type UnknownType struct {
	TypeName common.Name
}

func NewUnknownType(name common.Name) UnknownType {
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
	VarName common.Name
}

func NewNoExprToInferVariableType(varName common.Name) NoExprToInferVariableType {
	return NoExprToInferVariableType{VarName: varName}
}

func (verr NoExprToInferVariableType) Message() string {
	return fmt.Sprintf("can not infer type for variable `%s` without expression", verr.VarName.Value)
}

type CanNotInferTypeFromNull struct {
	Location common.Location
}

func NewCanNotInferTypeFromNull(location common.Location) CanNotInferTypeFromNull {
	return CanNotInferTypeFromNull{Location: location}
}

func (verr CanNotInferTypeFromNull) Message() string {
	return fmt.Sprintf("(%d,%d) can not infer type from null", verr.Location.Line, verr.Location.Column)
}

type CanNotInferTypeFromOtherVariable struct {
	Location common.Location
}

func NewCanNotInferTypeFromOtherVariable(location common.Location) CanNotInferTypeFromOtherVariable {
	return CanNotInferTypeFromOtherVariable{Location: location}
}

func (verr CanNotInferTypeFromOtherVariable) Message() string {
	return fmt.Sprintf("(%d,%d) not permitted to infer type from other variable", verr.Location.Line, verr.Location.Column)
}

type CanNotInferTypeFromCall struct {
	Location common.Location
}

func NewCanNotInferTypeFromCall(location common.Location) CanNotInferTypeFromCall {
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
	Name           common.Name
	ConflictedWith common.Location
}

func (verr NameConflict) Message() string {
	return fmt.Sprintf("Object name conflicts with name on %d,%d", verr.ConflictedWith.Line, verr.ConflictedWith.Column)
}

func NewNameConflict(name common.Name, conflictedWith common.Location) NameConflict {
	return NameConflict{Name: name, ConflictedWith: conflictedWith}
}

type UnknownObject struct {
	Name     string
	Location common.Location
}

func (verr UnknownObject) Message() string {
	return fmt.Sprintf("(%d,%d) unknown object: %s", verr.Location.Line, verr.Location.Column, verr.Name)
}

func NewUnknownObject(name string, location common.Location) UnknownObject {
	return UnknownObject{Name: name, Location: location}
}

type IncompatibleType struct {
	Location common.Location
	Actual   common.Type
	Expected common.Type
}

func NewIncompatibleType(actual common.Type, expected common.Type, location common.Location) IncompatibleType {
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
	Location common.Location
	Actual   int
	Expected int
}

func NewIncorrectArgumentCount(name string, location common.Location, actual, expected int) IncorrectArgumentCount {
	return IncorrectArgumentCount{ProcName: name, Location: location, Actual: actual, Expected: expected}
}

func (verr IncorrectArgumentCount) Message() string {
	return fmt.Sprintf("procedure expected %d,  %d provided", verr.Expected, verr.Actual)
}
