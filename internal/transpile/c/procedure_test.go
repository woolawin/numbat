package c

import (
	"testing"
)

func TestProgram(t *testing.T) {
	src := `
program do

end
`
	actual := transpile(src)
	expected := `
int main(int argc,char** argv) {
    return 0;
}
`
	assert(t, actual, expected)
}

func TestProcedures(t *testing.T) {

	src := `
proc apple do

end

proc pear Int32 do
	return 1
end

proc orange(a Int32, b Bool) Float64 do
	return 1.0
end
`

	actual := transpile(src)
	expected := `
void __prog_papple();
int __prog_ppear();
double __prog_porange(int __va,int __vb);

int main(int argc,char** argv) {
    return 0;
}

void __prog_papple() {
}

int __prog_ppear() {
    return 1;
}

double __prog_porange(int __va,int __vb) {
    return 1.0;
}
`
	assert(t, actual, expected)
}
