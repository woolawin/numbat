package c

import (
	"testing"
)

func TestProgram(t *testing.T) {
	src := `
program do

end
`
	transpiler := NewCTranspiler()
	transpiler.Transpile(readsrc(src))

	actual := transpiler.String()
	expected := `
int main(int argc, char** argv) {
}`
	assert(t, actual, expected)
}

func TestProcedures(t *testing.T) {

	src := `
proc apple do

end

proc pear Int32 do

end

proc orange(a Int32, b Bool) Float64 do
	return 1.0
end
`
	transpiler := NewCTranspiler()
	transpiler.Transpile(readsrc(src))

	actual := transpiler.String()
	expected := `
int main(int argc, char** argv) {
}

void __prog_proc_apple() {}

int __prog_proc_pear() {}

double __prog_proc_orange(a int ,b int ) {}
`
	assert(t, actual, expected)
}
