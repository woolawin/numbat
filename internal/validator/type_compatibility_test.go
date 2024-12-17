package validator

import (
	"numbat/internal/common"
	"testing"
)

func TestIncompatibleLiteralTypes(t *testing.T) {
	code := `
program do
	var a_1 Bool = true
	var a_2 Bool = "message"
	var a_3 Bool = 2
	var a_4 Bool = 3.14
	var a_5 Bool = null
	var a_6 Bool = 0xD2

	var b_1 Int32 = true
	var b_2 Int32 = "message"
	var b_3 Int32 = 2
	var b_4 Int32 = 3.14
	var b_5 Int32 = null
	var b_6 Int32 = 0xD2

	var c_1 Float64 = true
	var c_2 Float64 = "message"
	var c_3 Float64 = 2
	var c_4 Float64 = 3.14
	var c_5 Float64 = null
	var c_6 Float64 = 0xD2

	var d_1 Byte = true
	var d_2 Byte = "message"
	var d_3 Byte = 2
	var d_4 Byte = 3.14
	var d_5 Byte = null
	var d_6 Byte = 0xD2

	var e_1 Int64 = true
	var e_2 Int64 = "message"
	var e_3 Int64 = 2
	var e_4 Int64 = 3.14
	var e_5 Int64 = null
	var e_6 Int64 = 0xD2

	var f_1 Float32 = true
	var f_2 Float32 = "message"
	var f_3 Float32 = 2
	var f_4 Float32 = 3.14
	var f_5 Float32 = null
	var f_6 Float32 = 0xD2
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewIncompatibleType(common.NewStringType(), common.NewBoolType(), loc(4, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewBoolType(), loc(5, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewFloat64Type(), common.NewBoolType(), loc(6, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NullType{}, common.NewBoolType(), loc(7, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewByteType(), common.NewBoolType(), loc(8, 17)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewInt32Type(), loc(10, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewStringType(), common.NewInt32Type(), loc(11, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewFloat64Type(), common.NewInt32Type(), loc(13, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NullType{}, common.NewInt32Type(), loc(14, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewByteType(), common.NewInt32Type(), loc(15, 18)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewFloat64Type(), loc(17, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewStringType(), common.NewFloat64Type(), loc(18, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewFloat64Type(), loc(19, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NullType{}, common.NewFloat64Type(), loc(21, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewByteType(), common.NewFloat64Type(), loc(22, 20)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewByteType(), loc(24, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewStringType(), common.NewByteType(), loc(25, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewByteType(), loc(26, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewFloat64Type(), common.NewByteType(), loc(27, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NullType{}, common.NewByteType(), loc(28, 17)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewInt64Type(), loc(31, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewStringType(), common.NewInt64Type(), loc(32, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewInt64Type(), loc(33, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewFloat64Type(), common.NewInt64Type(), loc(34, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NullType{}, common.NewInt64Type(), loc(35, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewByteType(), common.NewInt64Type(), loc(36, 18)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewFloat32Type(), loc(38, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewStringType(), common.NewFloat32Type(), loc(39, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewFloat32Type(), loc(40, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewFloat64Type(), common.NewFloat32Type(), loc(41, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NullType{}, common.NewFloat32Type(), loc(42, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewByteType(), common.NewFloat32Type(), loc(43, 20)))

	assertValidationErrorCount(t, validation, 32)
}

func TestIncompatibleVarAndProcTypes(t *testing.T) {
	code := `
program do
	var i Int32
	var l Int64
	var b Bool
	var byt Byte
	var f Float32
	var d Float64

	var a_1 Bool = i
	var a_2 Bool = i64()
	var a_3 Bool = b
	var a_4 Bool = bool()

	var b_1 Int32 = b
	var b_2 Int32 = byte()
	var b_3 Int32 = i
	var b_4 Int32 = i32()

	var c_1 Float64 = b
	var c_2 Float64 = byte()
	var c_3 Float64 = d
	var c_4 Float64 = f64()

	var d_1 Byte = i32
	var d_2 Byte = f64()
	var d_3 Byte = byt
	var d_4 Byte = byte()

	var e_1 Int64 = b
	var e_2 Int64 = bool()
	var e_3 Int64 = l
	var e_4 Int64 = i64()

	var f_1 Float32 = i32
	var f_2 Float32 = bool()
	var f_3 Float32 = f
	var f_4 Float32 = f32()
end

proc i32 Int32 do
end

proc i64 Int64 do
end

proc bool Bool do
end

proc byte Byte do
end

proc f32 Float32 do
end

proc f64 Float64 do
end

`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewBoolType(), loc(10, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewInt64Type(), common.NewBoolType(), loc(11, 17)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewInt32Type(), loc(15, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewByteType(), common.NewInt32Type(), loc(16, 18)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewFloat64Type(), loc(20, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewByteType(), common.NewFloat64Type(), loc(21, 20)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewByteType(), loc(25, 17)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewFloat64Type(), common.NewByteType(), loc(26, 17)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewInt64Type(), loc(30, 18)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewInt64Type(), loc(31, 18)))

	assertValidationError(t, validation, NewIncompatibleType(common.NewInt32Type(), common.NewFloat32Type(), loc(35, 20)))
	assertValidationError(t, validation, NewIncompatibleType(common.NewBoolType(), common.NewFloat32Type(), loc(36, 20)))

	assertValidationErrorCount(t, validation, 12)
}

func TestIncompatibleWhenObjectIsNotFound(t *testing.T) {
	code := `
program do
	var a Int32 = unknown_var
	var b Int32 = unknown_proc()
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewUnknownObject("unknown_var", loc(3, 16)))
	assertValidationError(t, validation, NewUnknownObject("unknown_proc", loc(4, 16)))

	assertValidationErrorCount(t, validation, 2)
}
