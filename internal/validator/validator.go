package validator

import (
	"fmt"
	"numbat/internal/common"
	"numbat/internal/read"
	"strings"
	"unicode"
)

type UnknownType struct {
	TypeName string
}

func (verr UnknownType) Message() string {
	return fmt.Sprintf("unknown type: %s", verr.TypeName)
}

type UnInferrableVar struct {
	VarName string
}

func (verr UnInferrableVar) Message() string {
	return fmt.Sprintf("can not infer type for variable: %s", verr.VarName)
}

type NoExprToInferVariableType struct {
	VarName string
}

func (verr NoExprToInferVariableType) Message() string {
	return fmt.Sprintf("can not infer type for variable without expression: %s", verr.VarName)
}

type CanNotInferTypeFromNull struct {
	VarName string
}

func (verr CanNotInferTypeFromNull) Message() string {
	return fmt.Sprintf("can not infer type for variable with null expression: %s", verr.VarName)
}

type CanNotInferTypeFromOtherVariable struct {
	VarName string
}

func (verr CanNotInferTypeFromOtherVariable) Message() string {
	return fmt.Sprintf("type inference from variable not supported: %s", verr.VarName)
}

type CanNotInferTypeFromCall struct {
	VarName string
}

func (verr CanNotInferTypeFromCall) Message() string {
	return fmt.Sprintf("type inference from variable not supported: %s", verr.VarName)
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
				validation.errors = append(validation.errors, UnknownType{TypeName: t.Out.Name})
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
	// Check if variable name is already taken

	var variableDeclaration *common.VariableDeclaration
	var variableType *common.Type
	if stmt.VarType != nil {
		t := validation.validateType(stmt.VarType, context, true)
		variableType = &t
	}
	object, found := context.GetObject(stmt.Name.Value)
	if found {
		validation.addError(newNameConflict(stmt.Name.ToName(), object.GetName().Location))
	} else {
		exprType := validation.validateType(stmt.VarType, context, true)
		if stmt.VarType == nil {
			variableType = &exprType
		}
		// TODO check if expr type is compatible with var type
		a := common.NewVariableDeclaration(stmt.Name.ToName(), *variableType, nil)
		variableDeclaration = &a
		context.AddObject(stmt.Name.Value, variableDeclaration)
	}

	if len(stmt.Exprs) == 1 {
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

func (validation *Validation) HasProgram(src *read.Source) {
	if src.Program == nil {
		validation.addError(&ProgramingMissing{})
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
