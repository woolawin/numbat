package read

import (
	"github.com/google/go-cmp/cmp"
	"numbat/samples"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	src := Read(samples.HelloWorld).String()
	expected := `
PROGRAM
	DO
	CALL log #str "Hello World"
`
	assert(t, src, expected)
}

func TestComments(t *testing.T) {
	src := Read(samples.Comments).String()
	expected := `
PROGRAM
	DO
	CALL log #str "yes"
`
	assert(t, src, expected)
}

func TestProcs(t *testing.T) {
	src := Read(samples.Procs).String()
	expected := `
PROC foo
	DO

PROC bar
	OUT Int
	DO
	RET #num 1

PROC inc1
	IN num Int
	OUT Int
	DO
	RET $num

PROC swap
	IN a Int
	IN b Int
	OUT Int
	DO

PROC baz
	IN msg Str
	IN caps Bool
	IN times Int #num 1
	DO

PROC map
	IN arr Array
	IN fn (a Int) Int
	OUT Array
	DO
`
	assert(t, src, expected)
}

func TestVars(t *testing.T) {
	src := Read(samples.Vars).String()
	expected := `
PROC foo
	DO
	LET a #num 0
	LET b Int32 #num 1
	LET c Int32
	LET d Int32 &(baz)
	LET e Int64 #num 1e+22
	LET f $a
	LET h #hex 0x3F
`
	assert(t, src, expected)
}

func assert(t *testing.T, a, b string) {
	diff := cmp.Diff(strings.TrimSpace(a), strings.TrimSpace(b))
	if diff != "" {
		t.Fatalf(diff)
	}
}
