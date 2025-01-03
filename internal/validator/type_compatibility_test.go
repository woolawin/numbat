package validator

import (
	. "numbat/internal/common"
	"testing"
)

func TestSequenceCompatability(t *testing.T) {
	code := `
program do
	var a Int32[3] = [1 2 3] 		// OK
	var b Int32[3] = [1 2 3 4 5] 	// ERROR: Too Many
	var c Int32[5] = [1 2 3]		// ERROR: Too Few
	var d Int32[2] = [true false]	// ERROR: Wrong Type
end
`

	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewIncorrectSequenceSize(loc(4, 19), 3, 5))
	assertValidationError(t, validation, NewIncorrectSequenceSize(loc(5, 19), 5, 3))
	assertValidationError(t, validation, NewIncompatibleElementType(loc(6, 20), NewInt32InOutType(), NewBoolInOutType()))
	assertValidationError(t, validation, NewIncompatibleElementType(loc(6, 25), NewInt32InOutType(), NewBoolInOutType()))

	assertValidationErrorCount(t, validation, 4)
}

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

	assertValidationError(t, validation, NewIncompatibleType(NewStringInOutType(), NewBoolInOutType(), loc(4, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewBoolInOutType(), loc(5, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewFloat64InOutType(), NewBoolInOutType(), loc(6, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewAtomicInOutType(NullType{}), NewBoolInOutType(), loc(7, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewByteInOutType(), NewBoolInOutType(), loc(8, 17)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewInt32InOutType(), loc(10, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewStringInOutType(), NewInt32InOutType(), loc(11, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewFloat64InOutType(), NewInt32InOutType(), loc(13, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewAtomicInOutType(NullType{}), NewInt32InOutType(), loc(14, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewByteInOutType(), NewInt32InOutType(), loc(15, 18)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewFloat64InOutType(), loc(17, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewStringInOutType(), NewFloat64InOutType(), loc(18, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewFloat64InOutType(), loc(19, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewAtomicInOutType(NullType{}), NewFloat64InOutType(), loc(21, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewByteInOutType(), NewFloat64InOutType(), loc(22, 20)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewByteInOutType(), loc(24, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewStringInOutType(), NewByteInOutType(), loc(25, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewByteInOutType(), loc(26, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewFloat64InOutType(), NewByteInOutType(), loc(27, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewAtomicInOutType(NullType{}), NewByteInOutType(), loc(28, 17)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewInt64InOutType(), loc(31, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewStringInOutType(), NewInt64InOutType(), loc(32, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewInt64InOutType(), loc(33, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewFloat64InOutType(), NewInt64InOutType(), loc(34, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewAtomicInOutType(NullType{}), NewInt64InOutType(), loc(35, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewByteInOutType(), NewInt64InOutType(), loc(36, 18)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewFloat32InOutType(), loc(38, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewStringInOutType(), NewFloat32InOutType(), loc(39, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewFloat32InOutType(), loc(40, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewFloat64InOutType(), NewFloat32InOutType(), loc(41, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewAtomicInOutType(NullType{}), NewFloat32InOutType(), loc(42, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewByteInOutType(), NewFloat32InOutType(), loc(43, 20)))

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
	return 0
end

proc i64 Int64 do
	var v Int64
	return v
end

proc bool Bool do
	return true
end

proc byte Byte do
	return 0xFF
end

proc f32 Float32 do
	var f Float32
	return f
end

proc f64 Float64 do
	return 1.0
end

`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewBoolInOutType(), loc(10, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewInt64InOutType(), NewBoolInOutType(), loc(11, 17)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewInt32InOutType(), loc(15, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewByteInOutType(), NewInt32InOutType(), loc(16, 18)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewFloat64InOutType(), loc(20, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewByteInOutType(), NewFloat64InOutType(), loc(21, 20)))

	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewByteInOutType(), loc(25, 17)))
	assertValidationError(t, validation, NewIncompatibleType(NewFloat64InOutType(), NewByteInOutType(), loc(26, 17)))

	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewInt64InOutType(), loc(30, 18)))
	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewInt64InOutType(), loc(31, 18)))

	assertValidationError(t, validation, NewIncompatibleType(NewInt32InOutType(), NewFloat32InOutType(), loc(35, 20)))
	assertValidationError(t, validation, NewIncompatibleType(NewBoolInOutType(), NewFloat32InOutType(), loc(36, 20)))

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
