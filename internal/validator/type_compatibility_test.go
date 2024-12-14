package validator

import (
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
end
`
	src := readsrc(code)
	validation := NewValidation()
	validation.Validate(src)

	assertValidationError(t, validation, newIncompatibleType("String", "Bool", loc(4, 17)))
	assertValidationError(t, validation, newIncompatibleType("Int32", "Bool", loc(5, 17)))
	assertValidationError(t, validation, newIncompatibleType("Float64", "Bool", loc(6, 17)))
	assertValidationError(t, validation, newIncompatibleType("Null", "Bool", loc(7, 17)))
	assertValidationError(t, validation, newIncompatibleType("Byte", "Bool", loc(8, 17)))

	assertValidationError(t, validation, newIncompatibleType("Bool", "Int32", loc(10, 18)))
	assertValidationError(t, validation, newIncompatibleType("String", "Int32", loc(11, 18)))
	assertValidationError(t, validation, newIncompatibleType("Float64", "Int32", loc(13, 18)))
	assertValidationError(t, validation, newIncompatibleType("Null", "Int32", loc(14, 18)))
	assertValidationError(t, validation, newIncompatibleType("Byte", "Int32", loc(15, 18)))

	assertValidationError(t, validation, newIncompatibleType("Bool", "Float64", loc(17, 20)))
	assertValidationError(t, validation, newIncompatibleType("String", "Float64", loc(18, 20)))
	assertValidationError(t, validation, newIncompatibleType("Int32", "Float64", loc(19, 20)))
	assertValidationError(t, validation, newIncompatibleType("Null", "Float64", loc(21, 20)))
	assertValidationError(t, validation, newIncompatibleType("Byte", "Float64", loc(22, 20)))

	assertValidationError(t, validation, newIncompatibleType("Bool", "Byte", loc(24, 17)))
	assertValidationError(t, validation, newIncompatibleType("String", "Byte", loc(25, 17)))
	assertValidationError(t, validation, newIncompatibleType("Int32", "Byte", loc(26, 17)))
	assertValidationError(t, validation, newIncompatibleType("Float64", "Byte", loc(27, 17)))
	assertValidationError(t, validation, newIncompatibleType("Null", "Byte", loc(28, 17)))

	assertValidationErrorCount(t, validation, 20)
}
