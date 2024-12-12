package main

import (
	"flag"
	"fmt"
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

	fmt.Println(srcFiles)
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
