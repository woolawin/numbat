package main

import (
	"flag"
	"fmt"
	"io"
	"numbat/internal/read"
	"numbat/internal/transpile/c"
	validator "numbat/internal/validator"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var srcDir string
	var buildDir string
	flag.StringVar(&srcDir, "src-dir", "src", "Location of directory containing source code to compile")
	flag.StringVar(&buildDir, "build-dir", "build", "Location of directory to store transpiled code before compiling")

	flag.Parse()

	srcDirPath, err := filepath.Abs(srcDir)
	if err != nil {
		fmt.Printf("bad src-dir: %s", srcDir)
		return
	}

	buildDirPath, err := filepath.Abs(buildDir)
	if err != nil {
		fmt.Printf("bad build-dir: %s", buildDir)
		return
	}

	srcFiles, failed := listSrc(srcDirPath)
	if failed {
		fmt.Printf("src-dir '%s' does not exist or is not a directory\n", srcDir)
		return
	}
	reader := read.NewSourceReader()
	for _, srcFile := range srcFiles {
		file, err := os.Open(srcFile)
		if err != nil {
			fmt.Printf("failed to open file '%s': %v\n", srcFile, err)
			continue
		}
		data, err := io.ReadAll(file)
		file.Close()
		if err != nil {
			fmt.Printf("failed to read file '%s': %v\n", srcFile, err)
			continue
		}
		reader.Read(string(data), srcFile)
	}

	valid := validator.NewValidation()
	src := valid.Validate(reader.Source())

	if len(valid.GetErrors()) > 0 {
		fmt.Printf("%d Errors Found\n", len(valid.GetErrors()))
		for _, verr := range valid.GetErrors() {
			fmt.Println(verr.Message())
		}
		return
	}

	transpiler := c.NewCTranspiler()
	csrc := transpiler.Transpile(src)
	printer := c.NewCSourcePrinter()

	cFile := filepath.Join(buildDirPath, "main.c")
	cCode := []byte(printer.Print(csrc))
	if err := os.WriteFile(cFile, cCode, 0666); err != nil {
		fmt.Printf("failed to write file '%s': %v\n", cFile, err)
	}
}

func listSrc(srcDir string) ([]string, bool) {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return nil, true
	}
	var srcFiles []string
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".nb") {
			srcFiles = append(srcFiles, srcDir+"/"+entry.Name())
		}

		if strings.HasSuffix(entry.Name(), ".numbat") {
			srcFiles = append(srcFiles, srcDir+"/"+entry.Name())
		}
	}
	return srcFiles, false
}
