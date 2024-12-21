package common

import "strings"

const TypeAscii = "Ascii"
const TypeInt32 = "Int32"
const TypeInt64 = "Int64"
const TypeFloat32 = "Float32"
const TypeFloat64 = "Float64"
const TypeBool = "Bool"
const TypeByte = "Byte"

const LiteralTypeNull = "Null"
const LiteralTypeString = "String"

type AtomicType interface {
	GetName() string

	IsNoType() bool
	IsCompileError() bool
	IsNeverCompatible() bool
}

type SuperAtomicType struct {
	Type       AtomicType
	IsOptional bool
	IsError    bool
}

func (s *SuperAtomicType) IsNoType() bool {
	return s.Type.IsNoType()
}

func (s *SuperAtomicType) IsNeverCompatible() bool {
	return s.Type.IsNeverCompatible()
}

func (s *SuperAtomicType) IsCompileError() bool {
	return s.Type.IsCompileError()
}

func (s *SuperAtomicType) String() string {
	return s.Type.GetName()
}

func NewSuperAtomicType(t AtomicType) SuperAtomicType {
	return SuperAtomicType{Type: t}
}

type InType struct {
	Type         *InOutType
	Name         Name
	DefaultValue Expression
}

func NewInType(t InOutType, name Name, defaultValue Expression) InType {
	return InType{Type: &t, Name: name, DefaultValue: defaultValue}
}

type InOutType struct {
	In  []InType
	Out SuperAtomicType
}

func NewInOutType(in []InType, out SuperAtomicType) InOutType {
	return InOutType{In: in, Out: out}
}

func NewAtomicInOutType(atomic AtomicType) InOutType {
	return InOutType{Out: NewSuperAtomicType(atomic)}
}

func NoInOutType() InOutType {
	return NewInOutType(nil, NewSuperAtomicType(NewNoType()))
}

func (t *InOutType) String() string {
	var builder strings.Builder
	if len(t.In) > 0 {
		builder.WriteString("(")
		for idx, in := range t.In {
			builder.WriteString(in.Type.String())
			if idx < len(t.In)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(") ")
	}
	builder.WriteString(t.Out.String())
	return builder.String()
}

func GetStandardTypes() []AtomicType {
	return []AtomicType{
		NewByteType(),
		NewAsciiType(),
		NewInt32Type(),
		NewInt64Type(),
		NewFloat32Type(),
		NewFloat64Type(),
		NewBoolType(),
	}
}

// BoolType ---------------------------------------------------
//
// ------------------------------------------------------------

func NewBoolType() StandardType {
	return StandardType{Out: Name{Value: TypeBool}}
}

func NewBoolInOutType() InOutType {
	return NewAtomicInOutType(NewBoolType())
}

// Int32Type --------------------------------------------------
//
// ------------------------------------------------------------

func NewInt32Type() StandardType {
	return StandardType{Out: Name{Value: TypeInt32}}
}

func NewInt32InOutType() InOutType {
	return NewAtomicInOutType(NewInt32Type())
}

// Int64Type --------------------------------------------------
//
// ------------------------------------------------------------

func NewInt64Type() StandardType {
	return StandardType{Out: Name{Value: TypeInt64}}
}

func NewInt64InOutType() InOutType {
	return NewAtomicInOutType(NewInt64Type())
}

// Float32Type ------------------------------------------------
//
// ------------------------------------------------------------

func NewFloat32Type() StandardType {
	return StandardType{Out: Name{Value: TypeFloat32}}
}

func NewFloat32InOutType() InOutType {
	return NewAtomicInOutType(NewFloat32Type())
}

// Float64Type ------------------------------------------------
//
// ------------------------------------------------------------

func NewFloat64Type() StandardType {
	return StandardType{Out: Name{Value: TypeFloat64}}
}

func NewFloat64InOutType() InOutType {
	return NewAtomicInOutType(NewFloat64Type())
}

// ByteType ---------------------------------------------------
//
// ------------------------------------------------------------

func NewByteType() StandardType {
	return StandardType{Out: Name{Value: TypeByte}}
}

func NewByteInOutType() InOutType {
	return NewAtomicInOutType(NewByteType())
}

// AsciiType --------------------------------------------------
//
// ------------------------------------------------------------

func NewAsciiType() StandardType {
	return StandardType{Out: Name{Value: TypeAscii}}
}

// NullType ---------------------------------------------------
//
// ------------------------------------------------------------

type NullType struct {
}

func (t NullType) GetOut() Name {
	return Name{Value: LiteralTypeNull}
}

func (t NullType) GetName() string {
	return "<null>"
}

func (t NullType) IsCompileError() bool {
	return false
}

func (t NullType) IsNeverCompatible() bool {
	return false
}

func (t NullType) IsNoType() bool {
	return false
}

// StringType -------------------------------------------------
//
// ------------------------------------------------------------

func NewStringType() StandardType {
	return StandardType{Out: Name{Value: LiteralTypeString}}
}

func NewStringInOutType() InOutType {
	return NewAtomicInOutType(NewStringType())
}

// Misc -------------------------------------------------------
//
// ------------------------------------------------------------

type StandardType struct {
	Out Name
}

func (t StandardType) GetName() string {
	return t.Out.Value
}

func (t StandardType) IsCompileError() bool {
	return false
}

func (t StandardType) IsNeverCompatible() bool {
	return false
}

func (t StandardType) IsNoType() bool {
	return false
}

// NoType -----------------------------------------------------
//
// ------------------------------------------------------------

type NoType struct {
}

func NewNoType() NoType {
	return NoType{}
}

func (t NoType) GetName() string {
	return ""
}

func (t NoType) IsCompileError() bool {
	return false
}

func (t NoType) IsNeverCompatible() bool {
	return true
}

func (t NoType) IsNoType() bool {
	return true
}

// CompileErrorType -------------------------------------------
//
// ------------------------------------------------------------

type CompileErrorType struct {
}

func NewCompileErrorType() CompileErrorType {
	return CompileErrorType{}
}

func (e CompileErrorType) GetName() string {
	return "<error>"
}

func (e CompileErrorType) IsCompileError() bool {
	return true
}

func (e CompileErrorType) IsNeverCompatible() bool {
	return true
}

func (e CompileErrorType) IsNoType() bool {
	return false
}
