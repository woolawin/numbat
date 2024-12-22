package c

import "testing"

func TestVariables(t *testing.T) {
	src := `
program do
	var a Int32
end
`
	transpiler := NewCTranspiler()
	transpiler.Transpile(readsrc(src))

	actual := transpiler.String()
	expected := `
int main(int argc, char** argv) {
    int __va;
}`
	assert(t, actual, expected)
}
