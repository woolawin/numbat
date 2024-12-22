package validator

import (
	. "numbat/internal/common"
)

func typesAreNotEqual(a, b InOutType) bool {
	if a.Out != b.Out {
		return true
	}
	if len(a.In) != len(b.In) {
		return true
	}

	var inTypeAreNotEqual = func(a, b InType) bool {
		return typesAreNotEqual(*a.Type, *b.Type)
	}

	for i := range a.In {
		if inTypeAreNotEqual(a.In[i], b.In[i]) {
			return true
		}
	}
	return false
}

func inoutType(in AtomicType, out AtomicType) InOutType {
	return NewInOutType(
		[]InType{
			NewInType(
				outType(in),
				Name{Value: "", Location: Location{}},
				nil,
			),
		},
		NewSuperAtomicType(out),
	)
}

func outType(out AtomicType) InOutType {
	return NewInOutType(nil, NewSuperAtomicType(out))
}
