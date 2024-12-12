package main

import (
	"flag"
	"fmt"
	"io"
	"numbat/internal/read"
	"os"
	"strings"
)

func main() {
	var srcDir string
	var buildDir string
	flag.StringVar(&srcDir, "src-dir", "./src", "Location of directory containing source code to compile")
	flag.StringVar(&buildDir, "build-dir", "./build", "Location of directory to store transpiled code before compiling")

	flag.Parse()

	srcFiles, failed := listSrc(srcDir)
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
	fmt.Println(reader.Source().String())
}

func listSrc(srcDir string) ([]string, bool) {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return nil, true
	}
	var srcFiles []string
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".nbt") {
			srcFiles = append(srcFiles, srcDir+"/"+entry.Name())
		}

		if strings.HasSuffix(entry.Name(), ".numbat") {
			srcFiles = append(srcFiles, srcDir+"/"+entry.Name())
		}
	}
	return srcFiles, false
}
