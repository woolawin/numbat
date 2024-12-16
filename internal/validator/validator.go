package validator

import (
	"fmt"
	"numbat/internal/common"
	"numbat/internal/read"
	"strings"
	"unicode"
)

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

type ValidationError interface {
	Message() string
}

type Validation struct {
	errors []ValidationError
}

func NewValidation() Validation {
	return Validation{}
}

type Object struct {
	Name     read.Name
	Location common.Location
	Type     *read.Type
}

func (validation *Validation) Validate(src *read.Source) *common.Project {

	project := common.NewProject()

	for idx := range src.Procs {
		proc := &src.Procs[idx]
		object, found := project.Context.GetObject(proc.Name.Value)
		if found {
			validation.addError(newNameConflict(proc.Name.ToName(), object.GetName().Location))
		} else {
			t := validation.validateType(proc.Type, &project.Context, false)
			obj := common.NewProcedureDefinition(proc.Name.ToName(), &t)
			project.Context.AddObject(proc.Name.Value, &obj)
		}
	}

	if src.Program == nil {
		validation.addError(&ProgramingMissing{})
	} else {
		project.Program.AddStatements(validation.validateStatements(src.Program.Statements, &project.Context))
	}

	for idx := range src.Procs {
		proc := &src.Procs[idx]
		project.Context.GetObject(proc.Name.Value)
		t := validation.validateType(proc.Type, &project.Context, true)
		procedure := project.AddProcedure(proc.Name.ToName(), &t)
		validation.validateProc(proc, procedure)
	}

	return project
}

func (validation *Validation) validateProc(proc *read.Proc, procedure *common.Procedure) {

	if procedure.Type != nil {
		for _, in := range procedure.Type.In {
			parameter := common.Parameter{
				Name: in.Name,
				Type: in.Type,
			}
			object, found := procedure.Context.GetObject(in.Name.Value)
			if found {
				validation.addError(newNameConflict(in.Name, object.GetName().Location))
			} else {
				procedure.Context.AddObject(in.Name.Value, &parameter)
			}
		}
	}

	procedure.AddStatements(validation.validateStatements(proc.Statements, procedure.Context))
}

func (validation *Validation) validateStatements(stmts []read.Statement, context *common.Context) []common.Statement {
	var statements []common.Statement
	for idx := range stmts {
		stmt := validation.validateStatement(&stmts[idx], context)
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}
	return statements
}

func (validation *Validation) validateType(t *read.Type, context *common.Context, reportError bool) common.Type {
	if t == nil {
		return common.Type{}
	}
	out := common.Type{}
	if t.Out.Name != "" {
		outType, found := context.GetType(t.Out.Name)
		if !found {
			if reportError {
				validation.errors = append(validation.errors, UnknownType{TypeName: t.Out.ToName()})
			}
			return common.Type{}
		} else {
			out = outType
		}
	}

	var ins []common.InType
	for _, param := range t.In {
		in := validation.validateType(param.Typ, context, reportError)
		if in.Out.Value == "" { // TODO replace isCompilationErrorType
			return common.Type{}
		}
		var defaultValue *common.Expression
		if len(param.Expr) == 1 {
			e := validation.validateExpr(&param.Expr[0], nil, context)
			defaultValue = &e
		}
		ins = append(ins, common.NewInType(in, param.Name.ToName(), defaultValue))
	}
	return common.NewType(out.Out, ins)
}

func (validation *Validation) validateStatement(stmt *read.Statement, context *common.Context) common.Statement {
	if stmt.Var != nil {
		return validation.validateVar(stmt.Var, context)
	}

	if stmt.Call != nil {
		return validation.validateCall(stmt.Call, context)
	}

	return nil // ERROR
}

func (validation *Validation) validateVar(stmt *read.Var, context *common.Context) *common.VariableDeclaration {
	var variableDeclaration *common.VariableDeclaration
	var variableType *common.Type

	checkExprType := true

	if stmt.VarType != nil {
		// The specified a type so make sure it is a correct type
		t := validation.validateType(stmt.VarType, context, true)
		variableType = &t
	} else {
		// No type specified so lets try to infer it by the expression
		if len(stmt.Exprs) == 0 {
			// No expression to infer the type so error
			validation.addError(NewNoExprToInferVariableType(stmt.Name.ToName()))
		} else {
			exprType, errored := validation.inferVarType(&stmt.Exprs[0], stmt.Name.ToName(), context)
			// if we can not validate the expression type, do not try to validate the expression in general
			// otherwise there will just be double error reporting
			checkExprType = !errored
			variableType = &exprType
		}
	}

	// Check if variable name is already taken
	object, found := context.GetObject(stmt.Name.Value)
	if found {
		validation.addError(newNameConflict(stmt.Name.ToName(), object.GetName().Location))
	} else {
		if variableType == nil {
			// if we do not know the variable type, we can not work with it so skip, we should
			// add a placeholder object to say it exists but type is not known @TODO
			return nil
		}
		a := common.NewVariableDeclaration(stmt.Name.ToName(), *variableType, nil)
		variableDeclaration = &a
		context.AddObject(stmt.Name.Value, variableDeclaration)
	}

	if len(stmt.Exprs) == 1 && checkExprType {
		e := validation.validateExpr(&stmt.Exprs[0], variableType, context)
		variableDeclaration.Value = e
	}

	return variableDeclaration
}

