package c

func MangleProcedureName(name string) string {
	return "__prog_proc_" + name
}

func MangleVariableName(name string) string {
	return "__v" + name
}
