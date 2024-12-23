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
void __prog_proc_apple();
int __prog_proc_pear();
double __prog_proc_orange(int __va,int __vb);

int main(int argc,char** argv) {
    return 0;
}

void __prog_proc_apple() {
}

int __prog_proc_pear() {
    return 1;
}

double __prog_proc_orange(int __va,int __vb) {
    return 1.0;
}
`
	assert(t, actual, expected)
}
