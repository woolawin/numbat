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
}

func GetStandardTypes() []Type {
	return []Type{
		ByteType{},
		AsciiType{},
		Int32Type{},
		Int64Type{},
		Float32Type{},
		Float64Type{},
		BoolType{},
	}
}

// BoolType ---------------------------------------------------
//
// ------------------------------------------------------------
type BoolType struct {
}

func (t BoolType) GetOut() Name {
	return Name{Value: TypeBool}
}

func (t BoolType) GetIn() []InType {
	return nil
}

func (t BoolType) IsCompileError() bool {
	return false
}

func (t BoolType) IsNeverCompatible() bool {
	return false
}

// Int32Type --------------------------------------------------
//
// ------------------------------------------------------------

type Int32Type struct {
}

func (t Int32Type) GetOut() Name {
	return Name{Value: TypeInt32}
}

func (t Int32Type) GetIn() []InType {
	return nil
}

func (t Int32Type) IsCompileError() bool {
	return false
}

func (t Int32Type) IsNeverCompatible() bool {
	return false
}

// Int64Type --------------------------------------------------
//
// ------------------------------------------------------------

type Int64Type struct {
}

func (t Int64Type) GetOut() Name {
	return Name{Value: TypeInt64}
}

func (t Int64Type) GetIn() []InType {
	return nil
}

func (t Int64Type) IsCompileError() bool {
	return false
}

func (t Int64Type) IsNeverCompatible() bool {
	return false
}

// Float32Type ------------------------------------------------
//
// ------------------------------------------------------------

type Float32Type struct {
}

func (t Float32Type) GetOut() Name {
	return Name{Value: TypeFloat32}
}

func (t Float32Type) GetIn() []InType {
	return nil
}

func (t Float32Type) IsCompileError() bool {
	return false
}

func (t Float32Type) IsNeverCompatible() bool {
	return false
}

// Float64Type ------------------------------------------------
//
// ------------------------------------------------------------

type Float64Type struct {
}

func (t Float64Type) GetOut() Name {
	return Name{Value: TypeFloat64}
}

func (t Float64Type) GetIn() []InType {
	return nil
}

func (t Float64Type) IsCompileError() bool {
	return false
}

func (t Float64Type) IsNeverCompatible() bool {
	return false
}

// ByteType ---------------------------------------------------
//
// ------------------------------------------------------------

type ByteType struct {
}

func (t ByteType) GetOut() Name {
	return Name{Value: TypeByte}
}

func (t ByteType) GetIn() []InType {
	return nil
}

func (t ByteType) IsCompileError() bool {
	return false
}

func (t ByteType) IsNeverCompatible() bool {
	return false
}

// AsciiType --------------------------------------------------
//
// ------------------------------------------------------------

type AsciiType struct {
}

func (t AsciiType) GetOut() Name {
	return Name{Value: TypeAscii}
}

func (t AsciiType) GetIn() []InType {
	return nil
}

func (t AsciiType) IsCompileError() bool {
	return false
}

func (t AsciiType) IsNeverCompatible() bool {
	return false
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

// StringType -------------------------------------------------
//
// ------------------------------------------------------------

type StringType struct {
}

func (t StringType) GetOut() Name {
	return Name{Value: LiteralTypeString}
}

func (t StringType) GetIn() []InType {
	return nil
}

func (t StringType) IsCompileError() bool {
	return false
}

func (t StringType) IsNeverCompatible() bool {
	return false
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
