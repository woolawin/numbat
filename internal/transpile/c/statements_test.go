package c

import (
	"testing"
)

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
	var h Int32 = apple() 
	var i Int32 = banana(true)
    var j Int32 = pear(false, -4, -4.12)
end

proc apple Int32 do
	return 1
end

proc banana(a Bool) Int32 do
	return 1
end

proc pear(a Bool, b Int64, c Float64) Int32 do
	return 1
end
`
	actual := transpile(src)

	expected := `
int __prog_proc_apple();
int __prog_proc_banana(int __va);
int __prog_proc_pear(int __va,long int __vb,double __vc);

int main(int argc,char** argv) {
    int __va;
    int __vb = 1;
    int __vc = 0;
    int __vd = 3;
    int __ve = -3;
    double __vf = 3.14;
    long int __vg = 3;
    int __vh = __prog_proc_apple();
    int __vi = __prog_proc_banana(1);
    int __vj = __prog_proc_pear(0,-4,-4.12);
}

int __prog_proc_apple() {
}

int __prog_proc_banana(int __va) {
}

int __prog_proc_pear(int __va,long int __vb,double __vc) {
}
`
	assert(t, actual, expected)
}
