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
	return UnknownType{TypeName: name}
}

func (verr UnknownType) Message() string {
	return fmt.Sprintf(
		"(%d,%d) unknown type `%s`",
		verr.TypeName.Location.Line,
		verr.TypeName.Location.Column,
		verr.TypeName.Value,
	)
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
	Actual   InOutType
	Expected InOutType
}

func NewIncompatibleType(actual InOutType, expected InOutType, location Location) IncompatibleType {
	return IncompatibleType{Location: location, Actual: actual, Expected: expected}
}

func (verr IncompatibleType) Message() string {
	return fmt.Sprintf(
		"(%d,%d) can not use type `%s` for `%s`",
		verr.Location.Line,
		verr.Location.Column,
		verr.Actual.String(),
		verr.Expected.String(),
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

type DoesNotReturnValue struct {
	ObjectName string
	Location   Location
}

func NewDoesNotReturnValue(objectName string, location Location) DoesNotReturnValue {
	return DoesNotReturnValue{ObjectName: objectName, Location: location}
}

func (verr DoesNotReturnValue) Message() string {
	return fmt.Sprintf(
		"(%d,%d) `%s` does not return a value",
		verr.Location.Line,
		verr.Location.Column,
		verr.ObjectName,
	)
}

type ReturnValueRequired struct {
	ObjectName       string
	Location         Location
	ExpectedTypeName string
}

func NewReturnValueRequired(objectName string, location Location, expectedTypeName string) ReturnValueRequired {
	return ReturnValueRequired{ObjectName: objectName, Location: location, ExpectedTypeName: expectedTypeName}
}

func (verr ReturnValueRequired) Message() string {
	return fmt.Sprintf(
		"(%d,%d) `%s` must return value of type `%s`",
		verr.Location.Line,
		verr.Location.Column,
		verr.ObjectName,
		verr.ExpectedTypeName,
	)
}

type MissingReturnStatement struct {
	ObjectName string
	Location   Location
}

func NewMissingReturnStatement(objectName string, location Location) MissingReturnStatement {
	return MissingReturnStatement{ObjectName: objectName, Location: location}
}

func (verr MissingReturnStatement) Message() string {
	return fmt.Sprintf(
		"(%d,%d) `%s` does not return a value on all paths",
		verr.Location.Line,
		verr.Location.Column,
		verr.ObjectName,
	)
}

type IncompatibleReturnValueType struct {
	Location Location
	Actual   InOutType
	Expected InOutType
}

func NewIncompatibleReturnValueType(location Location, expected, actual InOutType) ValidationError {
	return IncompatibleReturnValueType{Location: location, Actual: actual, Expected: expected}
}

func (verr IncompatibleReturnValueType) Message() string {
	return fmt.Sprintf(
		"(%d,%d) can not use return value of type `%s` for `%s`",
		verr.Location.Line,
		verr.Location.Column,
		verr.Actual.String(),
		verr.Expected.String(),
	)
}

type InvalidSequenceSize struct {
	Location Location
	Value    string
}

func NewInvalidSequenceSize(location Location, value string) InvalidSequenceSize {
	return InvalidSequenceSize{Location: location, Value: value}
}

func (verr InvalidSequenceSize) Message() string {
	return fmt.Sprintf(
		"(%d,%d) invalid sequence size `%s`",
		verr.Location.Line,
		verr.Location.Column,
		verr.Value,
	)
}

type IncorrectSequenceSize struct {
	Location Location
	Expected int
	Actual   int
}

func NewIncorrectSequenceSize(location Location, expected int, actual int) IncorrectSequenceSize {
	return IncorrectSequenceSize{Location: location, Expected: expected, Actual: actual}
}

func (verr IncorrectSequenceSize) Message() string {
	return fmt.Sprintf(
		"(%d,%d) expected sequence size of `%d` but found `%d`",
		verr.Location.Line,
		verr.Location.Column,
		verr.Expected,
		verr.Actual,
	)
}

type IncompatibleElementType struct {
	Location Location
	Expected InOutType
	Actual   InOutType
}

func NewIncompatibleElementType(location Location, actual, expected InOutType) ValidationError {
	return IncompatibleElementType{Location: location, Actual: actual, Expected: expected}
}

func (verr IncompatibleElementType) Message() string {
	return fmt.Sprintf(
		"(%d,%d) value with type `%s` can not be used for sequence of `%s`",
		verr.Location.Line,
		verr.Location.Column,
		verr.Actual.String(),
		verr.Expected.String(),
	)
}

type IncompatibleParameterType struct {
	Location Location
	Expected InOutType
	Actual   InOutType
}

func NewIncompatibleParameterType(location Location, expected, actual InOutType) ValidationError {
	return IncompatibleParameterType{Location: location, Actual: actual, Expected: expected}
}

func (verr IncompatibleParameterType) Message() string {
	return fmt.Sprintf(
		"(%d, %d) value with type `%s` can not be used for paramter of `%s`",
		verr.Location.Line,
		verr.Location.Column,
		verr.Actual.String(),
		verr.Expected.String(),
	)
}
