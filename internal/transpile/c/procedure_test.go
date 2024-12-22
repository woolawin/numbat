package c

import (
	"testing"
)

func TestProcedureNoParamsNoReturn(t *testing.T) {

	src := `
proc apple do

end
`
	transpiler := NewCTranspiler()
	transpiler.Transpile(readsrc(src))

	actual := transpiler.String()
	expected := `
void __prog_proc_apple ( ) {
}
`
	assert(t, actual, expected)
}

func TestProcedureWithReturn(t *testing.T) {

	src := `
proc apple Int32 do

end
`
	transpiler := NewCTranspiler()
	transpiler.Transpile(readsrc(src))

	actual := transpiler.String()
	expected := `
int __prog_proc_apple ( ) {
}
`
	assert(t, actual, expected)
}

func TestProcedureParamsAndReturn(t *testing.T) {

	src := `
proc apple(a Int32, b Bool) Float64 do
	return 1.0
end
`
	transpiler := NewCTranspiler()
	transpiler.Transpile(readsrc(src))

	actual := transpiler.String()
	expected := `
double __prog_proc_apple ( a int , b int ) {
}
`
	assert(t, actual, expected)
}
