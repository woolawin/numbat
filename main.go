package main

import (
	"fmt"
	"numbat/internal/read"
)

func main() {

	code := `

program do
	(main)
end

proc main do
	return   // This is a comment at the end of a 
// This is a comment at the start of a line
end

/*
proc just_a_comment do

end
*/

proc hello Int do
	(log: name)
	(text to_int32: str)
	let max Int64 = 1e+22
	let avg = 0x3F
	let default = 1
	let bar Int16 
	(send: 1 true default "Welcome")
	(exit)
end

proc bye(msg Str, times Int = 0, fmt Str = "msg %s") Bool do
	let encoded = (encode: 32)
	return true
end

proc plus(num Int) do
	num = 3
	foo, bar = true, false
	return (get)
end
`

	src := read.Read(code)
	fmt.Println(src.String())
}
