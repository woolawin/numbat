package c

func MangleProcedureName(name string) string {
	return "__prog_p" + name
}

func MangleVariableName(name string) string {
	return "__v" + name
}
