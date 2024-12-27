package validator

import (
	. "numbat/internal/common"
	"testing"
)

func TestCommonType(t *testing.T) {

	// Bool and Int32 are not compatible
	result := commonType(NewBoolInOutType(), NewInt32InOutType())
	if result != nil {
		t.Errorf("expect nil, got %s", result)
	}

	// Int32 and Int64 have common type of Int64
	result = commonType(NewInt64InOutType(), NewInt32InOutType())
	if result.Out.Type.GetName() != TypeInt64 {
		t.Errorf("expect Int64, got %s", result)
	}

	// Int32 and Int64 have common type of Int64
	result = commonType(NewFloat32InOutType(), NewFloat64InOutType())
	if result.Out.Type.GetName() != TypeFloat64 {
		t.Errorf("expect Float64, got %s", result)
	}
}
