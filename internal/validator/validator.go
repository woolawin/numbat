package validator

import (
	. "numbat/internal/common"
	"numbat/internal/read"
	"strconv"
	"strings"
)

type Validation struct {
	errors []ValidationError
}

func (validation *Validation) GetErrors() []ValidationError {
	return validation.errors
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

	if len(src.Programs) == 0 {
		validation.addError(NewMissingProgram())
	} else {
		if len(src.Programs) > 1 {
			for i := 1; i < len(src.Programs); i++ {
				validation.addError(NewProgramRedefined(src.Programs[i].Location))
			}
		}
		for _, prog := range src.Programs {
			source.Program.AddStatements(validation.statements(prog.Statements, NewProgramContext(&source.Context)))
		}
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
	for _, in := range procedure.Type.In {
		paramType := NoInOutType()
		if in.Type != nil {
			paramType = *in.Type
		}
		parameter := Parameter{
			Name: in.Name,
			Type: paramType,
		}
		object, found := procedure.Context.GetObject(in.Name.Value)
		if found {
			validation.addError(NewNameConflict(in.Name, object.GetName().Location))
		} else {
			procedure.AddParameter(parameter)
			procedure.Context.AddObject(in.Name.Value, parameter)
		}
	}

	procedure.AddStatements(validation.statements(proc.Statements, procedure.Context))

	if !procedure.Type.Out.IsNoType() {
		// The procedure has a return type so we need to check that there is a value returned
		missingReturn := true
		for _, stmt := range procedure.Statements {
			if stmt.GuaranteedReturn() {
				missingReturn = false
				break
			}
		}
		if missingReturn {
			validation.addError(NewMissingReturnStatement(procedure.Name.Value, proc.Name.Location))
		}
	}
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

func (validation *Validation) statement(stmt *read.Statement, context *Context) (Statement, bool) {
	if stmt.Var != nil {
		return validation.variable(stmt.Var, context)
	}

	if stmt.Call != nil {
		return validation.procedureCall(stmt.Call, context)
	}

	if stmt.Ret != nil {
		return validation.returnStatement(stmt.Ret, stmt.Location, context), true
	}

	return nil, false // ERROR
}

func (validation *Validation) validateType(t *read.Type, context *Context, reportError bool) InOutType {
	if t == nil {
		return NoInOutType()
	}

	// If the type is `()`, ie no input and out then it is also a no type
	if t.Out.Name == "" && len(t.In) == 0 {
		return NoInOutType()
	}

	out := NewSuperAtomicType(NewNoType())

	if t.Out.Name != "" {
		// There is an out type specified so we must validate it
		foundType, found := context.GetType(t.Out.Name)
		if !found {
			if reportError {
				validation.errors = append(validation.errors, NewUnknownType(t.Out.ToName()))
			}
			return NewInOutType(nil, NewSuperAtomicType(NewCompileErrorType()))
		}
		out.Type = foundType
	}

	if t.Out.Sequence {
		out.IsSequence = true
		if t.Out.SequenceSize != "" {
			size, err := strconv.Atoi(t.Out.SequenceSize)
			if err != nil || size < 1 {
				validation.addError(NewInvalidSequenceSize(t.Out.SequenceLocation, t.Out.SequenceSize))
			} else {
				out.SequenceSize = size
			}
		}
	}

	if len(t.In) == 0 {
		return NewInOutType(nil, out)
	}

	var ins []InType
	for _, param := range t.In {
		in := validation.validateType(param.Typ, context, reportError)
		var defaultValue Expression
		if param.Value.IsNotEmpty() {
			// There is a default value to this in type so check that it is valid
			opts := NewExprOpts(ExprOpts{})
			e, ok := validation.expression(param.Value.First(), opts, context, false)
			if ok {
				defaultValue = e
			} else {
				// Assign a null expression just as a placeholder so follow on code knows this parameter is
				// optional
				defaultValue = NewLiteralExpression("", "", NewAtomicInOutType(NullType{}))
			}
		}
		ins = append(ins, NewInType(in, param.Name.ToName(), defaultValue))
	}

	return NewInOutType(ins, out)
}

func (validation *Validation) variable(stmt *read.Var, context *Context) (VariableDeclaration, bool) {

	object, found := context.GetObject(stmt.Name.Value)
	if found {
		// if there is already something with the same name, validate the expression if present and then
		// move on to the next statement
		validation.addError(NewNameConflict(stmt.Name.ToName(), object.GetName().Location))
		if stmt.Value.IsNotEmpty() {
			opts := NewExprOpts(ExprOpts{})
			validation.expression(stmt.Value.First(), opts, context, stmt.VarType == nil)
		}
		return VariableDeclaration{}, false
	}

	var variableDeclaration VariableDeclaration
	var variableType *InOutType
	var variableValue Expression

	if stmt.VarType != nil {
		// They specified a type so make sure it is a correct type
		t := validation.validateType(stmt.VarType, context, true)
		variableType = &t
	}

	if stmt.Value.IsEmpty() {
		if variableType == nil {
			// If they did not specify a type, then they must specify an expression to infer the type
			validation.addError(NewNoExprToInferVariableType(stmt.Name.ToName()))
		}
	} else {
		opts := NewExprOpts(ExprOpts{}, ExpectedTypeOpt(variableType))
		expr, ok := validation.expression(stmt.Value.First(), opts, context, variableType == nil)
		if ok {
			variableValue = expr
			if variableType == nil {
				t := expr.GetType()
				variableType = &t
			}
		}
	}

	if variableType == nil {
		t := NewAtomicInOutType(NewCompileErrorType())
		variableType = &t
	}

	variableDeclaration = NewVariableDeclaration(context, stmt.Name.ToName(), *variableType, variableValue)
	context.AddObject(stmt.Name.Value, &variableDeclaration)

	return variableDeclaration, true
}

func (validation *Validation) expression(expr *read.Expr, opts ExprOpts, context *Context, forInference bool) (Expression, bool) {

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

		call, ok := validation.procedureCall(expr.Call, context)
		if ok {
			ce := NewProcedureExpression(call)
			expression = &ce
		} else {
			return expression, false
		}
	}

	if expr.Boolean != nil {
		e := NewLiteralExpression(*expr.Boolean, "", NewAtomicInOutType(NewBoolType()))
		expression = &e
	}

	if expr.Number != nil {
		e := NewLiteralExpression(*expr.Number, "", NewAtomicInOutType(numericType(*expr.Number)))
		expression = &e
	}

	if expr.Str != nil {
		e := NewLiteralExpression(*expr.Str, "", NewAtomicInOutType(NewStringType()))
		expression = &e
	}

	if expr.Null {
		if forInference {
			validation.addError(NewCanNotInferTypeFromNull(expr.Location))
			return nil, false
		}
		e := NewLiteralExpression("", "", NewAtomicInOutType(NullType{}))
		expression = &e
	}

	if expr.Hex != nil {
		e := NewLiteralExpression(*expr.Hex, "", NewAtomicInOutType(NewByteType()))
		expression = &e
	}

	if expr.Seq != nil {
		var values []Expression
		var elemType = sequenceElementType(opts.ExpectedType)
		for _, val := range expr.Seq.Exprs {
			// Validate every expression in the sequence
			newOpts := NewExprOpts(opts, ExpectedTypeOpt(elemType), ValidationErrOpt(NewIncompatibleElementType))
			value, ok := validation.expression(&val, newOpts, context, false)
			if !ok {
				continue
			}
			values = append(values, value)
			valueType := value.GetType()
			if elemType == nil {
				// If this is the first type en-counted, use it as the sequence type for now
				elemType = &valueType
				continue
			}
			e := commonType(*elemType, valueType)
			if e == nil {
				continue
			}
			elemType = e

		}
		if opts.ExpectedType != nil && opts.ExpectedType.Out.IsSequence {
			expected := opts.ExpectedType.Out.SequenceSize
			if expected != len(values) {
				validation.addError(NewIncorrectSequenceSize(expr.Location, expected, len(values)))
			}
		}
		if elemType == nil {
			return expression, false
		}
		e := NewSequenceExpression(expr.Seq.Count(), values, *elemType)
		expression = &e
	}

	if expression == nil {
		// TODO return error
		return expression, false
	}

	if opts.ExpectedType == nil {
		return expression, true
	}
	et := expression.GetType()
	if areTypesIncompatible(*opts.ExpectedType, et) {
		validation.addError(opts.ErrorBuilder(expr.Location, *opts.ExpectedType, et))
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

	if call.Arguments.Count() != len(object.GetType().In) {
		validation.addError(
			NewIncorrectArgumentCount(
				call.Primary.Value,
				call.Primary.Location,
				call.Arguments.Count(),
				len(object.GetType().In),
			),
		)
	} else {
		validateParamAgainstType = true
	}

	var arguments []Expression
	argumentsOk := true
	for idx := range call.Arguments.Exprs {
		subexpr := &call.Arguments.Exprs[idx]
		var typeToCheckAgainst *InOutType
		if validateParamAgainstType {
			typeToCheckAgainst = object.GetType().In[idx].Type
		}
		opts := NewExprOpts(ExprOpts{}, ExpectedTypeOpt(typeToCheckAgainst), ValidationErrOpt(NewIncompatibleParameterType))
		e, ok := validation.expression(subexpr, opts, context, false)
		if !ok && argumentsOk {
			argumentsOk = false
		}
		arguments = append(arguments, e)
	}
	return NewProcedureCall(context, object, arguments), argumentsOk
}

func (validation *Validation) returnStatement(ret *read.Return, location Location, context *Context) ReturnStatement {

	if context.ReturnType.IsCompileError() {
		return NewReturnStatement(context, nil)
	}

	if context.ReturnType.IsNoType() {
		if ret.Value.IsNotEmpty() {
			validation.addError(NewDoesNotReturnValue(context.CurrentObjectName, ret.Value.First().Location))
			return NewReturnStatement(context, nil)
		}
		return NewReturnStatement(context, nil)
	}

	if ret.Value.IsEmpty() {
		validation.addError(NewReturnValueRequired(context.CurrentObjectName, location, context.ReturnType.String()))
		return NewReturnStatement(context, nil)
	}

	returnType := NewInOutType(nil, context.ReturnType)
	opts := NewExprOpts(ExprOpts{}, ExpectedTypeOpt(&returnType), ValidationErrOpt(NewIncompatibleReturnValueType))
	expr, ok := validation.expression(ret.Value.First(), opts, context, false)
	if !ok {
		return NewReturnStatement(context, nil)
	}

	return NewReturnStatement(context, expr)
}

func numericType(value string) AtomicType {
	if strings.Contains(value, "e+") {
		return NewFloat64Type()
	}

	if strings.Contains(value, ".") {
		return NewFloat64Type()
	}
	return NewInt32Type()

}

func areTypesIncompatible(left, right InOutType) bool {
	if left.Out.IsNeverCompatible() || right.Out.IsNeverCompatible() {
		return true
	}

	if left.Out.IsOptional != right.Out.IsOptional {
		return false
	}

	if left.Out.IsError != right.Out.IsError {
		return false
	}

	if left.Out.Type.GetName() != right.Out.Type.GetName() {
		return true
	}
	if len(left.In) != len(right.In) {
		return true
	}
	for idx := range left.In {
		if areTypesIncompatible(*left.In[idx].Type, *right.In[idx].Type) {
			return true
		}
	}
	return false
}

func (validation *Validation) addError(verr ValidationError) {
	validation.errors = append(validation.errors, verr)
}

func commonType(a, b InOutType) *InOutType {
	if a.Out.IsNeverCompatible() || b.Out.IsNeverCompatible() {
		return nil
	}
	isSameType := a.Out.Type.GetName() == b.Out.Type.GetName()
	isError := a.Out.IsError || b.Out.IsError
	isOptional := !isError && (a.Out.IsOptional || b.Out.IsOptional)

	if isSameType {
		return &InOutType{In: nil, Out: SuperAtomicType{Type: a.Out.Type, IsError: isError, IsOptional: isOptional}}
	}

	for _, other := range a.Out.Type.GetTypesCanCastTo() {
		if other.GetName() == b.Out.Type.GetName() {
			return &InOutType{In: nil, Out: SuperAtomicType{Type: other, IsError: isError, IsOptional: isOptional}}
		}
	}

	for _, other := range b.Out.Type.GetTypesCanCastTo() {
		if other.GetName() == a.Out.Type.GetName() {
			return &InOutType{In: nil, Out: SuperAtomicType{Type: other, IsError: isError, IsOptional: isOptional}}
		}
	}

	return nil
}

func sequenceElementType(seqType *InOutType) *InOutType {
	if seqType == nil {
		return nil
	}
	return &InOutType{In: nil, Out: SuperAtomicType{
		Type: seqType.Out.Type,
	}}
}

type ExprOpts struct {
	ExpectedType *InOutType
	ErrorBuilder func(location Location, expected InOutType, actual InOutType) ValidationError
}

type ExprOpt func(opts *ExprOpts)

func NewExprOpts(base ExprOpts, opts ...ExprOpt) ExprOpts {
	exprOpts := ExprOpts{
		ExpectedType: base.ExpectedType,
		ErrorBuilder: base.ErrorBuilder,
	}
	for _, opt := range opts {
		opt(&exprOpts)
	}
	if exprOpts.ErrorBuilder == nil {
		exprOpts.ErrorBuilder = func(location Location, expected InOutType, actual InOutType) ValidationError {
			return NewIncompatibleType(actual, expected, location)
		}
	}
	return exprOpts
}

func ExpectedTypeOpt(expectedType *InOutType) ExprOpt {
	return func(opts *ExprOpts) {
		opts.ExpectedType = expectedType
	}
}

func ValidationErrOpt(f func(location Location, expected InOutType, actual InOutType) ValidationError) ExprOpt {
	return func(opts *ExprOpts) {
		opts.ErrorBuilder = f
	}
}
