package c

import "testing"

func TestVariables(t *testing.T) {
	src := `
program do
	var a Int32
	var b Bool = true
	var c = false
	var d = 3
	var e = -3
    var f = 3.14
    var g Int64 = 3
end
`
	transpiler := NewCTranspiler()
	transpiler.Transpile(readsrc(src))

	actual := transpiler.String()
	expected := `
int main(int argc, char** argv) {
    int __va;
    int __vb = 1;
    int __vc = 0;
    int __vd = 3;
    int __ve = -3;
    double __vf = 3.14;
    long int __vg = 3;
}`
	assert(t, actual, expected)
}
