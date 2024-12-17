package validator

import (
	"numbat/internal/common"
	"numbat/internal/read"
	"strings"
	"unicode"
)

type Validation struct {
	errors []ValidationError
}

func NewValidation() Validation {
	return Validation{}
}

func (validation *Validation) Validate(src *read.Source) *common.Source {

	source := common.NewSource()

	// First before validating statements, we need to find all the procedures that could be used and first register
	// them to know they are valid references
	for idx := range src.Procs {
		proc := &src.Procs[idx]
		object, found := source.Context.GetObject(proc.Name.Value)
		if found {
			validation.addError(NewNameConflict(proc.Name.ToName(), object.GetName().Location))
		} else {
			t := validation.validateType(proc.Type, &source.Context, false)
			obj := common.NewProcedureForwardDeclaration(proc.Name.ToName(), t)
			source.Context.AddObject(proc.Name.Value, &obj)
		}
	}

	if src.Program == nil {
		validation.addError(&ProgramingMissing{})
	} else {
		source.Program.AddStatements(validation.statements(src.Program.Statements, &source.Context))
	}

	for idx := range src.Procs {
		proc := &src.Procs[idx]
		source.Context.GetObject(proc.Name.Value)
		t := validation.validateType(proc.Type, &source.Context, true)
		procedure := source.AddProcedure(proc.Name.ToName(), t)
		validation.procedure(proc, procedure)
	}

	return source
}

func (validation *Validation) procedure(proc *read.Proc, procedure *common.Procedure) {

	if procedure.Type != nil {
		for _, in := range procedure.Type.GetIn() {
			parameter := common.Parameter{
				Name: in.Name,
				Type: in.Type,
			}
			object, found := procedure.Context.GetObject(in.Name.Value)
			if found {
				validation.addError(NewNameConflict(in.Name, object.GetName().Location))
			} else {
				procedure.Context.AddObject(in.Name.Value, parameter)
			}
		}
	}

	procedure.AddStatements(validation.statements(proc.Statements, procedure.Context))
}

func (validation *Validation) statements(stmts []read.Statement, context *common.Context) []common.Statement {
	var statements []common.Statement
	for idx := range stmts {
		stmt := validation.statement(&stmts[idx], context)
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}
	return statements
}

func (validation *Validation) validateType(t *read.Type, context *common.Context, reportError bool) common.Type {
	if t == nil {
		return common.NewNoType()
	}

	// If the type is `()`, ie no input and out then it is also a no type
	if t.Out.Name == "" && len(t.In) == 0 {
		return common.NewNoType()
	}

	var outType common.Type = common.NewNoType()
	if t.Out.Name != "" {
		// There is an out type specified so we must validate it
		foundType, found := context.GetType(t.Out.Name)
		if !found {
			if reportError {
				validation.errors = append(validation.errors, UnknownType{TypeName: t.Out.ToName()})
			}
			return common.NewCompileErrorType()
		}
		outType = foundType
	}

	if len(t.In) == 0 {
		return outType
	}

	var ins []common.InType
	for _, param := range t.In {
		in := validation.validateType(param.Typ, context, reportError)
		if in.IsCompileError() {
			return common.NewCompileErrorType()
		}
		var defaultValue *common.Expression
		if len(param.Expr) == 1 {
			// There is a default value to this in type so check that it is valid
			e := validation.expression(&param.Expr[0], nil, context)
			defaultValue = &e
		}
		ins = append(ins, common.NewInType(in, param.Name.ToName(), defaultValue))
	}

	return common.NewStandardType(outType.GetOut(), ins)
}

func (validation *Validation) statement(stmt *read.Statement, context *common.Context) common.Statement {
	if stmt.Var != nil {
		return validation.variable(stmt.Var, context)
	}

	if stmt.Call != nil {
		return validation.procedureCall(stmt.Call, context)
	}

	return nil // ERROR
}

func (validation *Validation) variable(stmt *read.Var, context *common.Context) *common.VariableDeclaration {
	var variableDeclaration *common.VariableDeclaration
	var variableType common.Type

	checkExprType := true

	if stmt.VarType != nil {
		// They specified a type so make sure it is a correct type
		variableType = validation.validateType(stmt.VarType, context, true)
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
			variableType = exprType
		}
	}

	// Check if variable name is already taken
	object, found := context.GetObject(stmt.Name.Value)
	if found {
		validation.addError(NewNameConflict(stmt.Name.ToName(), object.GetName().Location))
	} else {
		if variableType == nil {
			// if we do not know the variable type, we can not work with it so skip, we should
			// add a placeholder object to say it exists but type is not known @TODO
			return nil
		}
		a := common.NewVariableDeclaration(stmt.Name.ToName(), variableType, nil)
		variableDeclaration = &a
		context.AddObject(stmt.Name.Value, variableDeclaration)
	}

	if len(stmt.Exprs) == 1 && checkExprType {
		e := validation.expression(&stmt.Exprs[0], variableType, context)
		variableDeclaration.Value = e
	}

	return variableDeclaration
}

