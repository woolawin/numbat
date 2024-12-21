package common

const TypeAscii = "Ascii"
const TypeInt32 = "Int32"
const TypeInt64 = "Int64"
const TypeFloat32 = "Float32"
const TypeFloat64 = "Float64"
const TypeBool = "Bool"
const TypeByte = "Byte"

const LiteralTypeNull = "Null"
const LiteralTypeString = "String"

type Type interface {
	GetOut() Name
	GetIn() []InType

	IsCompileError() bool
	IsNeverCompatible() bool
	IsNoType() bool
}

func GetStandardTypes() []Type {
	return []Type{
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

func NewBoolType() Type {
	return StandardType{Out: Name{Value: TypeBool}}
}

// Int32Type --------------------------------------------------
//
// ------------------------------------------------------------

func NewInt32Type() Type {
	return StandardType{Out: Name{Value: TypeInt32}}
}

// Int64Type --------------------------------------------------
//
// ------------------------------------------------------------

func NewInt64Type() Type {
	return StandardType{Out: Name{Value: TypeInt64}}
}

// Float32Type ------------------------------------------------
//
// ------------------------------------------------------------

func NewFloat32Type() Type {
	return StandardType{Out: Name{Value: TypeFloat32}}
}

// Float64Type ------------------------------------------------
//
// ------------------------------------------------------------

func NewFloat64Type() Type {
	return StandardType{Out: Name{Value: TypeFloat64}}
}

// ByteType ---------------------------------------------------
//
// ------------------------------------------------------------

func NewByteType() Type {
	return StandardType{Out: Name{Value: TypeByte}}
}

// AsciiType --------------------------------------------------
//
// ------------------------------------------------------------

func NewAsciiType() Type {
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

func (t NullType) GetIn() []InType {
	return nil
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

func NewStringType() Type {
	return StandardType{Out: Name{Value: LiteralTypeString}}
}

// Misc -------------------------------------------------------
//
// ------------------------------------------------------------

type StandardType struct {
	Out Name
	In  []InType
}

func NewStandardType(out Name, in []InType) StandardType {
	return StandardType{Out: out, In: in}
}

func (st StandardType) GetOut() Name {
	return st.Out
}

func (st StandardType) GetIn() []InType {
	return st.In
}

func (st StandardType) IsCompileError() bool {
	return false
}

func (t StandardType) IsNeverCompatible() bool {
	return false
}

func (t StandardType) IsNoType() bool {
	return false
}

type NoType struct {
}

func NewNoType() NoType {
	return NoType{}
}

func (n NoType) GetOut() Name {
	return Name{}
}

func (n NoType) GetIn() []InType {
	return nil
}

func (n NoType) IsCompileError() bool {
	return false
}

func (t NoType) IsNeverCompatible() bool {
	return true
}

func (t NoType) IsNoType() bool {
	return true
}

type CompileErrorType struct {
}

func NewCompileErrorType() CompileErrorType {
	return CompileErrorType{}
}

func (e CompileErrorType) GetOut() Name {
	return Name{}
}

func (e CompileErrorType) GetIn() []InType {
	return nil
}

func (e CompileErrorType) IsCompileError() bool {
	return true
}

func (e CompileErrorType) IsNeverCompatible() bool {
	return true
}

func (t CompileErrorType) IsNoType() bool {
	return false
}
