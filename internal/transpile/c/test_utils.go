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

	diffs := dmp.DiffMain(strings.TrimSpace(actual), strings.TrimSpace(expected), false)
	if len(diffs) == 1 && diffs[0].Type == diffmatchpatch.DiffEqual {
		return
	}
	t.Fatal(dmp.DiffPrettyText(diffs))
}