func (validation *Validation) validateExpr(expr *read.Expr, expectedType *common.Type, context *common.Context) common.Expression {

	var expression common.Expression
	if expr.Type != nil {
		validation.validateType(expr.Type, context, true)
	}

	if expr.VarName != nil {
		// Check if the RHS variable exists
		object, found := context.GetObject(expr.VarName.Value)
		if !found {
			validation.addError(newUnknownObject(expr.VarName.Value, expr.VarName.Location))
		} else {
			ve := common.NewVariableExpression(expr.VarName.ToName(), *object.GetType())
			expression = &ve
		}
	}

	if expr.Call != nil {
		call := validation.validateCall(expr.Call, context)
		if call != nil {
			ce := common.NewProcedureExpression(*call)
			expression = &ce
		}

	}

	if expr.Boolean != nil {
		e := common.NewLiteralExpression(*expr.Boolean, "", common.NewType(common.Name{Value: common.TypeBool}, nil))
		expression = &e
	}

	if expr.Number != nil {
		e := common.NewLiteralExpression(*expr.Number, "", numericType(*expr.Number))
		expression = &e
	}

	if expr.Str != nil {
		e := common.NewLiteralExpression(*expr.Str, "", common.NewType(common.Name{Value: common.LiteralTypeString}, nil))
		expression = &e
	}

	if expr.Null {
		e := common.NewLiteralExpression("", "", common.NewType(common.Name{Value: common.LiteralTypeNull}, nil))
		expression = &e
	}

	if expr.Hex != nil {
		e := common.NewLiteralExpression(*expr.Hex, "", common.NewType(common.Name{Value: common.TypeByte}, nil))
		expression = &e
	}

	if expression == nil {
		// TODO return error
		return expression
	}

	if expectedType == nil {
		return expression
	}
	et := expression.GetType()
	if areTypesIncompatible(expectedType, &et) {
		validation.addError(newIncompatibleType(et.Out.Value, expectedType.Out.Value, expr.Location))
	}

	return expression
}

func (validation *Validation) inferVarType(expr *read.Expr, varName common.Name, context *common.Context) (common.Type, bool) {

	if expr.Null {
		validation.addError(CanNotInferTypeFromNull{VarName: varName})
		return common.Type{}, true
	}

	if expr.VarName != nil {
		validation.addError(CanNotInferTypeFromOtherVariable{VarName: varName})
		return common.Type{}, true
	}

	if expr.Call != nil {
		if unicode.IsUpper([]rune(expr.Call.Primary.Value)[0]) && strings.TrimSpace(expr.Call.Secondary.Value) == "" {
			return validation.validateType(TypeOf(expr.Call.Primary.Value), context, true), false
		}
		validation.addError(CanNotInferTypeFromCall{VarName: varName})
		return common.Type{}, true
	}

	if expr.Boolean != nil {
		return common.NewType(common.Name{Value: common.TypeBool}, nil), false
	}

	if expr.Number != nil {
		return numericType(*expr.Number), false
	}

	if expr.Str != nil {
		return common.NewType(common.Name{Value: common.LiteralTypeString}, nil), false
	}

	if expr.Null {
		return common.NewType(common.Name{Value: common.LiteralTypeNull}, nil), false
	}

	if expr.Hex != nil {
		return common.NewType(common.Name{Value: common.TypeByte}, nil), false
	}

	// Should never get here
	return common.Type{}, false
}

func (validation *Validation) validateCall(call *read.Call, context *common.Context) *common.ProcedureCall {
	// Check if the  procedure exists

	object, found := context.GetObject(call.Primary.Value)
	if !found {
		validation.addError(newUnknownObject(call.Primary.Value, call.Primary.Location))
		return nil // Still validate arguments where possible
	}

	validateParamAgainstType := false

	if object.GetType() != nil {
		if len(call.Exprs) != len(object.GetType().In) {
			validation.addError(newIncorrectArgumentCount(call.Primary.Value, call.Primary.Location, len(call.Exprs), len(call.Exprs)))
		} else {
			validateParamAgainstType = true
		}
	}

	var arguments []common.Expression
	for idx := range call.Exprs {
		subexpr := &call.Exprs[idx]
		var typeToCheckAgainst *common.Type
		if validateParamAgainstType {
			typeToCheckAgainst = &object.GetType().In[idx].Type
		}
		e := validation.validateExpr(subexpr, typeToCheckAgainst, context)
		if e != nil {
			arguments = append(arguments, e)
		}
	}
	c := common.NewProcedureCall(object, arguments)
	return &c
}

func numericType(value string) common.Type {
	if strings.Contains(value, "e+") {
		return common.Type{Out: common.Name{Value: common.TypeFloat64}}
	}

	if strings.Contains(value, ".") {
		return common.Type{Out: common.Name{Value: common.TypeFloat64}}
	}
	return common.Type{Out: common.Name{Value: common.TypeInt32}}

}

func TypeOf(name string) *read.Type {
	return &read.Type{Out: read.TypeOut{Name: name}}
}

func areTypesIncompatible(left, right *common.Type) bool {
	if left == nil || right == nil {
		return true
	}
	if left.Out.Value != right.Out.Value {
		return true
	}
	if len(left.In) != len(right.In) {
		return true
	}
	for idx := range left.In {
		if areTypesIncompatible(&left.In[idx].Type, &right.In[idx].Type) {
			return true
		}
	}
	return false
}

func (validation *Validation) addError(verr ValidationError) {
	validation.errors = append(validation.errors, verr)
}
