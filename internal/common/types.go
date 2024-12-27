package common

import (
	"strconv"
	"strings"
)

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

	GetTypesCanCastTo() []AtomicType
}

type SuperAtomicType struct {
	Type         AtomicType
	IsSequence   bool
	SequenceSize int
	IsOptional   bool
	IsError      bool
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
	var builder strings.Builder
	builder.WriteString(s.Type.GetName())
	if s.IsSequence {
		builder.WriteString("[")
		if s.SequenceSize != 0 {
			builder.WriteString(strconv.Itoa(s.SequenceSize))
		}
		builder.WriteString("]")
	}
	return builder.String()
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

func NewSeqType(out AtomicType, size int) InOutType {
	return InOutType{Out: SuperAtomicType{Type: out, IsSequence: true, SequenceSize: size}}
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

func NewBoolSuperType() SuperAtomicType {
	return NewSuperAtomicType(NewBoolType())
}

// Int32Type --------------------------------------------------
//
// ------------------------------------------------------------

func NewInt32Type() StandardType {
	return StandardType{
		Out:         Name{Value: TypeInt32},
		CastToTypes: []AtomicType{NewInt64Type(), NewFloat64Type()},
	}
}

func NewInt32InOutType() InOutType {
	return NewAtomicInOutType(NewInt32Type())
}

func NewInt32SuperType() SuperAtomicType {
	return NewSuperAtomicType(NewInt32Type())
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

func NewInt64SuperType() SuperAtomicType {
	return NewSuperAtomicType(NewInt64Type())
}

// Float32Type ------------------------------------------------
//
// ------------------------------------------------------------

func NewFloat32Type() StandardType {
	return StandardType{
		Out:         Name{Value: TypeFloat32},
		CastToTypes: []AtomicType{NewInt64Type(), NewFloat64Type()},
	}
}

func NewFloat32InOutType() InOutType {
	return NewAtomicInOutType(NewFloat32Type())
}

func NewFloat32SuperType() SuperAtomicType {
	return NewSuperAtomicType(NewFloat32Type())
}

// Float64Type ------------------------------------------------
//
// ------------------------------------------------------------

func NewFloat64Type() StandardType {
	return StandardType{
		Out:         Name{Value: TypeFloat64},
		CastToTypes: []AtomicType{NewInt64Type()},
	}
}

func NewFloat64InOutType() InOutType {
	return NewAtomicInOutType(NewFloat64Type())
}

func NewFloat64SuperType() SuperAtomicType {
	return NewSuperAtomicType(NewFloat64Type())
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

func (t NullType) GetTypesCanCastTo() []AtomicType {
	return nil
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
	Out         Name
	CastToTypes []AtomicType
}

func (t StandardType) GetTypesCanCastTo() []AtomicType {
	return t.CastToTypes
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

func (t NoType) GetTypesCanCastTo() []AtomicType {
	return nil
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

func (e CompileErrorType) GetTypesCanCastTo() []AtomicType {
	return nil
}