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

func assert(t *testing.T, actual, expected string) {
	dmp := diffmatchpatch.New()

	a := strings.ReplaceAll(strings.TrimSpace(actual), " ", "•")
	b := strings.ReplaceAll(strings.TrimSpace(expected), " ", "•")
	diffs := dmp.DiffMain(a, b, false)
	if len(diffs) == 1 && diffs[0].Type == diffmatchpatch.DiffEqual {
		return
	}
	t.Fatalf("\n%s\n", dmp.DiffPrettyText(diffs))
}
