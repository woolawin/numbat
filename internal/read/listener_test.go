package read

import (
	"github.com/google/go-cmp/cmp"
	"numbat/samples"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	src := readSample(samples.HelloWorld).String()
	expected := `
PROGRAM
	DO
	CALL log #str "Hello World"
`
	assert(t, src, expected)
}

func TestComments(t *testing.T) {
	src := readSample(samples.Comments).String()
	expected := `
PROGRAM
	DO
	CALL log #str "yes"
`
	assert(t, src, expected)
}

func TestProcs(t *testing.T) {
	src := readSample(samples.Procs).String()
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
	src := readSample(samples.Vars).String()
	expected := `
PROC foo
	DO
	VAR a #num 0
	VAR b Int32 #num 1
	VAR c Int32
	VAR d Int32 &(baz)
	VAR e Int64 #num 1e+22
	VAR g Int64 #kg 12
	VAR f $a
	VAR h #hex 0x3F
	VAR i #num 1.2
	VAR j Float32 #num -3.8
	VAR k #kg/m^3 1
	VAR l #m^3/kg 2
	VAR m &(baz.bar)
`
	assert(t, src, expected)
}

func TestCalls(t *testing.T) {
	src := readSample(samples.Calls).String()
	expected := `
PROC main
	DO
	CALL foo
	CALL foo.bar
	CALL foo #str "baz"
	CALL foo.bar #str "baz"
	CALL alot #num 1 #bool true #bool false #num -45 #str "msg" #hex 0xFFF #num 3e+14
	VAR a &(foo)
	VAR b &(foo #str "baz" #bool true)
`
	assert(t, src, expected)
}

func readSample(sample string) *Source {
	listener := NewSourceReader()
	listener.Read(sample)
	return listener.Source()
}

func assert(t *testing.T, a, b string) {
	diff := cmp.Diff(strings.TrimSpace(a), strings.TrimSpace(b))
	if diff != "" {
		t.Fatalf(diff)
	}
}
