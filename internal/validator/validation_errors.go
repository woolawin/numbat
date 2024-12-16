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
	VarName common.Name
}

func NewCanNotInferTypeFromNull(name common.Name) CanNotInferTypeFromNull {
	return CanNotInferTypeFromNull{VarName: name}
}

func (verr CanNotInferTypeFromNull) Message() string {
	return fmt.Sprintf("can not infer type for variable `%s` with null expression", verr.VarName.Value)
}

type CanNotInferTypeFromOtherVariable struct {
	VarName common.Name
}

func NewCanNotInferTypeFromOtherVariable(name common.Name) CanNotInferTypeFromOtherVariable {
	return CanNotInferTypeFromOtherVariable{VarName: name}
}

func (verr CanNotInferTypeFromOtherVariable) Message() string {
	return fmt.Sprintf("infering type for variable `%s` from variable is not supported", verr.VarName.Value)
}

type CanNotInferTypeFromCall struct {
	VarName common.Name
}

func NewCanNotInferTypeFromCall(varName common.Name) CanNotInferTypeFromCall {
	return CanNotInferTypeFromCall{VarName: varName}
}

func (verr CanNotInferTypeFromCall) Message() string {
	return fmt.Sprintf("infering type for variable `%s` from procedure call is not supported", verr.VarName.Value)
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

func newNameConflict(name common.Name, conflictedWith common.Location) NameConflict {
	return NameConflict{Name: name, ConflictedWith: conflictedWith}
}

type UnknownObject struct {
	Name     string
	Location common.Location
}

func (verr UnknownObject) Message() string {
	return fmt.Sprintf("unknown object: %s", verr.Name)
}

func newUnknownObject(name string, location common.Location) UnknownObject {
	return UnknownObject{Name: name, Location: location}
}

type IncompatibleType struct {
	Location common.Location
	Actual   string
	Expected string
}

func newIncompatibleType(actual string, expected string, location common.Location) IncompatibleType {
	return IncompatibleType{Location: location, Actual: actual, Expected: expected}
}

func (verr IncompatibleType) Message() string {
	return fmt.Sprintf("can not use type `%s` for required type `%s`", verr.Actual, verr.Expected)
}

type IncorrectArgumentCount struct {
	ProcName string
	Location common.Location
	Actual   int
	Expected int
}

func newIncorrectArgumentCount(name string, location common.Location, actual, expected int) IncorrectArgumentCount {
	return IncorrectArgumentCount{ProcName: name, Location: location, Actual: actual, Expected: expected}
}

func (verr IncorrectArgumentCount) Message() string {
	return fmt.Sprintf("procedure expected %d,  %d provided", verr.Expected, verr.Actual)
}
