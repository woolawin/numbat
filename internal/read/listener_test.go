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
	CALL log "Hello World"
`
	assert(t, src, expected)
}

func TestComments(t *testing.T) {
	src := readSample(samples.Comments).String()
	expected := `
PROGRAM
	DO
	CALL log "yes"
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
	RET 1

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
	IN times Int 1
	DO

PROC map
	IN arr Array
	IN fn (a Int) Int
	OUT Array
	DO

PROC hello
	DO
`
	assert(t, src, expected)
}

func TestVars(t *testing.T) {
	src := readSample(samples.Vars).String()
	expected := `
PROC foo
	DO
	VAR a 0
	VAR b Int32 1
	VAR c Int32
	VAR d Int32 &(baz)
	VAR e Int64 1e+22
	VAR g Int64 12#kg
	VAR f $a
	VAR h 0x3F
	VAR i 1.2
	VAR j Float32 -3.8
	VAR k 1#kg/m^3
	VAR l 2#m^3/kg
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
	CALL foo
	CALL foo.bar
	CALL foo "baz"
	CALL foo.bar "baz"
	CALL foo.bar "baz"
	CALL alot 1 true false -45 "msg" 0xFFF 3e+14
	VAR a &(foo)
	VAR b &(foo "baz" true)
	CALL foo &(bar) &(baz true &(doh)) -1.7
`
	assert(t, src, expected)
}

func TestCallNested(t *testing.T) {
	src := readSample(samples.NestedCalls).String()
	expected := `
PROC main
	DO
	CALL foo &(bar) &(baz true &(doh)) -1.7
	CALL foo &(bar true &(baz) -1.8) "Hello"
	CALL foo &(lib2.bar "MSG" &(free)) &(baz true &(lib.doh)) -1.7
`
	assert(t, src, expected)
}

func readSample(sample string) *Source {
	listener := NewSourceReader()
	listener.Read(sample, "")
	return listener.Source()
}

func TestMultiplePrograms(t *testing.T) {
	code := `
program do
	var a = 1
end

program do
	var b = 2
end
`
	src := readSample(code).String()
	expected := `
PROGRAM
	DO
	VAR a 1

PROGRAM
	DO
	VAR b 2
`
	assert(t, src, expected)
}

func assert(t *testing.T, a, b string) {
	diff := cmp.Diff(strings.TrimSpace(a), strings.TrimSpace(b))
	if diff != "" {
		t.Fatalf(diff)
	}
}
