package c

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"numbat/internal/common"
	"numbat/internal/read"
	"numbat/internal/validator"
	"strings"
	"testing"
)

func readsrc(sample string) *common.Source {
	reader := read.NewSourceReader()
	reader.Read(sample, "")
	src := reader.Source()
	validation := validator.NewValidation()
	return validation.Validate(src)
}

func transpile(sample string) string {
	transpiler := NewCTranspiler()
	tsrc := transpiler.Transpile(readsrc(sample))
	printer := NewCSourcePrinter()
	return printer.Print(tsrc)
}

func assert(t *testing.T, actual, expected string) {
	dmp := diffmatchpatch.New()

	a := strings.ReplaceAll(strings.TrimSpace(actual), " ", "•")
	a = strings.ReplaceAll(a, "\n", "␊\n")
	b := strings.ReplaceAll(strings.TrimSpace(expected), " ", "•")
	b = strings.ReplaceAll(b, "\n", "␊\n")
	diffs := dmp.DiffMain(a, b, false)
	if len(diffs) == 1 && diffs[0].Type == diffmatchpatch.DiffEqual {
		return
	}
	t.Fatalf("\n%s\n", dmp.DiffPrettyText(diffs))
}
