package samples

import (
	_ "embed"
)

//go:embed hello_world.numbat
var HelloWorld string

//go:embed comments.numbat
var Comments string

//go:embed procs.numbat
var Procs string

//go:embed vars.numbat
var Vars string

//go:embed calls.numbat
var Calls string

//go:embed nest_call.numbat
var NestedCalls string

//go:embed sequences.numbat
var Sequences string