// @TODO merge this with inferVarType
func (validation *Validation) expression(expr *read.Expr, expectedType common.Type, context *common.Context) common.Expression {

	var expression common.Expression
	if expr.Type != nil {
		validation.validateType(expr.Type, context, true)
	}

	if expr.VarName != nil {
		// Check if the RHS variable exists
		object, found := context.GetObject(expr.VarName.Value)
		if !found {
			validation.addError(NewUnknownObject(expr.VarName.Value, expr.VarName.Location))
		} else {
			ve := common.NewVariableExpression(expr.VarName.ToName(), object.GetType())
			expression = &ve
		}
	}

	if expr.Call != nil {
		call := validation.procedureCall(expr.Call, context)
		if call != nil {
			ce := common.NewProcedureExpression(*call)
			expression = &ce
		}

	}

	if expr.Boolean != nil {
		e := common.NewLiteralExpression(*expr.Boolean, "", common.BoolType{})
		expression = &e
	}

	if expr.Number != nil {
		e := common.NewLiteralExpression(*expr.Number, "", numericType(*expr.Number))
		expression = &e
	}

	if expr.Str != nil {
		e := common.NewLiteralExpression(*expr.Str, "", common.StringType{})
		expression = &e
	}

	if expr.Null {
		e := common.NewLiteralExpression("", "", common.NullType{})
		expression = &e
	}

	if expr.Hex != nil {
		e := common.NewLiteralExpression(*expr.Hex, "", common.ByteType{})
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
	if areTypesIncompatible(expectedType, et) {
		validation.addError(NewIncompatibleType(et, expectedType, expr.Location))
	}

	return expression
}

func (validation *Validation) inferVarType(expr *read.Expr, varName common.Name, context *common.Context) (common.Type, bool) {

	if expr.Null {
		validation.addError(CanNotInferTypeFromNull{VarName: varName})
		return common.NewCompileErrorType(), true
	}

	if expr.VarName != nil {
		validation.addError(CanNotInferTypeFromOtherVariable{VarName: varName})
		return common.NewCompileErrorType(), true
	}

	if expr.Call != nil {
		if unicode.IsUpper([]rune(expr.Call.Primary.Value)[0]) && strings.TrimSpace(expr.Call.Secondary.Value) == "" {
			return validation.validateType(TypeOf(expr.Call.Primary.Value), context, true), false
		}
		validation.addError(CanNotInferTypeFromCall{VarName: varName})
		return common.NewCompileErrorType(), true
	}

	if expr.Boolean != nil {
		return common.BoolType{}, false
	}

	if expr.Number != nil {
		return numericType(*expr.Number), false
	}

	if expr.Str != nil {
		return common.StringType{}, false
	}

	if expr.Null {
		return common.NullType{}, false
	}

	if expr.Hex != nil {
		return common.ByteType{}, false
	}

	// Should never get here
	return common.NewCompileErrorType(), false
}

func (validation *Validation) procedureCall(call *read.Call, context *common.Context) *common.ProcedureCall {
	// Check if the  procedure exists

	object, found := context.GetObject(call.Primary.Value)
	if !found {
		validation.addError(NewUnknownObject(call.Primary.Value, call.Primary.Location))
		return nil // Still validate arguments where possible
	}

	validateParamAgainstType := false

	if object.GetType() != nil {
		if len(call.Exprs) != len(object.GetType().GetIn()) {
			validation.addError(NewIncorrectArgumentCount(call.Primary.Value, call.Primary.Location, len(call.Exprs), len(call.Exprs)))
		} else {
			validateParamAgainstType = true
		}
	}

	var arguments []common.Expression
	for idx := range call.Exprs {
		subexpr := &call.Exprs[idx]
		var typeToCheckAgainst common.Type
		if validateParamAgainstType {
			typeToCheckAgainst = object.GetType().GetIn()[idx].Type
		}
		e := validation.expression(subexpr, typeToCheckAgainst, context)
		if e != nil {
			arguments = append(arguments, e)
		}
	}
	c := common.NewProcedureCall(object, arguments)
	return &c
}

func numericType(value string) common.Type {
	if strings.Contains(value, "e+") {
		return common.Float64Type{}
	}

	if strings.Contains(value, ".") {
		return common.Float64Type{}
	}
	return common.Int32Type{}

}

func TypeOf(name string) *read.Type {
	return &read.Type{Out: read.TypeOut{Name: name}}
}

func areTypesIncompatible(left, right common.Type) bool {
	if left.IsNeverCompatible() || right.IsNeverCompatible() {
		return true
	}
	if left.GetOut().Value != right.GetOut().Value {
		return true
	}
	if len(left.GetIn()) != len(right.GetIn()) {
		return true
	}
	for idx := range left.GetIn() {
		if areTypesIncompatible(left.GetIn()[idx].Type, right.GetIn()[idx].Type) {
			return true
		}
	}
	return false
}

func (validation *Validation) addError(verr ValidationError) {
	validation.errors = append(validation.errors, verr)
}
