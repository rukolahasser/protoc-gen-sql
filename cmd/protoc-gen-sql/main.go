package main

import (
	"flag"
	"fmt"
	"github.com/rukolahasser/protoc-gen-sql/sqlgen"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
)

var (
	fileName string
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "See README.md for usage information.\n")
		os.Exit(0)
	}

	flag.StringVar(&fileName, "file", "", "Generated file name.")

	protogen.Options{
		ParamFunc: flag.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			sqlgen.GenerateTableCreationFile(gen, f, fileName)
		}
		return nil
	})
}
