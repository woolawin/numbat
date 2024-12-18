package validator

import (
	. "numbat/internal/common"
	"numbat/internal/read"
	"strings"
)

type Validation struct {
	errors []ValidationError
}

func NewValidation() Validation {
	return Validation{}
}

func (validation *Validation) Validate(src *read.Source) *Source {

	source := NewSource()

	// First before validating statements, we need to find all the procedures that could be used and first register
	// them to know they are valid references
	for idx := range src.Procs {
		proc := &src.Procs[idx]
		object, found := source.Context.GetObject(proc.Name.Value)
		if found {
			validation.addError(NewNameConflict(proc.Name.ToName(), object.GetName().Location))
		} else {
			t := validation.validateType(proc.Type, &source.Context, false)
			obj := NewProcedureForwardDeclaration(proc.Name.ToName(), t)
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

func (validation *Validation) procedure(proc *read.Proc, procedure *Procedure) {
	if procedure.Type != nil {
		for _, in := range procedure.Type.GetIn() {
			parameter := Parameter{
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

func (validation *Validation) statements(stmts []read.Statement, initialContext *Context) []Statement {
	var statements []Statement
	context := initialContext
	for idx := range stmts {
		stmt, ok := validation.statement(&stmts[idx], NewContext(context))
		if ok {
			context = NewContext(stmt.GetContext())
			statements = append(statements, stmt)
		}
	}
	return statements
}

func (validation *Validation) validateType(t *read.Type, context *Context, reportError bool) Type {
	if t == nil {
		return NewNoType()
	}

	// If the type is `()`, ie no input and out then it is also a no type
	if t.Out.Name == "" && len(t.In) == 0 {
		return NewNoType()
	}

	var outType Type = NewNoType()
	if t.Out.Name != "" {
		// There is an out type specified so we must validate it
		foundType, found := context.GetType(t.Out.Name)
		if !found {
			if reportError {
				validation.errors = append(validation.errors, UnknownType{TypeName: t.Out.ToName()})
			}
			return NewCompileErrorType()
		}
		outType = foundType
	}

	if len(t.In) == 0 {
		return outType
	}

	var ins []InType
	for _, param := range t.In {
		in := validation.validateType(param.Typ, context, reportError)
		if in.IsCompileError() {
			return NewCompileErrorType()
		}
		var defaultValue Expression
		if len(param.Expr) == 1 {
			// There is a default value to this in type so check that it is valid
			e, ok := validation.expression(&param.Expr[0], nil, context, false)
			if ok {
				defaultValue = e
			} else {
				// Assign a null expression just as a placeholder so follow on code knows this parameter is
				// optional
				defaultValue = NewLiteralExpression("", "", NullType{})
			}
		}
		ins = append(ins, NewInType(in, param.Name.ToName(), defaultValue))
	}

	return NewStandardType(outType.GetOut(), ins)
}

func (validation *Validation) statement(stmt *read.Statement, context *Context) (Statement, bool) {
	if stmt.Var != nil {
		return validation.variable(stmt.Var, context)
	}

	if stmt.Call != nil {
		return validation.procedureCall(stmt.Call, context)
	}

	return nil, false // ERROR
}

func (validation *Validation) variable(stmt *read.Var, context *Context) (VariableDeclaration, bool) {

	object, found := context.GetObject(stmt.Name.Value)
	if found {
		// if there is already something with the same name, validate the expression if present and then
		// move on to the next statement
		validation.addError(NewNameConflict(stmt.Name.ToName(), object.GetName().Location))
		if len(stmt.Exprs) == 1 {
			validation.expression(&stmt.Exprs[0], nil, context, stmt.VarType == nil)
		}
		return VariableDeclaration{}, false
	}

	var variableDeclaration VariableDeclaration
	var variableType Type
	var variableValue Expression

	if stmt.VarType != nil {
		// They specified a type so make sure it is a correct type
		variableType = validation.validateType(stmt.VarType, context, true)
	}

	if len(stmt.Exprs) == 0 {
		if variableType == nil {
			// If they did not specify a type, then they must specify an expression to infer the type
			validation.addError(NewNoExprToInferVariableType(stmt.Name.ToName()))
		}
	} else {
		expr, ok := validation.expression(&stmt.Exprs[0], variableType, context, variableType == nil)
		if ok {
			variableValue = expr
			variableType = expr.GetType()
		}
	}

	variableDeclaration = NewVariableDeclaration(context, stmt.Name.ToName(), variableType, variableValue)
	context.AddObject(stmt.Name.Value, &variableDeclaration)

	return variableDeclaration, true
}

func (validation *Validation) expression(expr *read.Expr, expectedType Type, context *Context, forInference bool) (Expression, bool) {

	var expression Expression
	if expr.Type != nil {
		validation.validateType(expr.Type, context, true)
	}

	if expr.VarName != nil {
		if forInference {
			// Not allowed to infer types from other variables
			validation.addError(NewCanNotInferTypeFromOtherVariable(expr.Location))
			return nil, false
		}

		// Check if the RHS variable exists
		object, found := context.GetObject(expr.VarName.Value)
		if !found {
			validation.addError(NewUnknownObject(expr.VarName.Value, expr.VarName.Location))
		} else {
			ve := NewVariableExpression(expr.VarName.ToName(), object.GetType())
			expression = &ve
		}
	}

	if expr.Call != nil {
		if forInference {
			// By design, we only support inferring a type from a procedure call
			validation.addError(NewCanNotInferTypeFromCall(expr.Location))
		}

		//fmt.Println("proc expr")
		//for o := range context.Objects {
		//	fmt.Println("proc obj", o)
		//}
		//fmt.Println("don")

		call, ok := validation.procedureCall(expr.Call, context)
		if ok {
			ce := NewProcedureExpression(call)
			expression = &ce
		} else {
			return expression, false
		}
	}

	if expr.Boolean != nil {
		e := NewLiteralExpression(*expr.Boolean, "", NewBoolType())
		expression = &e
	}

	if expr.Number != nil {
		e := NewLiteralExpression(*expr.Number, "", numericType(*expr.Number))
		expression = &e
	}

	if expr.Str != nil {
		e := NewLiteralExpression(*expr.Str, "", NewStringType())
		expression = &e
	}

	if expr.Null {
		if forInference {
			validation.addError(NewCanNotInferTypeFromNull(expr.Location))
			return nil, false
		}
		e := NewLiteralExpression("", "", NullType{})
		expression = &e
	}

	if expr.Hex != nil {
		e := NewLiteralExpression(*expr.Hex, "", NewByteType())
		expression = &e
	}

	if expression == nil {
		// TODO return error
		return expression, false
	}

	if expectedType == nil {
		return expression, true
	}
	et := expression.GetType()
	if areTypesIncompatible(expectedType, et) {
		validation.addError(NewIncompatibleType(et, expectedType, expr.Location))
	}
	return expression, true
}

func (validation *Validation) procedureCall(call *read.Call, context *Context) (ProcedureCall, bool) {
	// Check if the  procedure exists
	object, found := context.GetObject(call.Primary.Value)
	if !found {
		validation.addError(NewUnknownObject(call.Primary.Value, call.Primary.Location))
		return ProcedureCall{}, false // TODO Still validate arguments where possible
	}

	validateParamAgainstType := false

	if object.GetType() != nil {
		if len(call.Exprs) != len(object.GetType().GetIn()) {
			validation.addError(NewIncorrectArgumentCount(call.Primary.Value, call.Primary.Location, len(call.Exprs), len(call.Exprs)))
		} else {
			validateParamAgainstType = true
		}
	}

	var arguments []Expression
	argumentsOk := true
	for idx := range call.Exprs {
		subexpr := &call.Exprs[idx]
		var typeToCheckAgainst Type
		if validateParamAgainstType {
			typeToCheckAgainst = object.GetType().GetIn()[idx].Type
		}
		e, ok := validation.expression(subexpr, typeToCheckAgainst, context, false)
		if !ok && argumentsOk {
			argumentsOk = false
		}
		arguments = append(arguments, e)
	}
	return NewProcedureCall(context, object, arguments), argumentsOk
}

func numericType(value string) Type {
	if strings.Contains(value, "e+") {
		return NewFloat64Type()
	}

	if strings.Contains(value, ".") {
		return NewFloat64Type()
	}
	return NewInt32Type()

}

func areTypesIncompatible(left, right Type) bool {
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
