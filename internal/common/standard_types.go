package common

const TypeAscii = "Ascii"
const TypeStr = "Str"
const TypeInt32 = "Int32"
const TypeInt64 = "Int64"
const TypeFloat32 = "Float32"
const TypeFloat64 = "Float64"
const TypeUint32 = "Uint32"
const TypeUint64 = "Uint64"
const TypeBool = "Bool"
const TypeByte = "Byte"
const TypeSize = "Size"

const LiteralTypeNull = "Null"
const LiteralTypeString = "String"

func GetStandardTypes() []Type {
	return []Type{
		{Out: Name{Value: TypeByte}},
		{Out: Name{Value: TypeAscii}},
		{Out: Name{Value: TypeInt32}},
		{Out: Name{Value: TypeInt64}},
		{Out: Name{Value: TypeFloat32}},
		{Out: Name{Value: TypeFloat64}},
		{Out: Name{Value: TypeBool}},
	}
}

type BoolType struct {
}

func (t BoolType) Name() string {
	return "Bool"
}

type Int32Type struct {
}

func (t Int32Type) Name() string {
	return "Int32"
}

type Int64Type struct {
}

func (t Int64Type) Name() string {
	return "Int64"
}

type Float32Type struct {
}

func (t Float32Type) Name() string {
	return "Float32"
}

type Float64Type struct {
}

func (t Float64Type) Name() string {
	return "Float64"
}

type ByteType struct {
}

func (t ByteType) Name() string {
	return "Byte"
}
